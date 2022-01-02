package main

import (
	"fmt"
	"log"

	"github.com/WilliamMortl/GoTestPlugin/plugincommon"
)

// Main function
func main() {

	// test SimpleTest
	log.Print(SimpleTest("internal test"))
}

var callbacks plugincommon.CallbackFunctions

// simple plugin test
func SimpleTest(message string) string {
	var messageOut string = fmt.Sprintf("SimpleTest received this: %s", message)
	log.Print(messageOut)
	return messageOut
}

// set the callbacks
func SetCallbacks(callbacksIn plugincommon.CallbackFunctions) {
	callbacks = callbacksIn
}

// tests the callback
func CallbackTest(val int64, message string) (int64, string) {
	log.Printf("CallbackTest called, received: %d, %s", val, message)
	return callbacks.AddOne(val), callbacks.AddMessage(message)
}

// tests calling back into AddOne through C dylib
func CCallbackTest(val int64) int64 {
	log.Printf("CCallbackTest called, received: %d", val)

	// simple test
	CSimpleTest("test from main.go")

	// dynamically provide callback for C
	retVal := CCallbackCallTest(val)

	return retVal
}
