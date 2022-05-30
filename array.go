package main

import "fmt"

func main() {
  var names [3] string
  names[0] = "M."
  names[1] = "Randy"
  names[2] = "Anugerah"

  fmt.Println(names)

  var values = [3] int{
    90,
    80,
    50,
  }
  fmt.Println(values[2])

  fmt.Println(len(names))
  fmt.Println(len(values))
}