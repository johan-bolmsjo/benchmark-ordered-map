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
