package main

import "fmt"

func main() {
  name := "Anugerah"

  if name == "Randy"{
    fmt.Println(" Hello Randy")
  } else if name == "Anugerah"{
    fmt.Println(" Hello Anugerah")
  }else{
    fmt.Println("Hello Anonymous")
  }

  //ifelse dengan short statement
  if length:=len(name); length > 8{
    fmt.Println("terlalu panjang")
  } else {
    fmt.Println("udah pas")
  }
  
}