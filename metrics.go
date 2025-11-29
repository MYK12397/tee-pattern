package main

import (
	"context"
	"fmt"
	"sync"
)

func metrics(ctx context.Context, inpChan <-chan any, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("metrics: processed %d users before shutdown\n", count)
			return
		case val, ok := <-inpChan:
			if !ok {
				fmt.Printf("metrics: total processed %d items\n", count)
				return
			}
			count++
			fmt.Printf("metrics: user count = %d, value = %v\n", count, val)
		}
	}
}
