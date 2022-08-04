//multi Producer one Consumer Problem

package main

import (
	"fmt"
	"sync"
)

var messages = [][]string{
	{"The world itself's",
		"just one big hoax.",
		"Spamming each other with our",
		"running commentary of bullshit,",
	},
	{
		"masquerading as insight, our social media",
		"faking as intimacy.",
		"Or is it that we voted for this?",
		"Not with our rigged elections,",
	},
	{
		"but with our things, our property, our money.",
		"I'm not saying anything new.",
		"We all know why we do this,",
		"not because Hunger Games",
	},
	{
		"books make us happy,",
		"but because we wanna be sedated.",
	},
	{
		"Because it's painful not to pretend,",
		"because we're cowards.",
		"- Elliot Alderson",
		"Mr. Robot",
	},
}

var wg sync.WaitGroup

func produce(ch chan string, i int) {
	defer wg.Done()
	for _, val := range messages[i] {
		fmt.Printf("producer %d - %s\n", i, val)
		ch <- val
	}

}
func consume(ch chan string, done chan bool) {
	for val := range ch {
		fmt.Println(val)
	}
	done <- true
}
func main() {
	ch := make(chan string)
	done := make(chan bool)
	producers_count := 5

	for i := 0; i < producers_count; i++ {
		wg.Add(1)
		go produce(ch, i)

	}

	go consume(ch, done)
	wg.Wait()
	close(ch)
	<-done

}
