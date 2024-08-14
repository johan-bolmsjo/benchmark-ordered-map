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
- [github.com/johan-bolmsjo/gods](https://github.com/johan-bolmsjo/gods)
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
Benchmark_modernc-b-generic       49     21.40 ms/op    3298375 B/op     1845 allocs/op
Benchmark_modernc-b-codegen       51     24.02 ms/op    3295586 B/op     1841 allocs/op
Benchmark_ajwerner-btree          48     26.92 ms/op    3432930 B/op     1113 allocs/op
Benchmark_michal-leszczynski      39     32.50 ms/op    5322832 B/op     6918 allocs/op
Benchmark_gods-avltree            32     38.50 ms/op    4801620 B/op   100002 allocs/op
Benchmark_modernc-b               27     44.76 ms/op    6639328 B/op   201844 allocs/op
Benchmark_google-btree            24     49.38 ms/op    6293089 B/op   106919 allocs/op
Benchmark_kellydunn-art           22     49.41 ms/op   28780021 B/op   800744 allocs/op
Benchmark_biogo-llrb              20     53.12 ms/op    7200007 B/op   200000 allocs/op
Benchmark_petar-GoLLRB            19     56.39 ms/op    7200276 B/op   200000 allocs/op
Benchmark_glennbrown-skiplist     15     72.23 ms/op   12016821 B/op   400017 allocs/op
Benchmark_3xian-elephantlist      15     74.06 ms/op    9629110 B/op   227176 allocs/op
Benchmark_ryszard-skiplist         7    151.28 ms/op   66400426 B/op   500007 allocs/op
Benchmark_sortedslice              1   2779.18 ms/op   14077968 B/op       31 allocs/op
```

## Iterate
```
benchmark                        iter        time/iter   bytes alloc        allocs
---------                        ----        ---------   -----------        ------
Benchmark_sortedslice           10000     118.43 μs/op        0 B/op   0 allocs/op
Benchmark_ajwerner-btree         3105     382.21 μs/op        0 B/op   0 allocs/op
Benchmark_michal-leszczynski     2853     407.33 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b-codegen      2462     422.10 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b-generic      2643     458.76 μs/op        0 B/op   0 allocs/op
Benchmark_modernc-b              1485     814.41 μs/op        1 B/op   0 allocs/op
Benchmark_google-btree           1066     947.90 μs/op        0 B/op   0 allocs/op
Benchmark_gods-avltree            734    1458.60 μs/op        0 B/op   0 allocs/op
Benchmark_biogo-llrb              669    1621.20 μs/op        0 B/op   0 allocs/op
Benchmark_petar-GoLLRB            481    2323.82 μs/op        0 B/op   0 allocs/op
Benchmark_kellydunn-art           417    2740.19 μs/op        0 B/op   0 allocs/op
Benchmark_glennbrown-skiplist     302    3900.58 μs/op        0 B/op   0 allocs/op
Benchmark_ryszard-skiplist        231    4920.08 μs/op        0 B/op   0 allocs/op
```

## SortedInsert
```
benchmark                                      iter     time/iter     bytes alloc             allocs
---------                                      ----     ---------     -----------             ------
Benchmark_modernc-b-generic                     156    7.17 ms/op    2892029 B/op     1617 allocs/op
Benchmark_modernc-b-codegen                     146    8.08 ms/op    2890037 B/op     1614 allocs/op
Benchmark_ajwerner-btree                        123   10.08 ms/op    4903845 B/op     1589 allocs/op
Benchmark_sortedslice                           100   10.39 ms/op   14078118 B/op       30 allocs/op
Benchmark_modernc-b                              68   17.06 ms/op    6116779 B/op   201616 allocs/op
Benchmark_michal-leszczynski_ReplaceOrInsert     63   18.73 ms/op    7713979 B/op     9904 allocs/op
Benchmark_google-btree_ReplaceOrInsert           46   27.95 ms/op    8049773 B/op   109905 allocs/op
Benchmark_biogo-llrb_ReplaceOrInsert             39   32.39 ms/op    7200134 B/op   200000 allocs/op
Benchmark_biogo-llrb_InsertNoReplace             39   33.04 ms/op    7200003 B/op   200000 allocs/op
Benchmark_glennbrown-skiplist                    39   33.13 ms/op   12016815 B/op   400017 allocs/op
Benchmark_petar-GoLLRB_InsertNoReplace           34   34.17 ms/op    7200003 B/op   200000 allocs/op
Benchmark_gods-avltree                           36   37.87 ms/op    4800068 B/op   100001 allocs/op
Benchmark_petar-GoLLRB_ReplaceOrInsert           32   41.40 ms/op    7200009 B/op   200000 allocs/op
Benchmark_kellydunn-art                          25   43.29 ms/op   28770796 B/op   799621 allocs/op
Benchmark_3xian-elephantlist                     24   47.53 ms/op   12561963 B/op   237508 allocs/op
Benchmark_ryszard-skiplist                       19   65.56 ms/op   66400384 B/op   500006 allocs/op
```
