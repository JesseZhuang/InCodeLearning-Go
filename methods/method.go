package main

import "math"

type vertex struct {
	X, Y float64
}

// the abs method has a receiver of type vertex named v.
func (v vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// a method is just a function with a receiver argument.
func abs(v vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v vertex) cannotScale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// pointer receivers are more common than value receivers
// go uses pass by value, takes either pointer or value
// when to use pointer receiver: 1, modify value; 2, avoid value copying for large struct
// in general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.
func (v *vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type myFloat float64

// can only declare a method with a receiver whose type is defined in the same package as the method.
func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type myInt32 int32

func (i *myInt32) abs() float64 {
	if (*i) < 0 {
		return float64(-*i)
	}
	return float64(*i)
}
