package main

import "fmt"

func main(){

  clubs := map[int]string{ //[int] adalah tipe data KEY dan string tipe data VALUE
    1 : "Liverpool",
    2 : "Kloop",
    3 : "Salah",
    4 : "manie",
  }
  
  fmt.Println(clubs)
  fmt.Println(clubs[3])

  delete(clubs, 4) //menghapus

  clubs2 := make(map[string]string)
  clubs2["pemain1"] = "Mane" //menambah
  clubs2["pemain2"] = "Salah" //menambah
  
  fmt.Println(clubs2["club"])
  fmt.Println(clubs2["pemain2"])
}