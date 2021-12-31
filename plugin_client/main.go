package main

import (
	"fmt"
	"log"
	"plugin"
)

const plugin_name string = "plugin_server.so"
const plugin_function_simple string = "SimpleTest"

// Main function
func main() {

	p, err := plugin.Open(plugin_name)
	if err != nil {
		panic(err)
	}

	funcSimpleTest, err := p.Lookup(plugin_function_simple)
	if err != nil {
		panic(err)
	}
	messageOut := funcSimpleTest.(func(string) string)("hello from the client!")
	log.Print(fmt.Sprintf("Plugin client received this: %s", messageOut))
}
