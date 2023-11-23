package main

import (
	"fmt"
)

var a1 int
var b1 int = 2
var c1 = 3


/*
a2 := 1   
./prog.go:13:4: syntax error: unexpected :=, expecting =   */

func main() {
  var stud1 string = "Jora" //type is string
  var stud2 = "Divi" //type is inferred
  x := 2 //type is inferred

  var student1 string
  student1 = "John"
  /*It is possible to assign a value to a variable after it is declared. This is helpful for cases the value is not initially known. */

  var a string
  var b int
  var c bool

  a1 = 1
  fmt.Println(a1)
  fmt.Println(b1)
  fmt.Println(c1)

  // fmt.Println(a2)

  fmt.Println(stud1)
  fmt.Println(stud2)
  fmt.Println(x)

  fmt.Println(student1)

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

/*Note: In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).

Note: It is not possible to declare a variable using :=, without assigning a value to it.

a is ""
b is 0
c is false   */