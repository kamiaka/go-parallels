package parallels_test

import (
	"context"
	"fmt"
	"time"

	parallels "github.com/kamiaka/go-parallels"
)

func Example() {
	ls := []string{"A", "B", "C", "D", "E", "F"}
	start := time.Now()

	parallels.Do(func(i int) error {
		time.Sleep(sec(i + 1))
		fmt.Printf("#%v: end %v, %v\n", i, ls[i], sinceSec(start))
		return nil
	}, len(ls))

	// Output:
	// #0: end A, 1s
	// #1: end B, 2s
	// #2: end C, 3s
	// #3: end D, 4s
	// #4: end E, 5s
	// #5: end F, 6s
}

func ExampleDo_withConcurrent() {
	ls := []string{"A", "B", "C", "D", "E", "F"}
	start := time.Now()

	parallels.Do(func(i int) error {
		time.Sleep(sec(i + 1))
		fmt.Printf("#%v: end %v, %v\n", i, ls[i], sinceSec(start))
		return nil
	}, len(ls), parallels.Concurrent(2))

	// Output:
	// #0: end A, 1s
	// #1: end B, 2s
	// #2: end C, 4s
	// #3: end D, 6s
	// #4: end E, 9s
	// #5: end F, 12s
}

func ExampleDo_withContext() {
	ls := []string{"A", "B", "C", "D", "E", "F"}

	ctx, cancel := context.WithCancel(context.Background())
	ctxOpt, ctx := parallels.WithContext(ctx)

	start := time.Now()

	parallels.Do(func(i int) error {
		t := time.NewTimer(sec(i + 1))
		defer t.Stop()

		select {
		case <-ctx.Done():
			fmt.Printf("#%v: canceled %v, %v\n", i, ls[i], sinceSec(start))
		case <-t.C:
			fmt.Printf("#%v: end %v, %v\n", i, ls[i], sinceSec(start))
			if i == 2 {
				cancel()
			}
		}
		return nil
	}, len(ls), ctxOpt)

	// Unordered Output:
	// #0: end A, 1s
	// #1: end B, 2s
	// #2: end C, 3s
	// #5: canceled F, 3s
	// #3: canceled D, 3s
	// #4: canceled E, 3s
}

func ExampleDo_withError() {
	ls := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	start := time.Now()

	err := parallels.Do(func(i int) error {
		time.Sleep(sec(i + 1))
		fmt.Printf("#%v: end %v, %v\n", i, ls[i], sinceSec(start))
		if i >= 2 {
			return fmt.Errorf("#%v: an error occurred %v %v", i, ls[i], sinceSec(start))
		}
		return nil
	}, len(ls), parallels.Concurrent(2))
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// #0: end A, 1s
	// #1: end B, 2s
	// #2: end C, 4s
	// #3: end D, 6s
	// #2: an error occurred C 4s
}

// sec converts and returns a given integer into seconds.
// Note: tests emulates 1 / 10 seconds as 1 second.
func sec(i int) time.Duration {
	return time.Duration(i) * 100 * time.Millisecond
}

// sinceSec returns the time elapsed since t.
// Note: tests emulates 1 / 10 seconds as 1 second.
func sinceSec(t time.Time) time.Duration {
	return time.Since(t) / sec(1) * sec(1) * 10
}
