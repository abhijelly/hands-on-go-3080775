// challenges/types-primitive/begin/main.go
package main

import (
	"fmt"
)

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val string = "global"

// val = "global"

func main() {
	// create a local variable "val" and assign it an int value
	val := 21

	// print the value of the local variable "val"
	//fmt.Println(val)
	fmt.Printf("%T, local val = %v\n", val, val)
	// print the value of the package-level variable "val"
	// create a function outside the scope of main
	// then local val will not be referenced
	printPackageVal()

	// update the package-level variable "val" and print it again
	updatePackageVal("hello world")
	printPackageVal()

	// print the pointer address of the local variable "val"
	fmt.Println(&val)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 42
	fmt.Println(val)
}

func printPackageVal() {
	fmt.Println(val)
}

func updatePackageVal(newVal string) {
	val = newVal
}
