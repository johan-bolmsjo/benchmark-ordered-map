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
Benchmark_modernc-b-codegen       75     15.68 ms/op    3292157 B/op     1840 allocs/op
Benchmark_ajwerner-btree          70     17.02 ms/op    3659652 B/op     4689 allocs/op
Benchmark_google-btree            38     30.39 ms/op    6012120 B/op   106918 allocs/op
Benchmark_modernc-b               34     34.45 ms/op    6632838 B/op   201842 allocs/op
Benchmark_biogo-llrb              28     40.76 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB            24     42.75 ms/op    7200000 B/op   200000 allocs/op
Benchmark_3xian-elephantlist      25     47.13 ms/op    9629089 B/op   227176 allocs/op
Benchmark_kellydunn-art           22     48.60 ms/op   28746952 B/op   800744 allocs/op
Benchmark_glennbrown-skiplist     19     61.20 ms/op   12016788 B/op   400017 allocs/op
Benchmark_ryszard-skiplist         8    140.94 ms/op   66400310 B/op   500005 allocs/op
Benchmark_sortedslice              1   2339.28 ms/op   14077952 B/op       30 allocs/op
```

## Iterate
```
benchmark                        iter         time/iter   bytes alloc        allocs
---------                        ----         ---------   -----------        ------
Benchmark_sortedslice           53302       22.98 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b-codegen      3379      356.80 μs/op        0 B/op   0 allocs/op
Benchmark_google-btree           3184      378.22 μs/op        0 B/op   0 allocs/op
Benchmark_ajwerner-btree         2958      401.02 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b              2818      430.62 μs/op        0 B/op   0 allocs/op
Benchmark_biogo-llrb             1230      978.11 μs/op        0 B/op   0 allocs/op
Benchmark_petar-GoLLRB            758     1608.28 μs/op        0 B/op   0 allocs/op
Benchmark_glennbrown-skiplist     486     2362.92 μs/op        0 B/op   0 allocs/op
Benchmark_kellydunn-art           256     4583.21 μs/op        0 B/op   0 allocs/op
Benchmark_ryszard-skiplist         90    12516.47 μs/op        0 B/op   0 allocs/op
```

## SortedInsert
```
benchmark                                iter     time/iter     bytes alloc             allocs
---------                                ----     ---------     -----------             ------
Benchmark_modernc-b-codegen               219    5.41 ms/op    2887120 B/op     1613 allocs/op
Benchmark_sortedslice                     201    5.99 ms/op   14077952 B/op       30 allocs/op
Benchmark_ajwerner-btree                  156    7.67 ms/op    5225110 B/op     6667 allocs/op
Benchmark_modernc-b                        68   16.37 ms/op    6112113 B/op   201616 allocs/op
Benchmark_google-btree_ReplaceOrInsert     74   18.22 ms/op    7637176 B/op   109904 allocs/op
Benchmark_biogo-llrb_InsertNoReplace       43   24.26 ms/op    7200000 B/op   200000 allocs/op
Benchmark_biogo-llrb_ReplaceOrInsert       49   25.05 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB_InsertNoReplace     38   26.43 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB_ReplaceOrInsert     36   30.06 ms/op    7200000 B/op   200000 allocs/op
Benchmark_3xian-elephantlist               40   32.05 ms/op   12561975 B/op   237508 allocs/op
Benchmark_glennbrown-skiplist              33   34.14 ms/op   12016788 B/op   400017 allocs/op
Benchmark_kellydunn-art                    34   39.57 ms/op   28737712 B/op   799621 allocs/op
Benchmark_ryszard-skiplist                 18   65.49 ms/op   66400287 B/op   500005 allocs/op
```
