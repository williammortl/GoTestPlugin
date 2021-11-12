#ifndef TEST_H
#define TEST_H

#include <stdio.h>

// definitions
#define SHARED_LIBRARY "libdylib.so"
#define SHARED_LIBRARY_FUNCTION "CallbackTest"

// function types
typedef long (*CallbackTest)(long (*callback)(char* string), char* stringToPass);

// functions
long Callback(char* stringToPass);

#endif