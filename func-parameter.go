package main

import "fmt"

func sayHelloTO (fname string, lname string){
  fmt.Println("Hello",fname, lname)
}

func main() {
  firstname := "Randy"
  sayHelloTO(firstname, "Anugerah")
  sayHelloTO("Alfin", "Reza")
}