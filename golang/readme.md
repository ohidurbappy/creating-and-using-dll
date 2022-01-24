### Using golang to generate .dll

```
package main

import "C"

//export Square
func Square(number int) int {
	return number * number
}

func main() {

}
```

### Build

```
go build -o mydll.dll --buildmode=c-shared
```

### Using the generated dll file with python

```
import ctypes
lib=ctypes.CDLL("./mydll.dll")
print(lib.Square(23))
```

### Test dll

```
rundll32 mydll,Square
```