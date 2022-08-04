//multi Producer multi Consumer Problem

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

var pwg, cwg sync.WaitGroup

func produce(ch chan string, i int) {
	defer pwg.Done()
	for _, val := range messages[i] {
		fmt.Printf("producer %d - %s\n", i, val)
		ch <- val
	}

}
func consume(ch chan string, i int) {
	defer cwg.Done()
	for val := range ch {
		fmt.Printf("consumer - %d - %s", i, val)
	}

}
func main() {
	ch := make(chan string)

	producers_count := 5
	consumers_count := 3

	for i := 0; i < producers_count; i++ {
		pwg.Add(1)
		go produce(ch, i)

	}
	for i := 0; i < consumers_count; i++ {
		cwg.Add(1)
		go consume(ch, i)

	}

	pwg.Wait()
	close(ch)
	cwg.Wait()

}
