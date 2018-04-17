package main

import (
	"log"
	"os"
	"time"
	"./runner"
)

const timeout = 3 * time.Second

func main(){
	log.Println("work start")
	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
			case runner.ErrInterrupt:
				log.Printf(" interrupt error")
				os.Exit(1)
			case runner.ErrTimeout:
				log.Printf(" timeout error")
				os.Exit(2)
		}
	}
	log.Println("work finish")

}

func createTask() func(int) {
	return func(id int) {
		log.Printf("processor - task #%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
