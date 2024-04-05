package main

import "fmt"

func main() {
	// Declare variables



	// Integer types (-128 to 127)
	var intVar int = 10
	var int8Var int8 = 20   // Range: -128 to 127
	var int16Var int16 = 30 // Range: -32768 to 32767
	var int32Var int32 = 40 // Range: -2147483648 to 2147483647
	var int64Var int64 = 50 // Range: -9223372036854775808 to 9223372036854775807

	// Unsigned integer types (0 to 255)
	var uintVar uint = 60
	var uint8Var uint8 = 70    // Range: 0 to 255
	var uint16Var uint16 = 80  // Range: 0 to 65535
	var uint32Var uint32 = 90  // Range: 0 to 4294967295
	var uint64Var uint64 = 100 // Range: 0 to 18446744073709551615

	// Floating point types
	var float32Var float32 = 3.14 // Range: -3.4e38 to +3.4e38
	var float64Var float64 = 6.28 // Range: -1.7e308 to +1.7e308

	// Complex types
	var complex64Var complex64 = 1 + 2i   // 32-bit real and imaginary parts
	var complex128Var complex128 = 3 + 4i // 64-bit real and imaginary parts

	// Boolean type
	var boolVar bool = true // true or false

	// String type
	var stringVar string = "Hello, Go!"

	// Print variable values
	fmt.Println("Integer:", intVar)
	fmt.Println("Int8:", int8Var)
	fmt.Println("Int16:", int16Var)
	fmt.Println("Int32:", int32Var)
	fmt.Println("Int64:", int64Var)
	fmt.Println("Unsigned Integer:", uintVar)
	fmt.Println("Uint8:", uint8Var)
	fmt.Println("Uint16:", uint16Var)
	fmt.Println("Uint32:", uint32Var)
	fmt.Println("Uint64:", uint64Var)
	fmt.Println("Float32:", float32Var)
	fmt.Println("Float64:", float64Var)
	fmt.Println("Complex64:", complex64Var)
	fmt.Println("Complex128:", complex128Var)
	fmt.Println("Boolean:", boolVar)
	fmt.Println("String:", stringVar)

}
