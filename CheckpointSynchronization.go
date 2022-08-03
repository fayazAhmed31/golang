//Illustration of Checkpoint Synchronization

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	workers = []string{"A", "B", "C", "D", "E"}
	wg      sync.WaitGroup
)

func work(worker string) {
	defer wg.Done()
	fmt.Println(worker, "started working")
	time.Sleep(time.Second)
	fmt.Println(worker, "completed work")
}

func main() {
	fmt.Println("Hello")
	for j := 0; j < 4; j++ {
		fmt.Printf("Cycle %d started\n", j)
		for i := 0; i < len(workers); i++ {
			wg.Add(1)
			go work(workers[i])
		}
		wg.Wait()
		fmt.Printf("Cycle %d completed\n", j)
	}

}
