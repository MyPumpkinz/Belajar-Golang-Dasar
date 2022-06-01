package main

import "fmt"

func sayHello() {
  name := "Randy"
  fmt.Println("Hello ", name)
}

func main() {
  for i := 0 ; i <= 10; i++{
     sayHello() 
  }
  
}