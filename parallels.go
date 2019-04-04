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

	cnt := new(int64)

	errCnt := new(int32)

	for i := 0; i < g.parallels; i++ {
		g.Go(func() (err error) {
			ch <- 1
			defer func() {
				if r := recover(); r != nil {
					err = &Panic{
						Value: r,
					}
				}
				<-ch
			}()

			if atomic.LoadInt32(errCnt) != 0 {
				return nil
			}

			if err := f(int(atomic.AddInt64(cnt, 1) - 1)); err != nil {
				atomic.AddInt32(errCnt, 1)
				return err
			}

			return nil
		})
	}

	return g.Wait()
}
