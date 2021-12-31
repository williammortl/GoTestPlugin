#ifndef DYLIB_H
#define DYLIB_H

#include <stdio.h>

// exported functions 
extern void SimpleTest(char* stringToPass);

extern long CallbackTest(long (*callback) (char* string), char* stringToPass);

#endif