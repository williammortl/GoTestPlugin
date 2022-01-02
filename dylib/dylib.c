#include <stdio.h>
#include "dylib.h"

// simple dylib call test
extern void SimpleTest(char* stringToPass)
{
    printf("SimpleTest: %s\r\n", stringToPass);
    fflush(stdout);
    return;
}

// accepts a callback function and a string to pass to the callback
extern long CallbackTest(long (*callback) (long val), long val)
{
    printf("CallbackTest called in dylib.c\r\n");
    fflush(stdout);
    return (*callback)(val);
}