package main

import "fmt"

func main(){
	var arr = [5]int{1,2,3,4,5}
	changearr(arr)
	fmt.Println(arr)
	var slice = []int{1,2,3,4,5}
	changeslice(slice)
	fmt.Println(slice)
}

func changearr(a [5]int) int {
	for index,value := range a{
		a[index] = value*2
	}
	return 0;
}

func changeslice(b []int) int{
	for index,value := range b{
		b[index] = value*2
	}
	return 0;
}
