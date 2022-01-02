all: clean createdirs build

dylib: dylib/dylib.h dylib/dylib.c
	cd dylib; \
		gcc -c -Wall -fPIC -c dylib.c; \
		gcc -shared -o libdylib.so dylib.o; \
		rm -rf dylib.o;\
		gcc test.c -ldl -o testdylib.exe
	cp dylib/libdylib.so bin/.
	cp dylib/testdylib.exe bin/.
	rm -rf dylib/libdylib.so
	rm -rf dylib/testdylib.exe

pluginserver: pluginserver/main.go
	cd pluginserver; \
		go build -buildmode=plugin -o pluginserver.so
	cp pluginserver/pluginserver.so bin/.
	rm -rf pluginserver/pluginserver.so

pluginclient: pluginclient/main.go
	cd pluginclient; \
		go build -o pluginclient.exe
	cp pluginclient/pluginclient.exe bin/.
	rm -rf pluginclient/pluginclient.exe

clean:
	rm -rf bin

createdirs:
	mkdir bin

build: dylib pluginserver pluginclient
	ls bin
