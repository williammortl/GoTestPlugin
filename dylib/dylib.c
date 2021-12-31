#include <stdio.h>
#include "dylib.h"

// simple dylib call test
extern void SimpleTest(char* stringToPass)
{
    printf("simple test input: %s\r\n", stringToPass);
    fflush(stdout);
    return;
}

// accepts a callback function and a string to pass to the callback
extern long CallbackTest(long (*callback) (char* string), char* stringToPass)
{
    printf("callback test input: %s\r\n", stringToPass);
    fflush(stdout);
    return (*callback)(stringToPass);
}