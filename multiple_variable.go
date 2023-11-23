package main

import (
	"fmt"
)

func main() {
  //In Go, it is possible to declare multiple variables in the same line.
  var a, b, c, d int = 1, 3, 5, 7

  //If the type keyword is not specified, you can declare different types of variables in the same line.
  var a1, b1 = 6, "Hello"
  c1, d1 := 7, "World!"

  // Multiple variable declarations can also be grouped together into a block for greater readability.
  var (
    a2 int
    b2 int = 1
    c2 string = "hello"
  )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)

  fmt.Println(a1)
  fmt.Println(b1)
  fmt.Println(c1)
  fmt.Println(d1)

  fmt.Println(a2)
  fmt.Println(b2)
  fmt.Println(c2)
}