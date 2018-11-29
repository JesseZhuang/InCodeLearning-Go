package main

import (
	"fmt"
	"runtime"
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
		float32Num, float64Num, cmplx64Num, cmplx128Num, world, trueth,
	}
	printTypes(vars)

	printSeparator()
	sum := 1
	for sum < 1000 { // for loop as while loop
		sum += sum
	}
	fmt.Println(sum)

	printSeparator()
	// Both calls to pow return their results before the call to fmt.Println in main begins.
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	printSeparator()
	fmt.Println(sqrtNewton(2))

	printSeparator()
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	printSeparator()
	t := time.Now()
	// switch with no condition, clean way to write long if-then-else clauses
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	printSeparator()
	deferred()
	deferredCounting()

	printSeparator()
	pointers()

	printSeparator()
	demoStruct()

	printSeparator()
	demoSlice()

	printSeparator()
	ticTacToe()

	printSeparator()
	demoRange()

	printSeparator()
	demoMap()

	printSeparator()
	fmt.Println(wordCount("I ate a donut. Then I ate another donut."))

	printSeparator()
	demoFunction()

	printSeparator()
	demoClosure()

	printSeparator()
	demoFibonacci()
}
