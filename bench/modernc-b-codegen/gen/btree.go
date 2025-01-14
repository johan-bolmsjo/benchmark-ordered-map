// Copyright 2014 The b Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package b // import "modernc.org/b"

import bench "modernc.org/benchmark-ordered-map/fixture"

import (
	"fmt"
	"io"
	"sync"
)

const (
	kx = 32 //TODO benchmark tune this number if using custom key/value type(s).
	kd = 32 //TODO benchmark tune this number if using custom key/value type(s).
)

func init() {
	if kd < 1 {
		panic(fmt.Errorf("kd %d: out of range", kd))
	}

	if kx < 2 {
		panic(fmt.Errorf("kx %d: out of range", kx))
	}
}

var (
	btDPool = sync.Pool{New: func() interface{} { return &d{} }}
	btEPool = btEpool{sync.Pool{New: func() interface{} { return &Enumerator{} }}}
	btTPool = btTpool{sync.Pool{New: func() interface{} { return &Tree{} }}}
	btXPool = sync.Pool{New: func() interface{} { return &x{} }}
)

type btTpool struct{ sync.Pool }

func (p *btTpool) get(cmp Cmp) *Tree {
	x := p.Get().(*Tree)
	x.cmp = cmp
	return x
}

type btEpool struct{ sync.Pool }

func (p *btEpool) get(err error, hit bool, i int, k bench.Key, q *d, t *Tree, ver int64) *Enumerator {
	x := p.Get().(*Enumerator)
	x.err, x.hit, x.i, x.k, x.q, x.t, x.ver = err, hit, i, k, q, t, ver
	return x
}

type (
	// Cmp compares a and b. Return value is:
	//
	//	< 0 if a <  b
	//	  0 if a == b
	//	> 0 if a >  b
	//
	Cmp func(a, b bench.Key) int

	d struct { // data page
		c int
		d [2*kd + 1]de
		n *d
		p *d
		dTree
	}

	de struct { // d element
		k bench.Key
		v bench.Value
	}

	// Enumerator captures the state of enumerating a tree. It is returned
	// from the Seek* methods. The enumerator is aware of any mutations
	// made to the tree in the process of enumerating it and automatically
	// resumes the enumeration at the proper key, if possible.
	//
	// However, once an Enumerator returns io.EOF to signal "no more
	// items", it does no more attempt to "resync" on tree mutation(s).  In
	// other words, io.EOF from an Enumerator is "sticky" (idempotent).
	Enumerator struct {
		err error
		hit bool
		i   int
		k   bench.Key
		q   *d
		t   *Tree
		ver int64
	}

	// Tree is a B+tree.
	Tree struct {
		c     int
		cmp   Cmp
		first *d
		last  *d
		r     interface{}
		ver   int64
		treeInst
	}

	xe struct { // x element
		ch interface{}
		k  bench.Key
	}

	x struct { // index page
		c int
		x [2*kx + 2]xe
	}
)

var ( // R/O zero values
	zd  d
	zde de
	ze  Enumerator
	zk  bench.Key
	zt  Tree
	zx  x
	zxe xe
)

func clr(q interface{}) {
	switch x := q.(type) {
	case *x:
		for i := 0; i <= x.c; i++ { // Ch0 Sep0 ... Chn-1 Sepn-1 Chn
			clr(x.x[i].ch)
		}
		*x = zx
		btXPool.Put(x)
	case *d:
		*x = zd
		btDPool.Put(x)
	}
}

// -------------------------------------------------------------------------- x

func newX(ch0 interface{}) *x {
	r := btXPool.Get().(*x)
	r.x[0].ch = ch0
	return r
}

func (q *x) extract(i int) {
	q.c--
	if i < q.c {
		copy(q.x[i:], q.x[i+1:q.c+1])
		q.x[q.c].ch = q.x[q.c+1].ch
		q.x[q.c].k = zk  // GC
		q.x[q.c+1] = zxe // GC
	}
}

func (q *x) insert(i int, k bench.Key, ch interface{}) *x {
	c := q.c
	if i < c {
		q.x[c+1].ch = q.x[c].ch
		copy(q.x[i+2:], q.x[i+1:c])
		q.x[i+1].k = q.x[i].k
	}
	c++
	q.c = c
	q.x[i].k = k
	q.x[i+1].ch = ch
	return q
}

func (q *x) siblings(i int) (l, r *d) {
	if i >= 0 {
		if i > 0 {
			l = q.x[i-1].ch.(*d)
		}
		if i < q.c {
			r = q.x[i+1].ch.(*d)
		}
	}
	return
}

// -------------------------------------------------------------------------- d

func (l *d) mvL(r *d, c int) {
	r.didCopy(r.c)
	copy(l.d[l.c:], r.d[:c])
	copy(r.d[:], r.d[c:r.c])
	l.c += c
	r.c -= c
}

func (l *d) mvR(r *d, c int) {
	l.didCopy(r.c + c)
	copy(r.d[c:], r.d[:r.c])
	copy(r.d[:c], l.d[l.c-c:])
	r.c += c
	l.c -= c
}

// ----------------------------------------------------------------------- Tree

// TreeNew returns a newly created, empty Tree. The compare function is used
// for key collation.
func TreeNew(cmp Cmp) *Tree {
	return btTPool.get(cmp)
}

// Clear removes all K/V pairs from the tree.
func (t *Tree) Clear() {
	if t.r == nil {
		return
	}

	clr(t.r)
	t.c, t.first, t.last, t.r = 0, nil, nil, nil
	t.ver++
}

// Close performs Clear and recycles t to a pool for possible later reuse. No
// references to t should exist or such references must not be used afterwards.
func (t *Tree) Close() {
	t.Clear()
	*t = zt
	btTPool.Put(t)
}

func (t *Tree) cat(p *x, q, r *d, pi int) {
	t.ver++
	q.mvL(r, r.c)
	if r.n != nil {
		r.n.p = q
	} else {
		t.last = q
	}
	q.n = r.n
	*r = zd
	btDPool.Put(r)
	if p.c > 1 {
		p.extract(pi)
		p.x[pi].ch = q
		return
	}

	switch x := t.r.(type) {
	case *x:
		*x = zx
		btXPool.Put(x)
	case *d:
		*x = zd
		btDPool.Put(x)
	}
	t.r = q
}

func (t *Tree) catX(p, q, r *x, pi int) {
	t.ver++
	q.x[q.c].k = p.x[pi].k
	copy(q.x[q.c+1:], r.x[:r.c])
	q.c += r.c + 1
	q.x[q.c].ch = r.x[r.c].ch
	*r = zx
	btXPool.Put(r)
	if p.c > 1 {
		p.c--
		pc := p.c
		if pi < pc {
			p.x[pi].k = p.x[pi+1].k
			copy(p.x[pi+1:], p.x[pi+2:pc+1])
			p.x[pc].ch = p.x[pc+1].ch
			p.x[pc].k = zk     // GC
			p.x[pc+1].ch = nil // GC
		}
		return
	}

	switch x := t.r.(type) {
	case *x:
		*x = zx
		btXPool.Put(x)
	case *d:
		*x = zd
		btDPool.Put(x)
	}
	t.r = q
}

// Delete removes the k's KV pair, if it exists, in which case Delete returns
// true.
func (t *Tree) Delete(k bench.Key) (ok bool) {
	pi := -1
	var p *x
	q := t.r
	if q == nil {
		return false
	}

	for {
		var i int
		i, ok = t.find(q, k)
		if ok {
			switch x := q.(type) {
			case *x:
				if x.c < kx && q != t.r {
					x, i = t.underflowX(p, x, pi, i)
				}
				pi = i + 1
				p = x
				q = x.x[pi].ch
				continue
			case *d:
				t.extract(x, i)
				if x.c >= kd {
					return true
				}

				if q != t.r {
					t.underflow(p, x, pi)
				} else if t.c == 0 {
					t.Clear()
				}
				return true
			}
		}

		switch x := q.(type) {
		case *x:
			if x.c < kx && q != t.r {
				x, i = t.underflowX(p, x, pi, i)
			}
			pi = i
			p = x
			q = x.x[i].ch
		case *d:
			return false
		}
	}
}

func (t *Tree) extract(q *d, i int) { // (r bench.Value) {
	t.ver++
	//r = q.d[i].v // prepared for Extract
	q.c--
	if i < q.c {
		t.didCopy(q.c - i)
		copy(q.d[i:], q.d[i+1:q.c+1])
	}
	q.d[q.c] = zde // GC
	t.c--
}

func (t *Tree) find(q interface{}, k bench.Key) (i int, ok bool) {
	var mk bench.Key
	l := 0
	switch x := q.(type) {
	case *x:
		h := x.c - 1
		for l <= h {
			m := (l + h) >> 1
			mk = x.x[m].k
			switch cmp := t.cmp(k, mk); {
			case cmp > 0:
				l = m + 1
			case cmp == 0:
				return m, true
			default:
				h = m - 1
			}
		}
	case *d:
		h := x.c - 1
		for l <= h {
			m := (l + h) >> 1
			mk = x.d[m].k
			switch cmp := t.cmp(k, mk); {
			case cmp > 0:
				l = m + 1
			case cmp == 0:
				return m, true
			default:
				h = m - 1
			}
		}
	}
	return l, false
}

// First returns the first item of the tree in the key collating order, or
// (zero-value, zero-value) if the tree is empty.
func (t *Tree) First() (k bench.Key, v bench.Value) {
	if q := t.first; q != nil {
		q := &q.d[0]
		k, v = q.k, q.v
	}
	return
}

// Get returns the value associated with k and true if it exists. Otherwise Get
// returns (zero-value, false).
func (t *Tree) Get(k bench.Key) (v bench.Value, ok bool) {
	q := t.r
	if q == nil {
		return
	}

	for {
		var i int
		if i, ok = t.find(q, k); ok {
			switch x := q.(type) {
			case *x:
				q = x.x[i+1].ch
				continue
			case *d:
				return x.d[i].v, true
			}
		}
		switch x := q.(type) {
		case *x:
			q = x.x[i].ch
		default:
			return
		}
	}
}

func (t *Tree) insert(q *d, i int, k bench.Key, v bench.Value) *d {
	t.ver++
	q.setTree(t)
	c := q.c
	if i < c {
		t.didCopy(c - i)
		copy(q.d[i+1:], q.d[i:c])
	}
	c++
	q.c = c
	q.d[i].k, q.d[i].v = k, v
	t.c++
	return q
}

// Last returns the last item of the tree in the key collating order, or
// (zero-value, zero-value) if the tree is empty.
func (t *Tree) Last() (k bench.Key, v bench.Value) {
	if q := t.last; q != nil {
		q := &q.d[q.c-1]
		k, v = q.k, q.v
	}
	return
}

// Len returns the number of items in the tree.
func (t *Tree) Len() int {
	return t.c
}

func (t *Tree) overflow(p *x, q *d, pi, i int, k bench.Key, v bench.Value) {
	t.ver++
	l, r := p.siblings(pi)

	// s is the number of items to shift out of the full data container to
	// allow for the new data item. This logic shifts by half the available
	// space plus one. In the case where the new item is to be inserted within
	// the calculated shift space, s is reduced to include only the
	// data items up to the index of the new data item.
	//
	// This is more useful as kd gets larger, but only benefits inserts. It
	// has no effect on gets, but may marginally affect Next performance
	// (but not Prev). Omitted the 1e3/1.4 cases because they had >10%
	// variance. The performance differences may be more noticeable if you
	// have a non-interface key type, because that will generally be much
	// faster overall, making the copies are a bigger part of the workload.

	// kd=32
	// name               old time/op  new time/op  delta
	// DeCopies/Linear-8   281ns ± 4%   264ns ± 7%   -5.99%  (p=0.032 n=5+5)
	// DeCopies/Random-8  1.08µs ± 7%  1.07µs ± 5%     ~     (p=0.841 n=5+5)
	// SetSeq1e5-8        19.9ms ± 2%  17.9ms ± 5%   -9.81%  (p=0.008 n=5+5)
	// SetSeq1e6-8         230ms ± 2%   211ms ± 7%   -8.52%  (p=0.016 n=5+5)
	// SetRnd1e5-8        44.7ms ± 1%  41.4ms ± 4%   -7.51%  (p=0.008 n=5+5)
	// SetRnd1e6-8         827ms ± 2%   762ms ± 2%   -7.92%  (p=0.008 n=5+5)
	// Next1e5-8           774µs ± 2%   763µs ± 1%     ~     (p=0.222 n=5+5)
	// Next1e6-8          7.89ms ± 2%  7.77ms ± 2%   -1.54%  (p=0.008 n=5+5)

	// kd=128
	// name               old time/op  new time/op  delta
	// DeCopies/Linear-8   304ns ± 2%   266ns ± 7%  -12.31%  (p=0.008 n=5+5)
	// DeCopies/Random-8  1.28µs ± 2%  1.13µs ± 5%  -11.11%  (p=0.008 n=5+5)
	// SetSeq1e5-8        22.9ms ± 3%  18.1ms ± 9%  -20.98%  (p=0.008 n=5+5)
	// SetSeq1e6-8         257ms ± 2%   201ms ± 3%  -21.81%  (p=0.008 n=5+5)
	// SetRnd1e5-8        56.0ms ± 0%  49.6ms ± 3%  -11.41%  (p=0.008 n=5+5)
	// SetRnd1e6-8         1.13s ± 3%   1.04s ± 8%   -7.62%  (p=0.032 n=5+5)
	// Next1e5-8           714µs ± 2%   755µs ± 2%   +5.77%  (p=0.008 n=5+5)
	// Next1e6-8          7.48ms ± 1%  7.85ms ± 2%   +4.95%  (p=0.008 n=5+5)

	// kd=256
	// name               old time/op  new time/op  delta
	// DeCopies/Linear-8   350ns ± 1%   255ns ± 7%  -27.16%  (p=0.008 n=5+5)
	// DeCopies/Random-8  1.61µs ± 1%  1.32µs ± 7%  -18.28%  (p=0.008 n=5+5)
	// SetSeq1e5-8        27.7ms ± 2%  17.6ms ± 5%  -36.60%  (p=0.008 n=5+5)
	// SetSeq1e6-8         302ms ± 2%   196ms ± 1%  -35.21%  (p=0.008 n=5+5)
	// SetRnd1e5-8        74.9ms ± 2%  58.9ms ± 2%  -21.43%  (p=0.008 n=5+5)
	// SetRnd1e6-8         1.43s ± 0%   1.19s ± 2%  -16.45%  (p=0.008 n=5+5)
	// Next1e5-8           717µs ± 2%   754µs ± 2%   +5.21%  (p=0.008 n=5+5)
	// Next1e6-8          7.40ms ± 2%  7.77ms ± 1%   +4.99%  (p=0.008 n=5+5)

	if l != nil && l.c < 2*kd && i != 0 {
		s := (2*kd-l.c)/2 + 1 // half plus one
		if i < s {
			s = i
		}
		l.mvL(q, s)
		t.insert(q, i-s, k, v)
		p.x[pi-1].k = q.d[0].k
		return
	}

	if r != nil && r.c < 2*kd {
		if i < 2*kd {
			s := (2*kd-r.c)/2 + 1 // half plus one
			if 2*kd-i < s {
				s = 2*kd - i
			}
			q.mvR(r, s)
			t.insert(q, i, k, v)
			p.x[pi].k = r.d[0].k
			return
		}

		t.insert(r, 0, k, v)
		p.x[pi].k = k
		return
	}

	t.split(p, q, pi, i, k, v)
}

// Seek returns an Enumerator positioned on an item such that k >= item's key.
// ok reports if k == item.key The Enumerator's position is possibly after the
// last item in the tree.
func (t *Tree) Seek(k bench.Key) (e *Enumerator, ok bool) {
	q := t.r
	if q == nil {
		e = btEPool.get(nil, false, 0, k, nil, t, t.ver)
		return
	}

	for {
		var i int
		if i, ok = t.find(q, k); ok {
			switch x := q.(type) {
			case *x:
				q = x.x[i+1].ch
				continue
			case *d:
				return btEPool.get(nil, ok, i, k, x, t, t.ver), true
			}
		}

		switch x := q.(type) {
		case *x:
			q = x.x[i].ch
		case *d:
			return btEPool.get(nil, ok, i, k, x, t, t.ver), false
		}
	}
}

// SeekFirst returns an enumerator positioned on the first KV pair in the tree,
// if any. For an empty tree, err == io.EOF is returned and e will be nil.
func (t *Tree) SeekFirst() (e *Enumerator, err error) {
	q := t.first
	if q == nil {
		return nil, io.EOF
	}

	return btEPool.get(nil, true, 0, q.d[0].k, q, t, t.ver), nil
}

// SeekLast returns an enumerator positioned on the last KV pair in the tree,
// if any. For an empty tree, err == io.EOF is returned and e will be nil.
func (t *Tree) SeekLast() (e *Enumerator, err error) {
	q := t.last
	if q == nil {
		return nil, io.EOF
	}

	return btEPool.get(nil, true, q.c-1, q.d[q.c-1].k, q, t, t.ver), nil
}

// Set sets the value associated with k.
func (t *Tree) Set(k bench.Key, v bench.Value) {
	//dbg("--- PRE Set(%v, %v)\n%s", k, v, t.dump())
	//defer func() {
	//	dbg("--- POST\n%s\n====\n", t.dump())
	//}()

	pi := -1
	var p *x
	q := t.r
	if q == nil {
		z := t.insert(btDPool.Get().(*d), 0, k, v)
		t.r, t.first, t.last = z, z, z
		return
	}

	for {
		i, ok := t.find(q, k)
		if ok {
			switch x := q.(type) {
			case *x:
				i++
				if x.c > 2*kx {
					x, i = t.splitX(p, x, pi, i)
				}
				pi = i
				p = x
				q = x.x[i].ch
				continue
			case *d:
				x.d[i].v = v
			}
			return
		}

		switch x := q.(type) {
		case *x:
			if x.c > 2*kx {
				x, i = t.splitX(p, x, pi, i)
			}
			pi = i
			p = x
			q = x.x[i].ch
		case *d:
			switch {
			case x.c < 2*kd:
				t.insert(x, i, k, v)
			default:
				t.overflow(p, x, pi, i, k, v)
			}
			return
		}
	}
}

// Put combines Get and Set in a more efficient way where the tree is walked
// only once. The upd(ater) receives (old-value, true) if a KV pair for k
// exists or (zero-value, false) otherwise. It can then return a (new-value,
// true) to create or overwrite the existing value in the KV pair, or
// (whatever, false) if it decides not to create or not to update the value of
// the KV pair.
//
// 	tree.Set(k, v) call conceptually equals calling
//
// 	tree.Put(k, func(bench.Key, bool){ return v, true })
//
// modulo the differing return values.
func (t *Tree) Put(k bench.Key, upd func(oldV bench.Value, exists bool) (newV bench.Value, write bool)) (oldV bench.Value, written bool) {
	pi := -1
	var p *x
	q := t.r
	var newV bench.Value
	if q == nil {
		// new KV pair in empty tree
		newV, written = upd(newV, false)
		if !written {
			return
		}

		z := t.insert(btDPool.Get().(*d), 0, k, newV)
		t.r, t.first, t.last = z, z, z
		return
	}

	for {
		i, ok := t.find(q, k)
		if ok {
			switch x := q.(type) {
			case *x:
				i++
				if x.c > 2*kx {
					x, i = t.splitX(p, x, pi, i)
				}
				pi = i
				p = x
				q = x.x[i].ch
				continue
			case *d:
				oldV = x.d[i].v
				newV, written = upd(oldV, true)
				if !written {
					return
				}

				x.d[i].v = newV
			}
			return
		}

		switch x := q.(type) {
		case *x:
			if x.c > 2*kx {
				x, i = t.splitX(p, x, pi, i)
			}
			pi = i
			p = x
			q = x.x[i].ch
		case *d: // new KV pair
			newV, written = upd(newV, false)
			if !written {
				return
			}

			switch {
			case x.c < 2*kd:
				t.insert(x, i, k, newV)
			default:
				t.overflow(p, x, pi, i, k, newV)
			}
			return
		}
	}
}

func (t *Tree) split(p *x, q *d, pi, i int, k bench.Key, v bench.Value) {
	t.ver++
	r := btDPool.Get().(*d)
	r.setTree(t)
	if q.n != nil {
		r.n = q.n
		r.n.p = r
	} else {
		t.last = r
	}
	q.n = r
	r.p = q

	t.didCopy(kd)
	copy(r.d[:], q.d[kd:2*kd])
	for i := range q.d[kd:] {
		q.d[kd+i] = zde
	}
	q.c = kd
	r.c = kd
	var done bool
	if i > kd {
		done = true
		t.insert(r, i-kd, k, v)
	}
	if pi >= 0 {
		p.insert(pi, r.d[0].k, r)
	} else {
		t.r = newX(q).insert(0, r.d[0].k, r)
	}
	if done {
		return
	}

	t.insert(q, i, k, v)
}

func (t *Tree) splitX(p *x, q *x, pi int, i int) (*x, int) {
	t.ver++
	r := btXPool.Get().(*x)
	copy(r.x[:], q.x[kx+1:])
	q.c = kx
	r.c = kx
	if pi >= 0 {
		p.insert(pi, q.x[kx].k, r)
	} else {
		t.r = newX(q).insert(0, q.x[kx].k, r)
	}

	q.x[kx].k = zk
	for i := range q.x[kx+1:] {
		q.x[kx+i+1] = zxe
	}
	if i > kx {
		q = r
		i -= kx + 1
	}

	return q, i
}

func (t *Tree) underflow(p *x, q *d, pi int) {
	t.ver++
	l, r := p.siblings(pi)

	if l != nil && l.c+q.c >= 2*kd {
		l.mvR(q, 1)
		p.x[pi-1].k = q.d[0].k
		return
	}

	if r != nil && q.c+r.c >= 2*kd {
		q.mvL(r, 1)
		p.x[pi].k = r.d[0].k
		r.d[r.c] = zde // GC
		return
	}

	if l != nil {
		t.cat(p, l, q, pi-1)
		return
	}

	t.cat(p, q, r, pi)
}

func (t *Tree) underflowX(p *x, q *x, pi int, i int) (*x, int) {
	t.ver++
	var l, r *x

	if pi >= 0 {
		if pi > 0 {
			l = p.x[pi-1].ch.(*x)
		}
		if pi < p.c {
			r = p.x[pi+1].ch.(*x)
		}
	}

	if l != nil && l.c > kx {
		q.x[q.c+1].ch = q.x[q.c].ch
		copy(q.x[1:], q.x[:q.c])
		q.x[0].ch = l.x[l.c].ch
		q.x[0].k = p.x[pi-1].k
		q.c++
		i++
		l.c--
		p.x[pi-1].k = l.x[l.c].k
		return q, i
	}

	if r != nil && r.c > kx {
		q.x[q.c].k = p.x[pi].k
		q.c++
		q.x[q.c].ch = r.x[0].ch
		p.x[pi].k = r.x[0].k
		copy(r.x[:], r.x[1:r.c])
		r.c--
		rc := r.c
		r.x[rc].ch = r.x[rc+1].ch
		r.x[rc].k = zk
		r.x[rc+1].ch = nil
		return q, i
	}

	if l != nil {
		i += l.c + 1
		t.catX(p, l, q, pi-1)
		q = l
		return q, i
	}

	t.catX(p, q, r, pi)
	return q, i
}

// ----------------------------------------------------------------- Enumerator

// Close recycles e to a pool for possible later reuse. No references to e
// should exist or such references must not be used afterwards.
func (e *Enumerator) Close() {
	*e = ze
	btEPool.Put(e)
}

// Next returns the currently enumerated item, if it exists and moves to the
// next item in the key collation order. If there is no item to return, err ==
// io.EOF is returned.
func (e *Enumerator) Next() (k bench.Key, v bench.Value, err error) {
	if err = e.err; err != nil {
		return
	}

	if e.ver != e.t.ver {
		f, _ := e.t.Seek(e.k)
		*e = *f
		f.Close()
	}
	if e.q == nil {
		e.err, err = io.EOF, io.EOF
		return
	}

	if e.i >= e.q.c {
		if err = e.next(); err != nil {
			return
		}
	}

	i := e.q.d[e.i]
	k, v = i.k, i.v
	e.k, e.hit = k, true
	e.next()
	return
}

func (e *Enumerator) next() error {
	if e.q == nil {
		e.err = io.EOF
		return io.EOF
	}

	switch {
	case e.i < e.q.c-1:
		e.i++
	default:
		if e.q, e.i = e.q.n, 0; e.q == nil {
			e.err = io.EOF
		}
	}
	return e.err
}

// Prev returns the currently enumerated item, if it exists and moves to the
// previous item in the key collation order. If there is no item to return, err
// == io.EOF is returned.
func (e *Enumerator) Prev() (k bench.Key, v bench.Value, err error) {
	if err = e.err; err != nil {
		return
	}

	if e.ver != e.t.ver {
		f, _ := e.t.Seek(e.k)
		*e = *f
		f.Close()
	}
	if e.q == nil {
		e.err, err = io.EOF, io.EOF
		return
	}

	if !e.hit {
		// move to previous because Seek overshoots if there's no hit
		if err = e.prev(); err != nil {
			return
		}
	}

	if e.i >= e.q.c {
		if err = e.prev(); err != nil {
			return
		}
	}

	i := e.q.d[e.i]
	k, v = i.k, i.v
	e.k, e.hit = k, true
	e.prev()
	return
}

func (e *Enumerator) prev() error {
	if e.q == nil {
		e.err = io.EOF
		return io.EOF
	}

	switch {
	case e.i > 0:
		e.i--
	default:
		if e.q = e.q.p; e.q == nil {
			e.err = io.EOF
			break
		}

		e.i = e.q.c - 1
	}
	return e.err
}
