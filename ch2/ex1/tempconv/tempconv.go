package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CtoK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

func KtoC(k Kelvin) Celsius {
	return Celsius(k) + AbsoluteZeroC
}

func FtoK(f Fahrenheit) Kelvin {
	return CtoK(FtoC(f))
}

func KtoF(k Kelvin) Fahrenheit {
	return CtoF(KtoC(k))
}
