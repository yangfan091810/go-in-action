package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){

	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start")
	
	go func(){
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func(){
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	wg.Wait()
	fmt.Println()
	fmt.Println("finish")
}
