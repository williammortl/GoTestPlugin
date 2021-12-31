package main

import (
	"fmt"
	"log"
	"plugin"
)

const plugin_name string = "pluginserver.so"
const plugin_function_simple string = "SimpleTest"
const plugin_function_callback string = "CallbackTest"

type Callbacks struct{}

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

	funcCallbackTest, err := p.Lookup(plugin_function_callback)
	if err != nil {
		panic(err)
	}
	var callbackFunctions Callbacks
	numberOut, messageOut := funcCallbackTest.(func(CallbackFunctions, int) (int, string))(callbackFunctions, 4)
	log.Print(fmt.Sprintf("Plugin client received this: %s", messageOut))
}

func (c *Callbacks) AddOne(val int) int {
	log.Print("Called AddOne in pluginclient")
	return val + 1
}

func (c *Callbacks) AddMessage(message string) string {
	var messageOut string = fmt.Sprintf("%s - got it in AddMessage", message)
	log.Print(messageOut)
	return messageOut
}