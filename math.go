package main

import "fmt"

func main() {

  var result = 10 + 10
  fmt.Println(result)

  a := 10
  b := 30
  c := a + b

  fmt.Println(c)

  var i = 10
  i += 10
  fmt.Println(i)

  var inc = 10
  inc ++
  fmt.Println(inc)

  name1 := "randy"
  name2 := "randy"

  var hasil = name1 == name2
  fmt.Println(hasil)

  value1 := 100
  value2 := 200
  fmt.Println(value1 > value2)
}