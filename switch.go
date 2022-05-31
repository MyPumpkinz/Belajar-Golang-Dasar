package main

import "fmt"


func main() {
	name := "Randy"

  switch name {
    case "Randy":
      fmt.Println("hello Randy")
    case "Anugerah":
      fmt.Println("hello Anugerah")
    default:
      fmt.Println("hello Anonymous")
  }
}