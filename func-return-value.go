package main

import "fmt"

func getHello(name string) string {
  if name == ""{
    return "Hello bro!!"
  }else{
    return "Hello" + name
  }
}

func main() {
  result := getHello("Randy")
  fmt.Println(result)
  fmt.Println(getHello(""))
}