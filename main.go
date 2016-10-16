package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t0 := time.Now()
	rand.Seed(t0.Unix())

	for i := 0; i < 3; i++ {
		delay := rand.Intn(500)
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Printf("Collecting %d ...\n", i)
	}

	fmt.Printf("Done in %v\n", time.Since(t0))
}
