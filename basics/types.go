package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"unsafe"
)

// variable declarations may be "factored" into blocks, as with import statements.
var (
	toBe     bool
	text     string
	intNum   int   // in practice int typically equals the processor bit-size
	int8Num  int8  = -128
	int16Num int16 = math.MaxInt16
	// code points: ascii 128, extended ascii 256, unicode 17 * 2^16 = 1,114,112
	runeNum   rune   = 23 // alias for int32, represents a Unicode code point
	int32Num  int32  = math.MaxInt32
	maxInt64  int64  = math.MaxInt64
	unitNum   uint   = 600
	byteNum   byte   = 23        // alias for uint8
	maxuint8         = ^uint8(0) // go does not have ~
	minuint8         = ^maxuint8
	uint16Num uint16 = 23
	uint32Num uint32 = 23
	maxuint64 uint64 = 1<<64 - 1 // 0xFFFF FFFF FFFF FFFF
	// translating the pointer to the index it represents in your virtual memory address space
	// uintptr doesn't represent a pointer to an unsigned integer
	// but is a type to treat a pointer as an unsigned integer for pointer arithmetic
	uintPtr               = uintptr(unsafe.Pointer(&intNum))
	float32Num  float32   = 1.2
	float64Num            = 2.3
	cmplx64Num  complex64 = -5 + 12i
	cmplx128Num           = cmplx.Sqrt(-5 + 12i)
)

// const cannot be delared with := syntax
const (
	world  = "世界"
	trueth = true
)

// https://appliedgo.net/generics/

// empty interface can be anything
func printType(anything interface{}) {
	fmt.Printf("Type: %T Value: %v\n", anything, anything)
}

func printTypes(types []interface{}) {
	for _, anything := range types {
		printType(anything)
	}
}

// unlike C, no pointer arithmetic
func pointers() {
	var p *int
	fmt.Println(p) // nil
	// fmt.Println(*p) // Failed to continue - bad access
	i := 42
	p = &i
	fmt.Println(*p)
}

type vertex struct {
	X int
	Y int
}

func demoStruct() {
	v := vertex{1, 2}
	fmt.Println(v)
	v.X = 5
	fmt.Println(v)
	p := &v
	p.Y = 1e8
	fmt.Println(p, v)
	(*p).X = 123 // explicit dereference
	v2 := vertex{Y: 3}
	fmt.Println(v, v2)
}

// A slice does not store any data, a reference to array
func demoSlice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s = primes[1:4]
	fmt.Println(s)
	fmt.Println(primes[:2])
	fmt.Println(primes[1:])
	fmt.Println(primes[:])
	printSlice(s)
	s = primes[:0]
	if s != nil {
		fmt.Println("not nil")
	}
	printSlice(s)
	s = s[:6] // cannot slice out of bound
	printSlice(s)
	s = s[2:]
	printSlice(s)
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	// slice literal
	structSlice := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(structSlice)
	var zeroSlice []int
	printSlice(zeroSlice)
	if zeroSlice == nil {
		fmt.Println("nil")
	}
	s1 := make([]int, 5)
	printSlice(s1)
	s2 := make([]int, 0, 5)
	printSlice(s2)
	c := s2[:2]
	printSlice(c)
	d := c[2:5]
	printSlice(d)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
