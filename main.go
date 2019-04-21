package main

import "C"

//export CalcDouble
func CalcDouble(x float64) float64 {
	return 2 * x
}

// main func is needed
func main() {}
