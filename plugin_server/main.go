package main

import (
	"fmt"
	"log"
)

// Main function
func main() {

	// test SimpleTest
	log.Print(SimpleTest("internal test"))
}

func SimpleTest(message string) string {
	var messageOut string = fmt.Sprintf("SimpleTest received this: %s", message)
	log.Print(messageOut)
	return messageOut
}
