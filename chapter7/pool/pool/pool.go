package pool

import (
	"errors"
	"log"
	"io"
	"sync"
)

type Pool struct {
	m sync.Mutex
	resource chan io.Closer
	factory func()( io.Closer, error )
	closed bool
}

var ErrPoolClose = errors.New("pool has closed")

func New(fn func()(io.Closer, error), size int) (*Pool, error){
	if size < 0 {
		return nil,errors.New("size error")
	}
	return &Pool{
		resource: make(chan io.Closer, size),
		factory: fn,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
		case r,ok := <-p.resource:
			log.Println("get resource")
			if !ok {
				return nil, ErrPoolClose
			}
			return r,nil
		default:
			log.Println("create resource")
			return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
		case p.resource <- r:
			log.Println("put resource success")
		default:
			log.Println("resource full")
			r.Close()
	}
}

func (p *Pool) Close(){
	p.m.Lock()
	p.m.Unlock()

	if p.closed {
		log.Println("pool has closed ")
		return
	}

	p.closed = true
	close(p.resource)
	for r := range p.resource {
		r.Close()
	}
}














