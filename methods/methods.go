package main

import (
	"fmt"
	"math"
)

func printSeparator() {
	fmt.Println("--------------------")
}

func main() {
	v := vertex{3, 4}
	fmt.Println(v.abs())
	fmt.Println(abs(v))

	printSeparator()
	floatNum := myFloat(-2.34)
	fmt.Println(floatNum.abs())
	floatNum = myFloat(-math.Sqrt2)
	fmt.Println(floatNum.abs())

	printSeparator()
	v = vertex{3, 4}
	v.cannotScale(10)
	fmt.Println(v.abs())
	v.scale(10) // pointer indirection
	fmt.Println(v.abs())
	(&v).scale(10)
	fmt.Println(v.abs())

	printSeparator()
	var a abser
	a = myFloat(-math.Sqrt2)
	printType(a)
	fmt.Println(a.abs())
	a = vertex{3, 4}
	printType(a)
	fmt.Println(a.abs())
	a = &vertex{3, 4}
	printType(a)
	fmt.Println(a.abs())
	// a = myInt32(-6) // doe not compile, myInt32 does not implement abser
	myInt := myInt32(-6)
	a = &myInt
	printType(a)
	fmt.Println(a.abs())

	printSeparator()
	var i i
	printType(i)
	// i.m() runtime error
	var t *myString
	i = t
	printType(i)
	i.m()
	i = &myString{"hello"}
	printType(i)
	i.m()

	printSeparator()
	typeAssertion()

	printSeparator()
	typeSwitch(21)
	typeSwitch("j-test")
	typeSwitch(true)

	printSeparator()
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	printSeparator()
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))

	printSeparator()
	demoReader()
}
