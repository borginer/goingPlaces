package unitconv

import (
	"fmt"
	"strings"
)

type Celsius float64
type Fahrenheit float64
type Meters float64
type Feet float64
type Pounds float64
type Kilograms float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	KGtoLBSRatio          = 2.20462262
	MtoFRatio             = 3.28084
)

func FormatFloat(f float64) string {
	s := fmt.Sprintf("%.3f", f)
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, ".")
	return s
}

func (c Celsius) String() string {
	return fmt.Sprintf("%s°C", FormatFloat(float64(c)))
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%s°F", FormatFloat(float64(f)))
}

func (kg Kilograms) String() string {
	return fmt.Sprintf("%s'kg", FormatFloat(float64(kg)))
}

func (p Pounds) String() string {
	return fmt.Sprintf("%s'lb", FormatFloat(float64(p)))
}

func (m Meters) String() string {
	return fmt.Sprintf("%s'm", FormatFloat(float64(m)))
}

func (f Feet) String() string {
	return fmt.Sprintf("%s'ft", FormatFloat(float64(f)))
}

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KGtoLB(kg Kilograms) Pounds {
	return Pounds(kg * KGtoLBSRatio)
}

func LBtoKG(lb Pounds) Kilograms {
	return Kilograms(lb / KGtoLBSRatio)
}

func FtoM(f Feet) Meters {
	return Meters(f / MtoFRatio)
}

func MtoF(m Meters) Feet {
	return Feet(m * MtoFRatio)
}
