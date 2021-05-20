from ctypes import *
lib = cdll.LoadLibrary("./main.so")
lib.createGo("./src/examples/elelock.v")