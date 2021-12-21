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
Benchmark_modernc-b-generic       69     15.72 ms/op    3292461 B/op     1843 allocs/op
Benchmark_modernc-b-codegen       74     15.85 ms/op    3292152 B/op     1840 allocs/op
Benchmark_ajwerner-btree          73     16.28 ms/op    3430515 B/op     1112 allocs/op
Benchmark_google-btree            38     30.04 ms/op    6012120 B/op   106918 allocs/op
Benchmark_modernc-b               34     33.54 ms/op    6632821 B/op   201842 allocs/op
Benchmark_biogo-llrb              25     41.12 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB            24     42.73 ms/op    7200000 B/op   200000 allocs/op
Benchmark_kellydunn-art           25     47.80 ms/op   28746952 B/op   800744 allocs/op
Benchmark_3xian-elephantlist      22     48.21 ms/op    9628990 B/op   227176 allocs/op
Benchmark_glennbrown-skiplist     19     58.64 ms/op   12016787 B/op   400017 allocs/op
Benchmark_ryszard-skiplist         7    148.91 ms/op   66400248 B/op   500005 allocs/op
Benchmark_sortedslice              1   2360.89 ms/op   14077952 B/op       30 allocs/op
```

## Iterate
```
benchmark                        iter         time/iter   bytes alloc        allocs
---------                        ----         ---------   -----------        ------
Benchmark_sortedslice           12186       98.64 μs/op        0 B/op   0 allocs/op
Benchmark_ajwerner-btree         3830      316.18 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b-generic      2989      382.33 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b-codegen      3098      390.90 μs/op        0 B/op   0 allocs/op
Benchmark_google-btree           2216      548.08 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b              1861      665.92 μs/op        0 B/op   0 allocs/op
Benchmark_biogo-llrb             1069     1144.55 μs/op        0 B/op   0 allocs/op
Benchmark_petar-GoLLRB            745     1627.73 μs/op        0 B/op   0 allocs/op
Benchmark_glennbrown-skiplist     424     2843.59 μs/op        0 B/op   0 allocs/op
Benchmark_kellydunn-art           188     6497.36 μs/op        0 B/op   0 allocs/op
Benchmark_ryszard-skiplist         85    14012.93 μs/op        0 B/op   0 allocs/op
```

## SortedInsert
```
benchmark                                iter     time/iter     bytes alloc             allocs
---------                                ----     ---------     -----------             ------
Benchmark_modernc-b-codegen               217    5.50 ms/op    2887117 B/op     1613 allocs/op
Benchmark_modernc-b-generic               208    5.66 ms/op    2887448 B/op     1617 allocs/op
Benchmark_ajwerner-btree                  200    6.07 ms/op    4900997 B/op     1588 allocs/op
Benchmark_sortedslice                     196    6.13 ms/op   14077952 B/op       30 allocs/op
Benchmark_modernc-b                        72   16.64 ms/op    6112096 B/op   201615 allocs/op
Benchmark_google-btree_ReplaceOrInsert     68   18.15 ms/op    7637176 B/op   109904 allocs/op
Benchmark_biogo-llrb_ReplaceOrInsert       43   23.76 ms/op    7200000 B/op   200000 allocs/op
Benchmark_biogo-llrb_InsertNoReplace       51   24.30 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB_InsertNoReplace     44   26.16 ms/op    7200000 B/op   200000 allocs/op
Benchmark_petar-GoLLRB_ReplaceOrInsert     34   30.36 ms/op    7200000 B/op   200000 allocs/op
Benchmark_glennbrown-skiplist              33   32.14 ms/op   12016786 B/op   400017 allocs/op
Benchmark_3xian-elephantlist               40   32.25 ms/op   12562054 B/op   237508 allocs/op
Benchmark_kellydunn-art                    28   39.90 ms/op   28737712 B/op   799621 allocs/op
Benchmark_ryszard-skiplist                 20   62.04 ms/op   66400296 B/op   500005 allocs/op
```
