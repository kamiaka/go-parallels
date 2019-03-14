package parallels

import (
	"golang.org/x/sync/errgroup"
)

type parallels struct {
	repeat    int
	parallels int
	group     *errgroup.Group
}

func Do(fn func(i int) error, repeat int, opts ...Option) error {
	p := &parallels{
		repeat:    repeat,
		parallels: repeat,
		group:     &errgroup.Group{},
	}

	for _, opt := range opts {
		opt(p)
	}

	return p.Do(fn)
}

func (p *parallels) Do(fn func(i int) error) error {
	ch := make(chan int, p.parallels)

	for i := 0; i < p.repeat; i++ {
		p.group.Go(func() error {
			ch <- 1
			defer func() {
				<-ch
			}()

			return fn(i)
		})
	}

	return p.group.Wait()
}
