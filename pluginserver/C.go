package main

/*
#cgo LDFLAGS: -ldl

#include <stdio.h>
#include <dlfcn.h>

// import Go functions
extern long CallbackWrapper(long);

// definitions
#define SHARED_LIBRARY "./libdylib.so"
#define SHARED_LIBRARY_FUNCTION_SIMPLE "SimpleTest"
#define SHARED_LIBRARY_FUNCTION_CALLBACK "CallbackTest"

// function types
typedef void (SimpleTest)(char* stringToPass);
typedef long (CallbackTest)(long (*callback)(long val), long val);

static void DylibSimpleTest(char* stringToPass)
{
	void* libraryHandle = NULL;
    SimpleTest* simpleFunc = NULL;

	// load dll
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
            (*simpleFunc)(stringToPass);
        }
        else
        {
            printf(dlerror());
        }

		// close dll
		dlclose(libraryHandle);
	}
	else
    {
        printf(dlerror());
    }
}

// this is the pointer passed to the call back
static long CCallbackWrapper(long val)
{
	return CallbackWrapper(val);
}

static long DylibCallbackTest(long val)
{
	long retVal = -1;
	void* libraryHandle = NULL;
    CallbackTest* callbackFunc = NULL;

	// load dll
    libraryHandle = dlopen(SHARED_LIBRARY, RTLD_LAZY);
    if (libraryHandle != NULL)
    {
        printf("loaded lib!\r\n");
        fflush(stdout);

        // get the callback test procedure address
        callbackFunc = (CallbackTest*)dlsym(libraryHandle, SHARED_LIBRARY_FUNCTION_CALLBACK);
        if (callbackFunc != NULL)
        {
            printf("loaded func callback test!\r\n");
            fflush(stdout);

            //call function!
            retVal = (*callbackFunc)(CCallbackWrapper, val);
        }
        else
        {
            printf(dlerror());
        }

		// close dll
		dlclose(libraryHandle);
	}
	else
    {
        printf(dlerror());
    }

	return retVal;
}

*/
import "C"
import "log"

func CSimpleTest(message string) {
	C.DylibSimpleTest(C.CString(message))
}

func CCallbackCallTest(val int64) int64 {
	return int64(C.DylibCallbackTest(C.long(val)))
}

//export CallbackWrapper
func CallbackWrapper(val C.long) C.long {
	log.Print("CallbackWrapper called from dylib.c via C code")
	return C.long(callbacks.AddOne(int64(val)))
}
