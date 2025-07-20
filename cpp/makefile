CC=g++

target=output/main

#srcpp=main.cpp add.cpp

# srcpp:=$(shell find ./* -type f | egrep '\.c|\.cpp')
srcpp:=$(shell find ./* -type f | egrep '\.cpp')
obj=$(srcpp:.cpp=.o)
all:$(target)
$(target):$(obj)
	$(CC) -std=c++17 $(obj) -o $(target)

.cpp.o:
	$(CC) -std=c++17 -I inc -c $< -o $@

.PHONY:clean

clean:
	rm $(shell find ./src -type f | egrep '\.o')
	rm *.o $(target)