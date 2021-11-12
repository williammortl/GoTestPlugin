#include <stdio.h>
#include "dylib.h"

// accepts a callback function and a string to pass to the callback
extern long CallbackTest(long (*callback) (char* string), char* stringToPass)
{
    return callback(stringToPass);
}