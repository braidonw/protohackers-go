package main

import (
	"sync"

	primetime "github.com/braidonw/protohackers-go/internal/prime_time"
	"github.com/braidonw/protohackers-go/internal/smoketest"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		smoketest.Run()
	}()

	go func() {
		defer wg.Done()
		primetime.Run()
	}()

	wg.Wait()
}
