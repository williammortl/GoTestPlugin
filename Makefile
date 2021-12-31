all: clean createdirs build

dylib: dylib/dylib.h dylib/dylib.c
	$(MAKE) -C dylib test
	cp dylib/libdylib.so bin/.

plugin_server: plugin_server/main.go
	$(MAKE) -C plugin_server
	cp plugin_server/plugin_server.so bin/.

plugin_client: plugin_client/main.go
	$(MAKE) -C plugin_client
	cp plugin_client/plugin_client.exe bin/.

clean:
	rm -rf bin
	$(MAKE) -C dylib clean
	$(MAKE) -C plugin_server clean
	$(MAKE) -C plugin_client clean

createdirs:
	mkdir bin

build: dylib plugin_server plugin_client
	ls bin
