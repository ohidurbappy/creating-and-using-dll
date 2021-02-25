**Create an Entry Point**

Create `UnitEntryPoint.cpp`

```cpp
#include <windows.h>

int WINAPI DllEntryPoint(HINSTANCE hinst, unsigned long reason, void* lpReserved)
{
  return 1;
}

int WINAPI WinMain(      
    HINSTANCE hInstance,
    HINSTANCE hPrevInstance,
    LPSTR lpCmdLine,
    int nCmdShow)
{  
  return 0;
}

```

Create `UnitFunctions.h`

```cpp

#ifndef UnitFunctionsH
#define UnitFunctionsH

#ifdef __cplusplus
extern "C"
{
#endif

__declspec (dllexport) const int GetAnswerOfLife();

#ifdef __cplusplus
}
#endif


#endif
```

The function put in the DLL is called GetAnswerOfLife and will return the value of 42. Note the #ifdef's before and after the function. These are obligatory!

Now the function GetAnswerOfLife must be defined in UnitFunctions.cpp. Upon viewing it, the code looks like below

Create `UnitFunctions.cpp`

```cpp

#include "UnitFunctions.h"

const int GetAnswerOfLife()
{
  return 42;
}

```

Now we are ready to build our first dll

Here is the build script

```bash
g++ -c UnitFunctions.cpp
g++ -c UnitEntryPoint.cpp
g++ -o Functions.dll UnitEntryPoint.o UnitFunctions.o
```

Now we can call the dll. There are two ways to call a dll.

- Dynamically
- Statically


Let's use the dll functins dynamically first.

Create `main.cpp`


```cpp
#include <windows.h>

struct DllHandle
{
    DllHandle(const char *const filename)
        : h(LoadLibrary(filename)) {}
    ~DllHandle()
    {
        if (h)
            FreeLibrary(h);
    }
    const HINSTANCE Get() const { return h; }

private:
    HINSTANCE h;
};

int main()
{
    //Obtain a handle to the DLL
    const DllHandle h("Functions.DLL");
    if (!h.Get())
    {
        MessageBox(0, "Could not load DLL", "UnitCallDll", MB_OK);
        return 1;
    }

    //Obtain a handle to the GetAnswerOfLife function
    typedef const int (*GetAnswerOfLifeFunction)();
    const GetAnswerOfLifeFunction AnswerOfLife = reinterpret_cast<GetAnswerOfLifeFunction>(
        GetProcAddress(h.Get(), "GetAnswerOfLife"));

    if (!AnswerOfLife) //No handle obtained
    {
        MessageBox(0, "Loading AnswerOfLife failed", "UnitCallDll", MB_OK);
        return 1;
    }

    if (AnswerOfLife() != 42)
    {
        MessageBox(0, "Function AnswerOfLife failed", "UnitCallDll", MB_OK);
        return 1;
    }
    else
    {
        MessageBox(0, "Function AnswerOfLife successful!", "UnitCallDll", MB_OK);
    }
}

```

Now build and run

```bash
g++ main.cpp -o main.exe && ./main.exe
```


I have taken it from http://www.richelbilderbeek.nl/CppGppCreateDll.htm