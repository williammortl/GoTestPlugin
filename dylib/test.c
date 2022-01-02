#include <stdio.h>
#include <dlfcn.h>
#include <assert.h>
#include <string.h>
#include "test.h"

// main function
int main(int argc, char** argv)
{
    int retValue = -1;
    void* libraryHandle = NULL;
    SimpleTest* simpleFunc = NULL;
    CallbackTest* callbackFunc = NULL;
    char* message = "testtesttest";

    // load library
    libraryHandle = dlopen(SHARED_LIBRARY, RTLD_LAZY);
    if (libraryHandle != NULL)
    {
        printf("loaded lib!\r\n");
        fflush(stdout);

        // get the simple test procedure address
        simpleFunc = (SimpleTest*)dlsym(libraryHandle, SHARED_LIBRARY_FUNCTION_SIMPLE);
        if (simpleFunc != NULL)
        {
            printf("loaded func simple test!\r\n");
            fflush(stdout);

            //call function!
            (*simpleFunc)(STRING_TEST_STRING);
        }
        else
        {
            printf(dlerror());
        }

        // get the callback test procedure address
        callbackFunc = (CallbackTest*)dlsym(libraryHandle, SHARED_LIBRARY_FUNCTION_CALLBACK);
        if (callbackFunc != NULL)
        {
            printf("loaded func callback test!\r\n");
            fflush(stdout);

            //call function!
            retValue = (*callbackFunc)(AddOneCallback, TEST_VALUE - 1);
            printf("return value (should be %i): %i\r\n", TEST_VALUE, retValue);

            // make sure that return value is correct
            assert(retValue == TEST_VALUE);
        }
        else
        {
            printf(dlerror());
        }

        dlclose(libraryHandle);
    }
    else
    {
        printf(dlerror());
    }

    return 0;
}

// callback function
long AddOneCallback(long val)
{
    printf("Called AddOneCallback in test.c\r\n");
    fflush(stdout);
    return val + 1;
}