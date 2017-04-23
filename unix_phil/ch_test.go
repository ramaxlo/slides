package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	go writer(ch, &wg)
	go reader(0, ch, &wg)
	go reader(1, ch, &wg)
	wg.Add(3)

	wg.Wait()
}

func reader(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range ch {
		fmt.Printf("[%d] %d\n", id, i)
	}

	fmt.Printf("[%d] exit\n", id)
}

func writer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}

	close(ch)
}
