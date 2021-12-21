# Ordered map Go data structure benchmarks

## Workloads

### Insert

Inserts the test data set in the order it is given (pseudorandom),
replacing existing items.

### Iterate

Iterates all values in key order, accessing each key and value.

### SortedInsert

Inserts the test data, pre-sorted, in key order.


## Requirements for a library

  - Name the benchmark functions
    `Benchmark<Workload>[_<Alternative>]`.

  - If the API takes key and value separately, use `Key` and `Value`;
    if not, use `Item`.

  - Never store pointers to `Key`, `Value` or `Item`; we can simulate
	that by changing the size of the data types, instead.

  - Remember to loop `b.N` times.

  - When iterating, remember to assign both `Key` and `Value` to a static sink,
	after any type asserts if those are needed.

  - TODO: After the loop, `b.StopTimer()` and submit results to `Verify`.

## Participating libraries

- [github.com/3xian/elephantlist](https://github.com/3xian/elephantlist)
- [github.com/ajwerner/btree](https://github.com/ajwerner/btree)
- [github.com/biogo/store/llrb](https://github.com/biogo/store)
- [github.com/glenn-brown/skiplist](https://github.com/glenn-brown/skiplist)
- [github.com/google/btree](https://github.com/google/btree)
- [github.com/kellydunn/go-art](https://github.com/kellydunn/go-art)
- [github.com/petar/GoLLRB/llrb](https://github.com/petar/GoLLRB)
- [github.com/ryszard/goskiplist/skiplist](https://github.com/ryszard/goskiplist)
- [modernc.org/b](https://gitlab.com/cznic/b)

Do you want to see your library results here? Send me a merge request.

# Results

## Insert
```
benchmark                       iter       time/iter     bytes alloc             allocs
---------                       ----       ---------     -----------             ------
Benchmark_modernc-b-codegen       63     16.45 ms/op    3292154 B/op     1840 allocs/op
Benchmark_ajwerner-btree          72     16.78 ms/op    3430524 B/op     1112 allocs/op
Benchmark_google-btree            38     31.25 ms/op    6012120 B/op   106918 allocs/op
Benchmark_modernc-b               33     34.93 ms/op    6632812 B/op   201842 allocs/op
Benchmark_biogo-llrb              25     40.21 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB            25     42.05 ms/op    7200000 B/op   200000 allocs/op
Benchmark_3xian-elephantlist      25     47.31 ms/op    9629090 B/op   227176 allocs/op
Benchmark_kellydunn-art           21     48.30 ms/op   28746952 B/op   800744 allocs/op
Benchmark_glennbrown-skiplist     19     59.29 ms/op   12016788 B/op   400017 allocs/op
Benchmark_ryszard-skiplist         7    144.09 ms/op   66400245 B/op   500005 allocs/op
Benchmark_sortedslice              1   2343.96 ms/op   14077952 B/op       30 allocs/op
```

## Iterate
```
benchmark                        iter         time/iter   bytes alloc        allocs
---------                        ----         ---------   -----------        ------
Benchmark_sortedslice           12280       97.37 μs/op        0 B/op   0 allocs/op
Benchmark_ajwerner-btree         3924      310.46 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b-codegen      3045      397.19 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b              2030      597.72 μs/op        0 B/op   0 allocs/op
Benchmark_google-btree           1911      623.44 μs/op        0 B/op   0 allocs/op
Benchmark_biogo-llrb             1082     1112.01 μs/op        0 B/op   0 allocs/op
Benchmark_petar-GoLLRB            734     1665.02 μs/op        0 B/op   0 allocs/op
Benchmark_glennbrown-skiplist     392     3041.67 μs/op        0 B/op   0 allocs/op
Benchmark_kellydunn-art           188     6505.37 μs/op        0 B/op   0 allocs/op
Benchmark_ryszard-skiplist         86    13840.70 μs/op        0 B/op   0 allocs/op
```

## SortedInsert
```
benchmark                                iter     time/iter     bytes alloc             allocs
---------                                ----     ---------     -----------             ------
Benchmark_modernc-b-codegen               213    5.57 ms/op    2887121 B/op     1613 allocs/op
Benchmark_sortedslice                     199    6.02 ms/op   14077952 B/op       30 allocs/op
Benchmark_ajwerner-btree                  196    6.12 ms/op    4900999 B/op     1588 allocs/op
Benchmark_modernc-b                        75   16.03 ms/op    6112131 B/op   201616 allocs/op
Benchmark_google-btree_ReplaceOrInsert     69   18.12 ms/op    7637176 B/op   109904 allocs/op
Benchmark_biogo-llrb_ReplaceOrInsert       50   24.37 ms/op    7200000 B/op   200000 allocs/op
Benchmark_biogo-llrb_InsertNoReplace       43   24.50 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB_InsertNoReplace     45   25.75 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB_ReplaceOrInsert     39   29.79 ms/op    7200000 B/op   200000 allocs/op
Benchmark_3xian-elephantlist               34   32.11 ms/op   12561983 B/op   237508 allocs/op
Benchmark_glennbrown-skiplist              32   33.60 ms/op   12016787 B/op   400017 allocs/op
Benchmark_kellydunn-art                    30   39.04 ms/op   28737712 B/op   799621 allocs/op
Benchmark_ryszard-skiplist                 19   66.36 ms/op   66400274 B/op   500005 allocs/op
```
