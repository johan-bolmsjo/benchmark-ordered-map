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

## Insert-24
```
benchmark                       iter       time/iter     bytes alloc             allocs
---------                       ----       ---------     -----------             ------
Benchmark_modernc-b-codegen       72     28.15 ms/op    3296953 B/op     1841 allocs/op
Benchmark_google-btree            32     52.94 ms/op    6012494 B/op   106918 allocs/op
Benchmark_modernc-b               19     55.44 ms/op    6640263 B/op   201843 allocs/op
Benchmark_petar-GoLLRB            16     65.03 ms/op    7200000 B/op   200000 allocs/op
Benchmark_biogo-llrb              19     67.95 ms/op    7200212 B/op   200000 allocs/op
Benchmark_kellydunn-art           16     73.13 ms/op   28747037 B/op   800745 allocs/op
Benchmark_3xian-elephantlist      15     79.14 ms/op    9629817 B/op   227177 allocs/op
Benchmark_glennbrown-skiplist     13     94.15 ms/op   12016811 B/op   400017 allocs/op
Benchmark_ryszard-skiplist         6    193.43 ms/op   66400428 B/op   500008 allocs/op
Benchmark_sortedslice              1   2409.46 ms/op   14077960 B/op       31 allocs/op
```

## Iterate-24
```
benchmark                        iter         time/iter   bytes alloc        allocs
---------                        ----         ---------   -----------        ------
Benchmark_sortedslice           51717       23.51 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b-codegen      3213      361.38 μs/op        0 B/op   0 allocs/op
Benchmark_google-btree           3099      391.15 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b              2640      437.31 μs/op        1 B/op   0 allocs/op
Benchmark_biogo-llrb             1194      989.55 μs/op        0 B/op   0 allocs/op
Benchmark_petar-GoLLRB            644     1764.26 μs/op        0 B/op   0 allocs/op
Benchmark_glennbrown-skiplist     495     2306.66 μs/op        0 B/op   0 allocs/op
Benchmark_kellydunn-art           255     4423.19 μs/op        0 B/op   0 allocs/op
Benchmark_ryszard-skiplist         96    12164.58 μs/op        0 B/op   0 allocs/op
```

## SortedInsert
```
benchmark                                   iter     time/iter    bytes alloc             allocs
---------                                   ----     ---------    -----------             ------
Benchmark_google-btree_ReplaceOrInsert-24     36   31.27 ms/op   7637182 B/op   109904 allocs/op
Benchmark_biogo-llrb_ReplaceOrInsert-24       27   39.30 ms/op   7200018 B/op   200000 allocs/op
Benchmark_biogo-llrb_InsertNoReplace-24       27   39.45 ms/op   7200015 B/op   200000 allocs/op
Benchmark_petar-GoLLRB_InsertNoReplace-24     25   42.17 ms/op   7200008 B/op   200000 allocs/op
Benchmark_petar-GoLLRB_ReplaceOrInsert-24     21   48.27 ms/op   7200016 B/op   200000 allocs/op
```

## SortedInsert-24
```
benchmark                       iter     time/iter     bytes alloc             allocs
---------                       ----     ---------     -----------             ------
Benchmark_modernc-b-codegen      100   10.40 ms/op    2890845 B/op     1614 allocs/op
Benchmark_sortedslice             99   10.98 ms/op   14078027 B/op       30 allocs/op
Benchmark_modernc-b               58   26.68 ms/op    6118799 B/op   201616 allocs/op
Benchmark_glennbrown-skiplist     32   50.00 ms/op   12016841 B/op   400017 allocs/op
Benchmark_3xian-elephantlist      20   52.69 ms/op   12562128 B/op   237508 allocs/op
Benchmark_kellydunn-art           18   62.31 ms/op   28737750 B/op   799621 allocs/op
Benchmark_ryszard-skiplist        12   95.28 ms/op   66400380 B/op   500006 allocs/op
```
