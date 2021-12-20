.PHONY:	all clean

all:
	redo

clean:
	rm -rf *.list *.workload *.pretty .redo/
	find . -type f -name '*.results' -delete
	find . -type f -name '*.test' -delete
