package main

import (
	"sync"

	"github.com/hoangnguyen1247/quickstart-lib-go/quickstart"
)

var wg sync.WaitGroup
var needToCreateANewGoroutine bool = true

func main() {
	// Create at least 1 goroutine
	wg.Add(1)
	go f()

	quickstart.Hello()

	go forever()
	wg.Wait()
}

// Conceptual code
func forever() {
	for {
		if needToCreateANewGoroutine {
			wg.Add(1)
			go f()
		}
	}
}

func f() {
	// When the termination condition for this goroutine is detected, do:
	wg.Done()
}
