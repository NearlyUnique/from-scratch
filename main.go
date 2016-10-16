package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	blog struct {
		name string
	}
	blogReader interface {
		collect(result chan string)
	}
)

func main() {
	t0 := time.Now()
	rand.Seed(t0.Unix())
	ch := make(chan string)
	blogs := getBlogs()

	for _, blog := range blogs {
		go blog.collect(ch)
	}

	for _ = range blogs {
		result := <-ch
		fmt.Println(result)
	}

	fmt.Printf("Done in %v\n", time.Since(t0))
}

func (b blog) collect(ch chan string) {
	delay := rand.Intn(500)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	ch <- fmt.Sprintf("Collecting %s waited(%d)...\n", b.name, delay)
}

func getBlogs() []blogReader {
	return []blogReader{
		blog{name: "golang.org"},
		blog{name: "medium.com"},
		blog{name: "xkcd.com"},
	}
}
