package main

import (
	"fmt"
	"os"
	"strconv"
	"unitconv/unitconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
			os.Exit(1)
		}

		var Far unitconv.Fahrenheit = unitconv.Fahrenheit(v)
		var Feet unitconv.Feet = unitconv.Feet(v)
		var Cel unitconv.Celsius = unitconv.Celsius(v)
		var Meters unitconv.Meters = unitconv.Meters(v)
		var KG unitconv.Kilograms = unitconv.Kilograms(v)
		var LB unitconv.Pounds = unitconv.Pounds(v)

		fmt.Println("conversion from", v)
		fmt.Printf("%s is %s : %s is %s\n", Far, unitconv.FtoC(Far), Cel, unitconv.CtoF(Cel))
		fmt.Printf("%s is %s : %s is %s\n", Feet, unitconv.FtoM(Feet), Meters, unitconv.MtoF(Meters))
		fmt.Printf("%s is %s : %s is %s\n", LB, unitconv.LBtoKG(LB), KG, unitconv.KGtoLB(KG))
	}
}
