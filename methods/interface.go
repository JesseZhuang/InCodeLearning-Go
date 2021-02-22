package main

import (
	"fmt"
	"io"
	"math"
	"strings"
)

// implicit implementation, no implements clause needed
type abser interface {
	abs() float64
}

// interface values can be thought of as a tuple of a value and a concrete type
func printType(anything interface{}) {
	fmt.Printf("Type: %T Value: %v\n", anything, anything)
}

type i interface {
	m()
}

type myString struct {
	S string
}

// in Go it is common to write methods that gracefully handle being called with a nil receiver
func (s *myString) m() {
	if s == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(s.S)
}

func typeAssertion() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(s, ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)
	// f = i.(float64) // panic
}

func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// IPAddr represents an IP address
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func demoReader() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
