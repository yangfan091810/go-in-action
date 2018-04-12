package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start")

	go func(){
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char <= 'z'; char++ {
				fmt.Printf("%c", char)
			}
			fmt.Println()
		}
	}()
	
	go func(){
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char <= 'Z'; char++ {
				fmt.Printf("%c", char)
			}
			fmt.Println()
		}
	}()

	fmt.Println("wait goroute")
	wg.Wait()
	fmt.Println("goroute finish")
}
