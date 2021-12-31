all: clean build

clean:
	rm -rf *.so

build: main.go
	go build -buildmode=plugin -o pluginserver.so
