package parallels

import (
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

type group struct {
	*errgroup.Group
	parallels  int
	concurrent int
}

// Do executes function in parallel n times.
func Do(fn func(i int) error, n int, opts ...Option) error {
	g := &group{
		Group:      &errgroup.Group{},
		parallels:  n,
		concurrent: n,
	}

	for _, opt := range opts {
		opt(g)
	}

	return g.Do(fn)
}

func (g *group) Do(f func(i int) error) error {
	ch := make(chan int, g.concurrent)

	c := new(int64)

	var skip bool

	for i := 0; i < g.parallels; i++ {
		g.Go(func() error {
			ch <- 1
			defer func() {
				<-ch
			}()

			if skip {
				return nil
			}

			if err := f(int(atomic.AddInt64(c, 1) - 1)); err != nil {
				skip = true
				return err
			}

			return nil
		})
	}

	return g.Wait()
}
