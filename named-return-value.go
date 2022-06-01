package main

import "fmt"

func getCompleteName()(fname, mname, lname string){
  fname = "M."
  mname = "Randy"
  lname = "Anugerah"
  
  return //return, doesn't required to take var from parameter
}

func main() {
  fname, mname, lname := getCompleteName()
  fmt.Println(fname, mname, lname)
}