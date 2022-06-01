package main

import "fmt"

func main() {
  counter := 1 //init statement

  
  for counter <= 10 {
   fmt.Println("Perulangan ke ", counter)
   counter++ //post statement
  }

  for number := 1; number <= 10; number++{
    fmt.Println("Ini short Statement, Perulangan ke ", number)
  }


  slice := []string{"Randy", "Reza", "Alfin"}
  
  for i := 0; i < len(slice); i++{
    fmt.Println(slice[i])
  }

  //fast way, to takee all value with range
  for i, slic := range slice{ 
    fmt.Println("index", i, "=", slic)
  }

  clubs := map[string]string{
    "pemain"   : "salah",
    "club"     : "liverpool",
    "lapangan" : "Anfield",
  }

  for i, club := range clubs{
    fmt.Println("i", i, "=", club)
  }

}