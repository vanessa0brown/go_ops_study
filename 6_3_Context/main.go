package main

import (
	"context"
	"fmt"
	"time"
)

func runGoroutines(ctx context.Context) {
	f := func(idx int, ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Printf("Горутина #%v завершила работу\n", idx)
		}
	}

	for i := range 5 {
		go f(i, ctx)
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	runGoroutines(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

}
