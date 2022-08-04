/one Producer one Consumer Problem

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
func consume(ch chan string) {
	defer wg.Done()
	for val := range ch {
		fmt.Println(val)
	}
}
func main() {
	ch := make(chan string)

	wg.Add(1)
	go produce(ch)
	wg.Add(1)
	go consume(ch)

	wg.Wait()
}
