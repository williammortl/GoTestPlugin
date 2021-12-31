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

// simple plugin test
func SimpleTest(message string) string {
	var messageOut string = fmt.Sprintf("SimpleTest received this: %s", message)
	log.Print(messageOut)
	return messageOut
}

// accepts a callback function and a string to pass to the callback
/*func CallbackTest(long (*callback) (char* string), char* stringToPass) long
{
    printf("callback test input: %s\r\n", stringToPass);
    fflush(stdout);
    return (*callback)(stringToPass);
}*/
