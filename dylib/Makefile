all: clean build

clean:
	rm -rf *.o 
	rm -rf *.so

build: dylib.h dylib.c
	gcc -c -Wall -fPIC -c dylib.c
	gcc -shared -o libdylib.so dylib.o
	rm -rf dylib.o

test: clean build test.c
	gcc test.c -o test
	./test
