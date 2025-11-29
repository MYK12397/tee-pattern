package main

import (
	"context"
	"fmt"
	"sync"
)

func logger(ctx context.Context, inpChan <-chan any, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("logger shutting down")
			return
		case val, ok := <-inpChan:
			if !ok {
				fmt.Println("logger channel closed")
				return
			}
			fmt.Println("logger: ", val)
		}
	}
}
