all: clean build

buildDylib: dylib/dylib.h dylib/dylib.c
	$(MAKE) -C dylib
	cp dylib/libdylib.so bin/.

buildPlugin: plugin/main.go
	$(MAKE) -C plugin
	cp plugin/plugin bin/.

buildCli: cli/main.go
	$(MAKE) -C cli
	cp cli/cli bin/.

clean:
	rm -rf bin

createDirectories:
	mkdir bin

build: createDirectories buildDylib buildPlugin buildCli
	ls bin
