package main

import "fmt"

// in integers, unsigned integers means that the number cannot be negative
// signed integers means that the number can be negative

func main() {
	// uint8
	// 0 <-> 255
	var a uint8 = 255

	// uint16
	// 0 <-> 65535
	var b uint16 = 65535

	// uint32
	// 0 <-> 4294967295
	var c uint32 = 4294967295

	// uint64
	// 0 <-> 18446744073709551615
	var d uint64 = 18446744073709551615

	// uint
	// 0 <-> 18446744073709551615 or 0 <-> 4294967295 depending on the system
	var e uint = 18446744073709551615

	// int8
	// -128 <-> 127
	var f int8 = -128

	// int16
	// -32768 <-> 32767
	var g int16 = -32768

	// int32
	// -2147483648 <-> 2147483647
	var h int32 = -2147483648

	// int64
	// -9223372036854775808 <-> 9223372036854775807

	// int
	// -9223372036854775808 <-> 9223372036854775807 or -2147483648 <-> 2147483647 depending on the system
	var i int = -9223372036854775808

	// rune
	// rune in go is an alias for int32
	// it represents a Unicode code point
	var j rune = 2147483647

	fmt.Print(a, b, c, d, e, f, g, h, i, j)

	// float32
	// 1.18e-38 <-> 3.4e38
	var k float32 = 3.4e38

	// float64
	// 2.23e-308 <-> 1.8e308
	var l float64 = 1.8

	fmt.Println(k, l)

	var exists bool = true

	fmt.Println(exists)

	var message string = "Hello, World!"

	fmt.Println(message)

	// nil in go doesn't mean null, it means that the variable has no value
	/// nil is the zero value for pointers, interfaces, maps, slices, channels and functions
	// zero value from strings is an empty string ""
	// zero value for booleans is false
	var emptyString string

	fmt.Println(emptyString)
}