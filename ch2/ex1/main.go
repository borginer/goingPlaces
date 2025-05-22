package main

import (
	"fmt"
	"tempconv/tempconv"
)

func assert(cond bool, msg string) {
	if !cond {
		panic("assertion failed: " + msg)
	}
}

func main() {
	assert(tempconv.KtoC(275.15) == 2, "k to c conversion check")
	assert(tempconv.CtoK(25) == 298.15, "c to k conversion check")
	assert(tempconv.FtoC(95) == 35, "f to c conversion check")
	assert(tempconv.CtoF(70) == 158, "c to f conversion check")
	assert(tempconv.KtoF(373.15) == 212, "k to f conversion check")
	assert(tempconv.FtoK(392) == 473.15, "f to k conversion check")

	var c tempconv.Celsius = 2
	var f tempconv.Fahrenheit = 30
	var k tempconv.Kelvin = 200

	assert(fmt.Sprintf("%s, %s, %s", c, f, k) == "2°C, 30°F, 200°K",
		"String formating check")

	fmt.Println("tempconv works! (probably)")
}
