.PHONY:	all clean nuke README.md

all:
	redo
	make README.md

clean:
	rm -rf *.list *.workload *.pretty .redo/ bench/modernc-b-codegen/gen/btree.go
	find . -type f -name '*.results' -delete
	find . -type f -name '*.test' -delete

README.md:
	rm -f README.md
	redo README.md
