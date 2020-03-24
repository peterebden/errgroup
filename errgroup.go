// Package errgroup implements a synchronisation mechanism similar to x/sync/errgroup,
// with the main difference being that it reports errors eagerly rather than waiting for all routines to terminate.
package errgroup

import "sync"

// A Group manages a group of goroutines running in parallel.
// A zero Group is *not* valid, use New instead.
// Groups should not be copied.
type Group struct {
	wg sync.WaitGroup
	err error
	ch chan struct{}
	mutex sync.Mutex
}

// New creates a new Group.
func New() *Group {
	return &Group{
		ch: make(chan struct{}),
	}
}

// Go calls the given function in a new goroutine.
// The first such function to return a non-nil error will cancel the group; its error will be returned by Wait.
func (g *Group) Go(f func() error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		if err := f(); err != nil {
			g.Cancel(err)
		}
	}()
}

// Wait waits for the first call to Go to return with an error, or for them all to return successfully,
// whichever happens first.
// If Wait returns an error, subsequent calls always return the same error.
func (g *Group) Wait() error {
	ch := make(chan struct{})
	go func() {
		g.wg.Wait()
		close(ch)
	}()
	select {
	case <-g.ch:
	case <-ch:
	}
	return g.err
}

// Cancel cancels this error group immediately with the given error, which will be
// returned from Wait() as though it had been returned from a call to Go().
func (g *Group) Cancel(err error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	if g.err == nil {
		g.err = err
		close(g.ch)
	}
}
