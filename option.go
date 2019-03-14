package parallels

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Option func(p *parallels)

func Parallels(num int) Option {
	return func(p *parallels) {
		p.parallels = num
	}
}

func WithContext(ctx context.Context) (Option, context.Context) {
	g, ctx := errgroup.WithContext(ctx)
	return func(p *parallels) {
		p.group = g
	}, ctx
}
