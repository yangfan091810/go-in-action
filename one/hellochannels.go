package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main(){
	c := make(chan int)
	go printer(c)
	wg.Add(1)

	for i := 1; i < 10; i++ {
		c <- i
	}
	close(c)
	wg.Wait()
}

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("receiveed %d \n", i)
	}
	wg.Done()
}
