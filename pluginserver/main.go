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

// simple plugin test
func SimpleTest(message string) string {
	var messageOut string = fmt.Sprintf("SimpleTest received this: %s", message)
	log.Print(messageOut)
	return messageOut
}

// accepts a callback function and a string to pass to the callback
func CallbackTest(callbacks plugincommon.CallbackFunctions, val int, message string) (int, string) {
	log.Print("CallbackTest called")
	return callbacks.AddOne(val), callbacks.AddMessage(message)
}
