package main

import (
	"fmt"
)

func main() {
  var a bool = true     // Boolean
  var b int = 5         // Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hi!"  // String

  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float:   ", c)
  fmt.Println("String:  ", d)


  //Boolean Data Type
  var b1 bool = true // typed declaration with initial value
  var b2 = true // untyped declaration with initial value
  var b3 bool // typed declaration without initial value 
  b4 := true // untyped declaration with initial value

  fmt.Println(b1)
  fmt.Println(b2)
  fmt.Println(b3)
  fmt.Println(b4)


// Integer Data Types:
//Signed integers, declared with one of the int keywords, can store "both positive and negative" values.
var x1 int = 500
  var y1 int = -4500
  fmt.Printf("Type: %T, value: %v", x1, x1)
  fmt.Printf("Type: %T, value: %v", y1, y1)

  //Unsigned Integers
 //Unsigned integers, declared with one of the uint keywords, can only store "non-negative" values:
 var x uint = 500
  var y uint = 4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}