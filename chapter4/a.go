package main

import "fmt"

func main(){
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	fmt.Println(slice)
	fmt.Println(newSlice)
	newSlice1 := append(newSlice, 60)
	fmt.Println(newSlice1)
	fmt.Println(newSlice)
	fmt.Println(slice)
}
