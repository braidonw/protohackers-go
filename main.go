package main

import (
	"sync"

	"github.com/braidonw/protohackers-go/internal/smoketest"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		smoketest.Run()
	}()

	wg.Wait()
}
