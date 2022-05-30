package main

import "fmt"

func main() {
  var name string //harus deklarasikan tipe data karena belum dimasukkan data/value

  name = "M. Randy Anugerah"
  fmt.Println(name)

  name = "Anugerah Randy"
  fmt.Println(name)

  var age = 12 //TIDAK deklarasikan tipe data karena ada value didalamnya 
  fmt.Println(age)

  born := 2000 //TIDAK menggunakan var, penggantinya :=, tetapi hnya untuk deklarasi pertama
  fmt.Println(born)

  born = 1999
  fmt.Println(born) 

  var(
    fname = "M."
    mname = "Randy"
    lname = "Anugerah"
  )

  fmt.Println(fname + " " + mname + " " + lname)
}