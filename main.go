package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	t0 := time.Now()
	wg := sync.WaitGroup{}
	rand.Seed(t0.Unix())

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(index int) {
			delay := rand.Intn(500)
			time.Sleep(time.Duration(delay) * time.Millisecond)
			fmt.Printf("Collecting %d waited(%d)...\n", index, delay)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Printf("Done in %v\n", time.Since(t0))
}
