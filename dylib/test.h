#ifndef TEST_H
#define TEST_H

#include <stdio.h>

// definitions
#define SHARED_LIBRARY "./libdylib.so"
#define SHARED_LIBRARY_FUNCTION_SIMPLE "SimpleTest"
#define SHARED_LIBRARY_FUNCTION_CALLBACK "CallbackTest"
#define TEST_VALUE 4

// function types
typedef void (SimpleTest)(char* stringToPass);
typedef long (CallbackTest)(long (*callback)(char* string), char* stringToPass);

// functions
long Callback(char* stringToPass);

#endif