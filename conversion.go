package conversion

import "fmt"

// Celcius Temperature Type
type Celcius float64

// Fahrenheit Temperature Type
type Fahrenheit float64

const (
	// AbsoluteZeroC The point of absolute zero in celcius
	AbsoluteZeroC Celcius = -273.15
	// FreezingC The temperatute in Celcius that water freezes
	FreezingC Celcius = 0
	// BoilingC the temperature in Celcius that water boils.
	BoilingC Celcius = 100
)

func (c Celcius) String() string    { return fmt.Sprintf("%g•C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g•F", f) }

// CToF Convert Celcius to Fahrenheit
func CToF(c Celcius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC Convert Fahrenheit to Celcius
func FToC(f Fahrenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}
