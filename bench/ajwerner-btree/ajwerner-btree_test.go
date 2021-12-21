package ajwerner_btree_test

import (
	"testing"

	btree "github.com/ajwerner/btree"
	"modernc.org/benchmark-ordered-map/fixture"
)

func cmp(a, b fixture.Key) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return +1
	default:
		return 0
	}
}

func new() btree.Map[fixture.Key, fixture.Value] {
	return btree.MakeMap[fixture.Key, fixture.Value](cmp)
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := new()
		for i := 0; i < len(fixture.TestData); i++ {
			tree.Upsert(fixture.TestData[i].Key, fixture.TestData[i].Value)
		}
	}
}

func BenchmarkSortedInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := new()
		for i := 0; i < len(fixture.SortedTestData); i++ {
			tree.Upsert(fixture.SortedTestData[i].Key, fixture.SortedTestData[i].Value)
		}
	}
}

var (
	ksink fixture.Key
	vsink fixture.Value
)

func BenchmarkIterate(b *testing.B) {
	tree := new()
	for i := 0; i < len(fixture.TestData); i++ {
		tree.Upsert(fixture.TestData[i].Key, fixture.TestData[i].Value)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		iter := tree.Iterator()
		for iter.First(); iter.Valid(); iter.Next() {
			ksink = iter.Cur()
			vsink = iter.Value()
		}
	}
}
