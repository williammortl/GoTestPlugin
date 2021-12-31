all: clean createdirs build

dylib: dylib/dylib.h dylib/dylib.c
	$(MAKE) -C dylib test
	cp dylib/libdylib.so bin/.

pluginserver: pluginserver/main.go
	$(MAKE) -C pluginserver
	cp pluginserver/pluginserver.so bin/.

pluginclient: pluginclient/main.go
	$(MAKE) -C pluginclient
	cp pluginclient/pluginclient.exe bin/.

clean:
	rm -rf bin
	$(MAKE) -C dylib clean
	$(MAKE) -C pluginserver clean
	$(MAKE) -C pluginclient clean

createdirs:
	mkdir bin

build: dylib pluginserver pluginclient
	ls bin
