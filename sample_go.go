package main

import (
	"errors"
	"fmt"
)

// Basic Data Types
// By default they take on the null values
var boolean bool         // false
var myString string      // ""
var signedInteger int    // 0
var unsignedInteger uint // 0

// Declaring 2 variables of the same data type at the same time
var value1, value2 int
var me, myFriend string

// Initializing a variable
var myAge int = 20
var myName string = "God"

// Integers can either be signed or unsigned with different bit size
// Calculating acceptable range of values for signed Int is -2^n to (2^n)-1 and for an unsigned Int 0 to (2^n)-1
// The default int or uint are 32-bits wide on 32-bit systems and 64-bits wide on 64-bit system, so it's safer to declare the width
var myInt8 int8   // Can hold valued from -128 to 127
var myInt16 int16 // -2^16 to (2^16)-1
var myInt32 int32
var myInt64 int64

var myUnsignedInt8 uint8 // Can hold valued from 0 to 255 i.e, 2^8
var myUnsignedInt16 uint16
var myUnsignedInt32 uint32
var myUnsignedInt64 uint64

var myRune rune // same as int32
var myByte byte // same as uint8

// Floats
var myFloat32 float32 // Precision upto 7 decimal places
var myFloat64 float64 // Precision upto 15 to 17 decimal places

// Constants
const internalMarks uint8 = 20
const externalMarks uint8 = 71

// User Define Types
type Person struct {
	Name   string
	Age    uint8
	Status bool
}

// Function
func (p Person) Jump() {
	fmt.Println("God jumped")
}

// Syntax: funct (object DataType) functioName() (return) {}
func (p Person) driveCar() (bool, error) {
	if p.Age < 18 {
		return false, errors.New("Too young to drive")
	}
	return true, nil
}

func main() {
	fmt.Println("Hello World")

	// Simple Operator - Note: declaring a variable inside the main function and not using it will throw an error
	var totalMarks uint8 = internalMarks + externalMarks
	fmt.Println(totalMarks)

	var isItGreater bool = externalMarks > internalMarks
	fmt.Println(isItGreater)

	// String concatenation
	mySchool := "MCC Campus" // Short assignment operator and assumes the string data type, notice that even the var keyword is not necessary
	schoolBoard := "Matriculation"
	fmt.Println(mySchool + " " + schoolBoard)

	// Arrays
	var primes = [6]int{2, 3, 5, 7, 9, 11}
	odd := [6]int{1, 3, 5, 7, 9, 11}
	fmt.Println(primes[0])
	fmt.Println(odd[3])

	// Slices = Arrays without having to define the size
	even := []int{2, 4, 6, 8}
	fmt.Println(even[3])

	// Maps - Key value pairs, similar to hashmap in Java
	myMap := make(map[string]int)
	myMap["God"] = 20
	myMap["Godness"] = 21
	myMap["Evil"] = 15
	fmt.Println(myMap)

	// Loops, note: Go doesnt have a while loop
	// The count for till how much a loop should run is an int and so specifying unsigned and or bit width will throw an error
	//const count uint64 = 5 // throws an error
	const count int = 6
	sum := 0
	for i := 0; i < count; i++ {
		sum += 1
		fmt.Println(sum)
	}

	// Looping over slices
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		fmt.Println(i, v)
	}

	// Creating a structure object
	GG := Person{
		Name:   "God",
		Age:    20,
		Status: true,
	}
	GG.Jump()
	canDrive, err := GG.driveCar()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Can drive? ", canDrive)
	}

}
