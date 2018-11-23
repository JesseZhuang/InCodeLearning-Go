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
