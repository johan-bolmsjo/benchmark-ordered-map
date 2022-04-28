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
- [github.com/Michal-Leszczynski/btree](https://github.com/Michal-Leszczynski/btree)
- [github.com/ajwerner/btree](https://github.com/ajwerner/btree)
- [github.com/biogo/store/llrb](https://github.com/biogo/store)
- [github.com/glenn-brown/skiplist](https://github.com/glenn-brown/skiplist)
- [github.com/google/btree](https://github.com/google/btree)
- [github.com/kellydunn/go-art](https://github.com/kellydunn/go-art)
- [github.com/petar/GoLLRB/llrb](https://github.com/petar/GoLLRB)
- [github.com/ryszard/goskiplist/skiplist](https://github.com/ryszard/goskiplist)
- [modernc.org/b/v2](https://gitlab.com/cznic/b/v2)
- [modernc.org/b](https://gitlab.com/cznic/b)

Do you want to see your library results here? Send me a merge request.

# Results

## Insert
```
benchmark                       iter       time/iter     bytes alloc             allocs
---------                       ----       ---------     -----------             ------
Benchmark_modernc-b-codegen       73     16.21 ms/op    3292148 B/op     1840 allocs/op
Benchmark_modernc-b-generic       73     16.46 ms/op    3292441 B/op     1843 allocs/op
Benchmark_ajwerner-btree          69     17.21 ms/op    3430521 B/op     1112 allocs/op
Benchmark_michal-leszczynski      56     20.79 ms/op    5322432 B/op     6918 allocs/op
Benchmark_google-btree            38     30.22 ms/op    6012120 B/op   106918 allocs/op
Benchmark_modernc-b               33     33.04 ms/op    6632815 B/op   201842 allocs/op
Benchmark_biogo-llrb              26     40.59 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB            24     42.98 ms/op    7200000 B/op   200000 allocs/op
Benchmark_kellydunn-art           22     47.79 ms/op   28746952 B/op   800744 allocs/op
Benchmark_3xian-elephantlist      21     49.00 ms/op    9628982 B/op   227176 allocs/op
Benchmark_glennbrown-skiplist     18     59.76 ms/op   12016787 B/op   400017 allocs/op
Benchmark_ryszard-skiplist         8    144.13 ms/op   66400258 B/op   500005 allocs/op
Benchmark_sortedslice              1   2350.67 ms/op   14077952 B/op       30 allocs/op
```

## Iterate
```
benchmark                        iter         time/iter   bytes alloc        allocs
---------                        ----         ---------   -----------        ------
Benchmark_sortedslice           12259       97.71 μs/op        0 B/op   0 allocs/op
Benchmark_ajwerner-btree         3872      312.32 μs/op        0 B/op   0 allocs/op
Benchmark_michal-leszczynski     3561      338.07 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b-generic      3042      393.25 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b-codegen      3117      395.25 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b              2017      619.37 μs/op        0 B/op   0 allocs/op
Benchmark_google-btree           1905      625.50 μs/op        0 B/op   0 allocs/op
Benchmark_biogo-llrb             1028     1211.41 μs/op        0 B/op   0 allocs/op
Benchmark_petar-GoLLRB            778     1534.41 μs/op        0 B/op   0 allocs/op
Benchmark_glennbrown-skiplist     388     3009.81 μs/op        0 B/op   0 allocs/op
Benchmark_kellydunn-art           170     7617.26 μs/op        0 B/op   0 allocs/op
Benchmark_ryszard-skiplist         76    14288.58 μs/op        0 B/op   0 allocs/op
```

## SortedInsert
```
benchmark                                      iter     time/iter     bytes alloc             allocs
---------                                      ----     ---------     -----------             ------
Benchmark_sortedslice                           211    5.61 ms/op   14077952 B/op       30 allocs/op
Benchmark_modernc-b-generic                     211    5.71 ms/op    2887452 B/op     1617 allocs/op
Benchmark_modernc-b-codegen                     210    5.85 ms/op    2887115 B/op     1613 allocs/op
Benchmark_ajwerner-btree                        194    6.10 ms/op    4901002 B/op     1588 allocs/op
Benchmark_michal-leszczynski_ReplaceOrInsert     93   13.70 ms/op    7713720 B/op     9904 allocs/op
Benchmark_modernc-b                              80   15.47 ms/op    6112100 B/op   201615 allocs/op
Benchmark_google-btree_ReplaceOrInsert           76   16.60 ms/op    7637176 B/op   109904 allocs/op
Benchmark_biogo-llrb_InsertNoReplace             48   24.66 ms/op    7200000 B/op   200000 allocs/op
Benchmark_biogo-llrb_ReplaceOrInsert             48   25.93 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB_InsertNoReplace           38   26.63 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB_ReplaceOrInsert           38   29.90 ms/op    7200000 B/op   200000 allocs/op
Benchmark_3xian-elephantlist                     38   33.11 ms/op   12562089 B/op   237508 allocs/op
Benchmark_glennbrown-skiplist                    37   33.95 ms/op   12016788 B/op   400017 allocs/op
Benchmark_kellydunn-art                          34   39.63 ms/op   28737712 B/op   799621 allocs/op
Benchmark_ryszard-skiplist                       18   68.32 ms/op   66400298 B/op   500005 allocs/op
```
