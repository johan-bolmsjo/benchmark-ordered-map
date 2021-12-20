.PHONY:	all clean

all:
	redo

clean:
	rm -rf *.list *.workload *.pretty .redo/ bench/modernc-b-codegen/gen/btree.go
	find . -type f -name '*.results' -delete
	find . -type f -name '*.test' -delete
