package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	people = []string{"A", "B", "C", "D", "E"}
	wg     sync.WaitGroup
)

type fork struct{ sync.Mutex }

type pihlosopher struct {
	name                string
	leftfork, rightfork *fork
}

func dine(ph *pihlosopher) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		ph.leftfork.Lock()
		ph.rightfork.Lock()
		fmt.Println(ph.name, "eating")
		time.Sleep(time.Second)
		fmt.Println(ph.name, "eating done")
		ph.leftfork.Unlock()
		ph.rightfork.Unlock()
		fmt.Println(ph.name, "thinking")
		time.Sleep(time.Second)
	}
}

func main() {
	forks := make([]*fork, 5)
	for i := 0; i < 5; i++ {
		forks[i] = new(fork)
	}
	for i := 0; i < len(people); i++ {
		ph := &pihlosopher{
			name:      people[i],
			leftfork:  forks[i],
			rightfork: forks[(i+1)%5],
		}
		wg.Add(1)
		go dine(ph)
	}
	wg.Wait()
}
