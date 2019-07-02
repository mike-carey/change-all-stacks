package commands

import (
	"sync"
)

type AsyncErrorPool struct {
	pool []error
	mutex sync.Mutex
}

func NewAsyncErrorPool() *AsyncErrorPool {
	return &AsyncErrorPool{
		pool: make([]error, 0),
		mutex: sync.Mutex{},
	}
}

func (p *AsyncErrorPool) Add(err error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.pool = append(p.pool, err)
}

func (p *AsyncErrorPool) Pool() []error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.pool
}

func (p *AsyncErrorPool) Len() int {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return len(p.pool)
}
