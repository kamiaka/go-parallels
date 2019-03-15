package parallels

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// Option is option to specify the bahavior of parallel execution.
type Option func(*group)

// Concurrent returns an Option to specify the number of concurrent executions.
func Concurrent(num int) Option {
	return func(g *group) {
		g.concurrent = num
	}
}

// WithContext returns a Option for using Context
// and an associated Context derived from ctx.
func WithContext(ctx context.Context) (Option, context.Context) {
	grp, ctx := errgroup.WithContext(ctx)
	return func(g *group) {
		g.Group = grp
	}, ctx
}
