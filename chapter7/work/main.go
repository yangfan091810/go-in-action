package main

import (
	"sync"
	"time"
	"log"
	"./work"
)

var names = []string{
	"aaaa",
	"bbbb",
	"cccc",
	"dddd",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task(){
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main(){
	p := work.New(2)
	var wg sync.WaitGroup
	wg.Add(20 * len(names))
	for i := 0; i < 20; i++ {
		for _,name := range names {
			np := namePrinter{
				name:name,
			}
			go func (){
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()

	p.Shutdown()
}
