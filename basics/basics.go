package main

import (
	"fmt"
	"time"
)

// A var statement can be at package or function level
var c, python, java bool

func printSeparator() {
	fmt.Println("--------------------")
}

// go run <filename> on command line, ^F5 for VS code go extension
func main() {
	fmt.Println("Hello World Go!")
	fmt.Println("The time is", time.Now())

	printSeparator()
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	printSeparator()
	j := 3 // type inferred from literal, := only in function
	fmt.Println(c, python, java, j)

	printSeparator()
	values := []interface{}{
		map[string]string{"name": "ravi"},
		[]string{"art", "coding", "music", "travel"},
	}
	fmt.Println(values)

	printSeparator()
	vars := []interface{}{
		toBe, text, intNum, int8Num, int16Num, runeNum, int32Num, maxInt64,
		unitNum, byteNum, maxuint8, minuint8, uint16Num, uint32Num, maxuint64, uintPtr,
		float32Num, float64Num, cmplx64Num, cmplx128Num,
	}
	printTypes(vars)
}
