package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hei this is separate function ")
	}()
	fmt.Println("Hellow world !!")

	wg.Wait()
}
