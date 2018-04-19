package main

import (
	"sync"
	"io"
	"log"
	"math/rand"
	"time"
	"sync/atomic"
	"./pool"
)

const (
	maxGoroutines = 10
	pooledResources = 4
)

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close () error {
	log.Println("dbConn close :", dbConn.ID)
	return nil
}

var idCounter int32

func createConnect() (io.Closer, error){
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{id}, nil
}

func main(){
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p,err := pool.New(createConnect, pooledResources)
	if err != nil {
		log.Println(err)
	}
	for query := 0; query < maxGoroutines; query++ {
		go func(q int){
			performQueries(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()
	p.Close()
	log.Println("finish")
}


func performQueries(query int, p *pool.Pool){
	conn,err := p.Acquire()
	if err != nil {
		log.Println(err)
	}
	defer p.Release(conn)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}
