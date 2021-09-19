package main

import (
	"fmt"
	"math"
)

//Print integers 1 to N,
//but print “Fizz” if an integer is divisible by 3,
//“Buzz” if an integer is divisible by 5, and
//0“FizzBuzz” if an integer is divisible by both 3 and 5.
func main() {



	buzzer:=New()
	fmt.Printf("Buzzering out put is %s  ",buzzer.PrintInteger(15))
}

func New() Buzz {
	return &buzzer{
		Divider3: 3,
		Divider5: 5,
	}
}

type Buzz interface {
	PrintInteger(input float64) string
}

type  buzzer struct {
	Divider3 float64
	Divider5 float64
}

func(b buzzer)PrintInteger(input float64)(result string)   {

	if math.Mod(input,b.Divider3)==0 && math.Mod(input,b.Divider5)==0{
		result ="FizzBuzz"
	}else if answer:=math.Mod(input,b.Divider3);answer==0{
		result= "Fizz"
	}else if answer:=math.Mod(input,b.Divider5);answer==0{
		result= "Buzz"
	}
	return result
}