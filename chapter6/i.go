package main

import(
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad = 10
)

var wg sync.WaitGroup

func init(){
	rand.Seed(time.Now().Unix())
}

func main(){
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)

	for i := 1; i <= numberGoroutines; i++ {
		go Work(tasks, i)
	}

	for k := 1; k <= taskLoad; k++ {
		tasks <- fmt.Sprintf("task : %d ", k)
	}

	close(tasks)

	wg.Wait()
}

func Work(tasks chan string, i int){
	defer wg.Done()
	for {
		task,ok := <-tasks
		if !ok {
			fmt.Printf("worker %d shut down \n", i)
			return
		}
		fmt.Printf("worker %d received %s \n", i, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("worker %d finish %s \n", i, task)
	}
}





