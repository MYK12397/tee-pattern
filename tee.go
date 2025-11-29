package main

import (
	"context"
	"sync"
)

func Tee(ctx context.Context, inpChan <-chan any, numOutputs int) []<-chan any {
	outputs := make([]chan any, numOutputs)
	for i := range outputs {
		outputs[i] = make(chan any)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer func() {
			for _, ch := range outputs {
				close(ch)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-inpChan:
				if !ok {
					return
				}

				//sending to all output channels
				var sendWg sync.WaitGroup
				for _, ch := range outputs {
					sendWg.Add(1)
					go func(c chan any) {
						defer sendWg.Done()
						select {
						case <-ctx.Done():
							return
						case c <- val:
						}
					}(ch)
				}

				sendWg.Wait()
			}
		}
	}()

	results := make([]<-chan any, numOutputs)
	for i, ch := range outputs {
		results[i] = ch
	}
	return results
}
