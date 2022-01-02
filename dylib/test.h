#ifndef TEST_H
#define TEST_H

#include <stdio.h>

// definitions
#define SHARED_LIBRARY "./libdylib.so"
#define SHARED_LIBRARY_FUNCTION_SIMPLE "SimpleTest"
#define SHARED_LIBRARY_FUNCTION_CALLBACK "CallbackTest"
#define STRING_TEST_STRING "testtesttest"
#define TEST_VALUE 1604

// function types
typedef void (SimpleTest)(char* stringToPass);
typedef long (CallbackTest)(long (*callback)(long val), long val);

// functions
long AddOneCallback(long val);

#endif