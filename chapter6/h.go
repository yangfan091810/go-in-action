package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main(){
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	baton <- 1

	wg.Wait()

}

func Runner(baton chan int){
	var newRunner int
	runner := <-baton
	fmt.Printf("Running user is %d \n", runner)

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("runner %d is to over line \n", runner)

	if runner == 4{
		wg.Done()
		fmt.Println("running is finish")
		return
	}

	//fmt.Printf("runner %d is to over line \n", runner)

	newRunner = runner + 1

	fmt.Printf("runner %d exchange baton to runner %d \n", runner, newRunner)

	go Runner(baton)

	baton <- newRunner
}
