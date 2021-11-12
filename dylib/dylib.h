#ifndef DYLIB_H
#define DYLIB_H

#include <stdio.h>

// exported functions
extern long CallbackTest(long (*callback) (char* string), char* stringToPass);

#endif