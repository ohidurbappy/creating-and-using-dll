import ctypes

lib=ctypes.CDLL("./mydll.dll")

print(lib.Square(23))

