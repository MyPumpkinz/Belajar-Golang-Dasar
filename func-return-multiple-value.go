package main

import "fmt"

func getFullName() (string, string, int){
  return "Randy", "Anugerah", 22
}

func getAddress() (string, string, int){
  return "Tirto", "Malang", 71418
}

func main() {
  fname, lname, age := getFullName()
  fmt.Println(fname, lname, age)

  //ignore return multiple value
  street, _, _ := getAddress()
  fmt.Println(street)
}