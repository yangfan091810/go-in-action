package main

import "fmt"

func main(){
	slice := []int{10,20,30,40}
	for index, value := range slice {
		fmt.Printf(" value: %d value address %X sliceitem address %X\n", value, &value, &slice[index] )
	}
}
