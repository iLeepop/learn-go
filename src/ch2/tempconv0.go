package tempconv0

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BolilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g`C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g`F", f)
}
