package parallels

import "fmt"

type Panic struct {
	Value interface{}
}

func (p *Panic) Error() string {
	return fmt.Sprintf("caught panic in goroutine: %v", p.Value)
}
