package main

import (
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	words := strings.Split(msg, " ")
	var newWord string

	for _, w := range words {
		
		for i, v := range strings.Split(w, "") {
			newWord += strings.Repeat(v, i+1)
		}
		newWord += " "
	}
	print(newWord)
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}