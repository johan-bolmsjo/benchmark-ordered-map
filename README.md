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

  - When iterating, remember to assign both `Key` and `Value` to `_`,
	after any type asserts if those are needed.

  - TODO: After the loop, `b.StopTimer()` and submit results to `Verify`.

# Results

## Insert
```
benchmark                       iter       time/iter     bytes alloc             allocs
---------                       ----       ---------     -----------             ------
Benchmark_modernc-b-codegen       76     15.47 ms/op    3292153 B/op     1840 allocs/op
Benchmark_google-btree            39     29.54 ms/op    6012120 B/op   106918 allocs/op
Benchmark_modernc-b               33     33.76 ms/op    6632851 B/op   201842 allocs/op
Benchmark_biogo-llrb              30     39.26 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB            24     41.94 ms/op    7200000 B/op   200000 allocs/op
Benchmark_kellydunn-art           21     47.80 ms/op   28746952 B/op   800744 allocs/op
Benchmark_3xian-elephantlist      22     48.99 ms/op    9628991 B/op   227176 allocs/op
Benchmark_glennbrown-skiplist     19     60.08 ms/op   12016787 B/op   400017 allocs/op
Benchmark_ryszard-skiplist         8    145.87 ms/op   66400258 B/op   500005 allocs/op
Benchmark_sortedslice              1   2355.25 ms/op   14077952 B/op       30 allocs/op
```

## Iterate
```
benchmark                        iter         time/iter   bytes alloc        allocs
---------                        ----         ---------   -----------        ------
Benchmark_sortedslice           53152       22.58 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b-codegen      3310      365.27 μs/op        0 B/op   0 allocs/op
Benchmark_google-btree           3290      368.17 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b              2756      441.15 μs/op        0 B/op   0 allocs/op
Benchmark_biogo-llrb             1274      969.28 μs/op        0 B/op   0 allocs/op
Benchmark_petar-GoLLRB            729     1636.97 μs/op        0 B/op   0 allocs/op
Benchmark_glennbrown-skiplist     519     2312.26 μs/op        0 B/op   0 allocs/op
Benchmark_kellydunn-art           267     4466.77 μs/op        0 B/op   0 allocs/op
Benchmark_ryszard-skiplist         93    12881.09 μs/op        0 B/op   0 allocs/op
```

## SortedInsert
```
benchmark                                iter     time/iter     bytes alloc             allocs
---------                                ----     ---------     -----------             ------
Benchmark_modernc-b-codegen               219    5.48 ms/op    2887116 B/op     1613 allocs/op
Benchmark_sortedslice                     199    6.00 ms/op   14077952 B/op       30 allocs/op
Benchmark_modernc-b                        67   16.91 ms/op    6112105 B/op   201616 allocs/op
Benchmark_google-btree_ReplaceOrInsert     75   17.75 ms/op    7637176 B/op   109904 allocs/op
Benchmark_biogo-llrb_ReplaceOrInsert       44   23.66 ms/op    7200000 B/op   200000 allocs/op
Benchmark_biogo-llrb_InsertNoReplace       52   23.69 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB_InsertNoReplace     45   25.28 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB_ReplaceOrInsert     40   29.31 ms/op    7200000 B/op   200000 allocs/op
Benchmark_3xian-elephantlist               40   33.09 ms/op   12562052 B/op   237508 allocs/op
Benchmark_glennbrown-skiplist              38   33.43 ms/op   12016787 B/op   400017 allocs/op
Benchmark_kellydunn-art                    34   39.51 ms/op   28737712 B/op   799621 allocs/op
Benchmark_ryszard-skiplist                 19   65.88 ms/op   66400292 B/op   500005 allocs/op
```
