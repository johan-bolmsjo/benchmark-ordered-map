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

## Participating libraries

[github.com/3xian/elephantlist](https://github.com/3xian/elephantlist)
[github.com/biogo/store/llrb](hppts://github.com/biogo/store/llrb)
[github.com/glenn-brown/skiplist](https://github.com/glenn-brown/skiplist)
[github.com/google/btree](https://github.com/google/btree)
[github.com/kellydunn/go-art](https://github.com/kellydunn/go-art)
[github.com/petar/GoLLRB/llrb](https://github.com/petar/GoLLRB/llrb)
[github.com/ryszard/goskiplist/skiplist](https://github.com/ryszard/goskiplist/skiplist)
[modernc.org/b](https://modernc.org/b)

Do you want to see your library results here? Send me a merge request.

# Results

## Insert
```
benchmark                       iter       time/iter     bytes alloc             allocs
---------                       ----       ---------     -----------             ------
Benchmark_modernc-b-codegen       74     15.79 ms/op    3292154 B/op     1840 allocs/op
Benchmark_google-btree            38     31.18 ms/op    6012120 B/op   106918 allocs/op
Benchmark_modernc-b               32     33.10 ms/op    6632856 B/op   201842 allocs/op
Benchmark_biogo-llrb              25     42.91 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB            25     43.93 ms/op    7200000 B/op   200000 allocs/op
Benchmark_3xian-elephantlist      21     49.17 ms/op    9628983 B/op   227176 allocs/op
Benchmark_kellydunn-art           25     49.22 ms/op   28746952 B/op   800744 allocs/op
Benchmark_glennbrown-skiplist     19     59.38 ms/op   12016787 B/op   400017 allocs/op
Benchmark_ryszard-skiplist         7    156.70 ms/op   66400298 B/op   500005 allocs/op
Benchmark_sortedslice              1   2379.49 ms/op   14077952 B/op       30 allocs/op
```

## Iterate
```
benchmark                        iter         time/iter   bytes alloc        allocs
---------                        ----         ---------   -----------        ------
Benchmark_sortedslice           50992       23.49 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b-codegen      3350      358.78 μs/op        0 B/op   0 allocs/op
Benchmark_google-btree           3165      373.26 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b              2869      432.62 μs/op        0 B/op   0 allocs/op
Benchmark_biogo-llrb             1248      978.60 μs/op        0 B/op   0 allocs/op
Benchmark_petar-GoLLRB            798     1497.52 μs/op        0 B/op   0 allocs/op
Benchmark_glennbrown-skiplist     510     2335.80 μs/op        0 B/op   0 allocs/op
Benchmark_kellydunn-art           262     4452.33 μs/op        0 B/op   0 allocs/op
Benchmark_ryszard-skiplist         96    12261.36 μs/op        0 B/op   0 allocs/op
```

## SortedInsert
```
benchmark                                iter     time/iter     bytes alloc             allocs
---------                                ----     ---------     -----------             ------
Benchmark_modernc-b-codegen               207    5.58 ms/op    2887109 B/op     1613 allocs/op
Benchmark_sortedslice                     196    6.12 ms/op   14077952 B/op       30 allocs/op
Benchmark_modernc-b                        75   16.39 ms/op    6112124 B/op   201616 allocs/op
Benchmark_google-btree_ReplaceOrInsert     73   18.06 ms/op    7637176 B/op   109904 allocs/op
Benchmark_biogo-llrb_InsertNoReplace       44   23.81 ms/op    7200000 B/op   200000 allocs/op
Benchmark_biogo-llrb_ReplaceOrInsert       52   24.00 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB_InsertNoReplace     45   25.23 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB_ReplaceOrInsert     38   28.68 ms/op    7200000 B/op   200000 allocs/op
Benchmark_3xian-elephantlist               38   32.70 ms/op   12562090 B/op   237508 allocs/op
Benchmark_glennbrown-skiplist              38   33.43 ms/op   12016787 B/op   400017 allocs/op
Benchmark_kellydunn-art                    34   39.92 ms/op   28737712 B/op   799621 allocs/op
Benchmark_ryszard-skiplist                 18   67.67 ms/op   66400286 B/op   500005 allocs/op
```
