#include <stdio.h>
#include <dlfcn.h>
#include "test.h"

// main function
int main(int argc, char** argv)
{
    int retValue = -1;
    void* libraryHandle = NULL;
    CallbackTest* func = null;

    // load library
    libraryHandle = dlopen(SHARED_LIBRARY, RTLD_LAZY);
    if (libraryHandle != NULL)
    {

        // get the procedures address
        func = (CallbackTest*)dlsym(handle, SHARED_LIBRARY_FUNCTION);
        if (func != NULL)
        {
        
            //call function!

            // close library
            dlclose(libraryHandle);
        }
    }

    return retValue;
}

// callback function
long Callback(char* stringToPass)
{
    return 4;
}