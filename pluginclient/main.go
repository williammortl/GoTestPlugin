package main

import (
	"fmt"
	"log"
	"plugin"

	"github.com/WilliamMortl/GoTestPlugin/plugincommon"
)

const plugin_name string = "pluginserver.so"
const plugin_function_simple string = "SimpleTest"
const plugin_function_setcallbacks string = "SetCallbacks"
const plugin_function_callback string = "CallbackTest"
const plugin_function_ccallback string = "CCallbackTest"

type Callback struct{}

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

	funcSetCallbacks, err := p.Lookup(plugin_function_setcallbacks)
	if err != nil {
		panic(err)
	}
	var callbackFunctions Callback
	funcSetCallbacks.(func(plugincommon.CallbackFunctions))(&callbackFunctions)
	log.Print("Set callback")

	funcCallbackTest, err := p.Lookup(plugin_function_callback)
	if err != nil {
		panic(err)
	}
	numberOut, messageOut := funcCallbackTest.(func(int64, string) (int64, string))(4, "hello")
	log.Print(fmt.Sprintf("Plugin client received this returned number: %d", numberOut))
	log.Print(fmt.Sprintf("Plugin client received this returned string: %s", messageOut))

	funcCCallbackTest, err := p.Lookup(plugin_function_ccallback)
	if err != nil {
		panic(err)
	}
	numberOut = funcCCallbackTest.(func(int64) int64)(4)
	log.Print(fmt.Sprintf("Plugin client received this returned number from a C dylib call back from server into this client: %d", numberOut))
}

func (c *Callback) AddOne(val int64) int64 {
	log.Print("Called AddOne in pluginclient")
	return val + 1
}

func (c *Callback) AddMessage(message string) string {
	var messageOut string = fmt.Sprintf("%s - got it in AddMessage", message)
	log.Print("Called AddMessage in pluginclient")
	return messageOut
}
