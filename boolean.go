package main

import "fmt"

func main(){
  fmt.Println("Benar =", true)
  fmt.Println("Salah =", false)

  var nilai = 90
  var absen = 80

  // var lulusnilai = nilai > 80
  // var lulusabsen = absen > 80

  // var hasil = lulusnilai && lulusabsen
  // fmt.Println(hasil)

  fmt.Println(nilai > 75 && absen > 70)
}