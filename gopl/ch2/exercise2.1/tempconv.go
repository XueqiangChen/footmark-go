// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度
type Kelvin float64 //绝对零度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度

	AbsoluteZeroK Kelvin = 0
	FreezingK Kelvin = 273.15
	BoilingK Kelvin = 373.15

	AbsoluteZeroF Fahrenheit = -459.66999999999996
	FreezingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) string() string { return fmt.Sprintf("%g°K", k) }
