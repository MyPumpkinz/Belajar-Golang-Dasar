// MEMBUAT SLICE
// array[low:high]
// array[low:]
// array[:high]
// array[:]

package main

import "fmt"

func main() {
  months := [...]string{ //kalo kita gak tau kapasitas array bisa [...]
    "Januari",
    "Februari",
    "Maret",
    "April",
    "Mei",
    "Juni",
    "juli",
    "Agustus",
    "September",
    "Oktober",
    "November",
    "Desember",
  }
  
  slice2 := months[4:8]
  fmt.Println(slice2)
  fmt.Println(len(slice2))
  fmt.Println(cap(slice2))

  slice2 = append(slice2, "tambah data")
  fmt.Println(slice2)

  newSlice := make([]string, 3, 5)

  newSlice[0] = "randy"
  newSlice[1] = "anugeerah"
  newSlice[2] = "aassa"
  
  fmt.Println(newSlice)

  copySlice := make([]string, len(newSlice), cap(newSlice))
  copy(copySlice, newSlice)
  fmt.Println(copySlice)


  //hati-hati dengan penulisan array dan slice
  iniArray := [5]int{1,2,3,4,5}
  iniSlice := []int{1,2,3,4,5}

  fmt.Println(iniArray)
  fmt.Println(iniSlice)
  
}