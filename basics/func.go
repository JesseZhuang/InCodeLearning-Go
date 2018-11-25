package main

import (
	"fmt"
	"math"
)

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // "naked" return
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func sqrtNewton(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / 2 / z
		fmt.Printf("Round %d: value %g.\n", i, z)
	}
	return z
}

func deferred() {
	// arugments evaluated but not executed until surrounding function returns
	defer fmt.Println("world")
	fmt.Println("Hello")
}

// stacking defers
func deferredCounting() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
