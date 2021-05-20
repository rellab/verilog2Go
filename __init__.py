from ctypes import *

def main():
    lib = cdll.LoadLibrary("./verilog2Go/main.so")
    lib.createGo("./verilog2Go/src/examples/elelock.v")