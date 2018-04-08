package main

import "fmt"

func main(){
	slice := [][]int{{10}, {100,200}}
	slice[0] = append(slice[0], 30, 50, 60)
	for _,value := range slice {
		//fmt.Println(value)
		for _,val := range value {
			fmt.Println(val)
		}
	}
}
