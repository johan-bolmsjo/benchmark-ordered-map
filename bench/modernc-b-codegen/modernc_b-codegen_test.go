package moderncBCodegen_test

import (
	"io"
	"testing"

	btree "modernc.org/benchmark-ordered-map/bench/modernc-b-codegen/gen"
	"modernc.org/benchmark-ordered-map/fixture"
)

func moderncCmp(a, b fixture.Key) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return +1
	default:
		return 0
	}
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := btree.TreeNew(moderncCmp)
		for i := 0; i < len(fixture.TestData); i++ {
			tree.Set(fixture.TestData[i].Key, fixture.TestData[i].Value)
		}
	}
}

func BenchmarkSortedInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := btree.TreeNew(moderncCmp)
		for i := 0; i < len(fixture.SortedTestData); i++ {
			tree.Set(fixture.SortedTestData[i].Key, fixture.SortedTestData[i].Value)
		}
	}
}

var (
	ksink fixture.Key
	vsink fixture.Value
)

func BenchmarkIterate(b *testing.B) {
	tree := btree.TreeNew(moderncCmp)
	for i := 0; i < len(fixture.TestData); i++ {
		tree.Set(fixture.TestData[i].Key, fixture.TestData[i].Value)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// only errors on empty trees; meh api
		iter, err := tree.SeekFirst()
		if err != nil {
			b.Fatalf("tree.SeekFirst: %v", err)
		}
		for {
			ksink, vsink, err = iter.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				b.Fatalf("iter.Next: %v", err)
			}
		}
		iter.Close()
	}
}
