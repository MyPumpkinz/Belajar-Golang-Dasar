package main

import "fmt"

func main() {
  const fname string = "Randy"
  const lname = "Anugerah"
  const value = 100 //walau gak digunain gak error

  //fname = "tidak bisa diubah"

  fmt.Println(fname + " " + lname)

  const (
    age     = 12
    country = "Indonesia"
  )

  fmt.Println(age)
  fmt.Println(country)
}