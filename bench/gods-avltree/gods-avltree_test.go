package godsAvltree_test

import (
	"testing"

	"github.com/johan-bolmsjo/gods/v3/avltree"
	"github.com/johan-bolmsjo/gods/v3/math"
	"modernc.org/benchmark-ordered-map/fixture"
)

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := avltree.New[fixture.Key, fixture.Value](math.CompareOrdered[fixture.Key])
		for i := 0; i < len(fixture.TestData); i++ {
			tree.Add(fixture.TestData[i].Key, fixture.TestData[i].Value)
		}
	}
}

func BenchmarkSortedInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := avltree.New[fixture.Key, fixture.Value](math.CompareOrdered[fixture.Key])
		for i := 0; i < len(fixture.SortedTestData); i++ {
			tree.Add(fixture.TestData[i].Key, fixture.TestData[i].Value)
		}
	}
}

var (
	ksink fixture.Key
	vsink fixture.Value
)

func BenchmarkIterate(b *testing.B) {
	tree := avltree.New[fixture.Key, fixture.Value](math.CompareOrdered[fixture.Key])
	for i := 0; i < len(fixture.TestData); i++ {
		tree.Add(fixture.TestData[i].Key, fixture.TestData[i].Value)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for k, v := range tree.All() {
			ksink = k
			vsink = v
		}
	}
}
