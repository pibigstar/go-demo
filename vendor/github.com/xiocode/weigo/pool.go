/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          pool.go
 * Description:   generic pool
 */

package weigo

import (
	"errors"
	"fmt"
	"sync"
)

// Factory is a function to create new connections.
type Factory func() (interface{}, error)

// Pool allows you to use a pool of net.Conn connections.
type Pool struct {
	mu       sync.Mutex
	elements chan interface{} // storage for interface{}
	factory  Factory          // net.Conn generator
}

// New returns a new pool with an initial capacity and maximum capacity.
// Factory is used when initial capacity is greater than zero to fill the
// pool.
func NewConnPool(size, capacity int, factory Factory) (*Pool, error) {
	if size <= 0 || capacity <= 0 || size > capacity {
		return nil, errors.New("invalid capacity settings")
	}

	pool := &Pool{
		elements: make(chan interface{}, capacity),
		factory:  factory,
	}

	// create initial connections, if something goes wrong,
	// just close the pool error out.
	for i := 0; i < size; i++ {
		element, err := factory()
		if err != nil {
			return nil, fmt.Errorf("factory is not able to fill the pool: %s", err)
		}
		pool.elements <- element
	}

	return pool, nil
}

func (p *Pool) getElements() chan interface{} {
	p.mu.Lock()
	elements := p.elements
	p.mu.Unlock()
	return elements
}

// Get returns a new connection from the pool. After using the connection it
// should be put back via the Put() method. If there is no new connection
// available in the pool, a new connection will be created via the Factory()
// method.
func (p *Pool) Get() (interface{}, error) {
	elements := p.getElements()
	if elements == nil {
		return nil, errors.New("pool is closed")
	}

	select {
	case element := <-elements:
		if element == nil {
			return nil, errors.New("pool is closed")
		}
		return element, nil
	default:
		return p.factory()
	}
}

// Put puts an existing connection into the pool. If the pool is full or
// closed, conn is simply closed. A nil conn will be rejected. Putting into a
// destroyed or full pool will be counted as an error.
func (p *Pool) Put(element interface{}) error {
	if element == nil {
		return errors.New("connection is nil. rejecting")
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.elements == nil {
		return errors.New("pool is closed")
	}

	select {
	case p.elements <- element:
		return nil
	default:
		return errors.New("pool is full")
	}
}

// Close closes the pool and all its connections. After Close() the
// pool is no longer usable.
func (p *Pool) Close() {
	p.mu.Lock()
	elements := p.elements
	p.elements = nil
	p.factory = nil
	p.mu.Unlock()

	if elements == nil {
		return
	}

	close(elements)
	for _ = range elements {
	}
}

// MaximumCapacity returns the maximum capacity of the pool
func (p *Pool) MaximumCapacity() int { return cap(p.getElements()) }

// CurrentCapacity returns the current capacity of the pool.
func (p *Pool) CurrentCapacity() int { return len(p.getElements()) }
