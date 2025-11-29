package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	inputChan := make(chan any)

	outputChans := Tee(ctx, inputChan, 2)

	var wg sync.WaitGroup

	wg.Add(2)
	go logger(ctx, outputChans[0], &wg)
	go metrics(ctx, outputChans[1], &wg)

	go func() {
		defer close(inputChan)
		for i := range 5 {
			select {
			case <-ctx.Done():
				fmt.Println("context cancelled")
				return
			case inputChan <- fmt.Sprintf("user-%d signed up using email", i+1):
				time.Sleep(150 * time.Millisecond)
			}
		}
	}()
	wg.Wait()
	fmt.Println("all consumers finished")
}
