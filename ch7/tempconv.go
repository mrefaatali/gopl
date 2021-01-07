package main

import (
	"flag"
	"fmt"
)
/*
type Celsius float64
type Fahrenheit float64

type celsiusFlag struct{ Celsius }

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func FToC(f Fahrenheit) Celsius {
	c := (f - 32) * 5 / 9
	return Celsius(c)
}

//tough defined for Celsius it is promoted to celsiusFlag, though satisfying the String() function of flag.Value interface
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

//satisfies the flag.Value Set method
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}


func main() {
	var temp = CelsiusFlag("temp", 20.0, "the temperature")

	flag.Parse()
	fmt.Println(*temp)
}
*/