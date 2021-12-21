.PHONY:	all clean nuke README.md edit editor

all:
	GOMAXPROCS=1 redo
	make README.md

clean:
	rm -rf *.list *.workload *.pretty .redo/ bench/modernc-b-codegen/gen/btree.go
	find . -type f -name '*.results' -delete
	find . -type f -name '*.test' -delete

README.md:
	rm -f README.md
	redo README.md

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile README.edit.md & fi

editor:
	gofmt -l -s -w .
