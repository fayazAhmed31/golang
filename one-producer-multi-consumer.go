//one Producer multi Consumer Problem

package main

import (
	"fmt"
	"sync"
)

var messages = []string{
	"The world itself's",
	"just one big hoax.",
	"Spamming each other with our",
	"running commentary of bullshit,",
	"masquerading as insight, our social media",
	"faking as intimacy.",
	"Or is it that we voted for this?",
	"Not with our rigged elections,",
	"but with our things, our property, our money.",
	"I'm not saying anything new.",
	"We all know why we do this,",
	"not because Hunger Games",
	"books make us happy,",
	"but because we wanna be sedated.",
	"Because it's painful not to pretend,",
	"because we're cowards.",
	"- Elliot Alderson",
	"Mr. Robot",
}

var wg sync.WaitGroup

func produce(ch chan string) {
	defer wg.Done()
	for _, val := range messages {
		ch <- val
	}
	close(ch)
}
func consume(ch chan string, i int) {
	defer wg.Done()
	for val := range ch {
		fmt.Printf("consumer %d - %s\n", i, val)
	}
}
func main() {
	ch := make(chan string)
	consumers_count := 5
	wg.Add(1)
	go produce(ch)
	for i := 0; i < consumers_count; i++ {
		wg.Add(1)
		go consume(ch, i)
	}

	wg.Wait()
}
