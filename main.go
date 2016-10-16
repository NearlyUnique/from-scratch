package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const blogs = 3
	t0 := time.Now()
	rand.Seed(t0.Unix())
	ch := make(chan string)

	for i := 0; i < blogs; i++ {
		go collect(ch, i)
	}

	for i := 0; i < blogs; i++ {
		result := <-ch
		fmt.Println(result)
	}
	fmt.Printf("Done in %v\n", time.Since(t0))
}

func collect(ch chan string, index int) {
	delay := rand.Intn(500)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	ch <- fmt.Sprintf("Collecting %d waited(%d)...\n", index, delay)
}
