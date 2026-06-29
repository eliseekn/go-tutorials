package main

import (
	"fmt"
	"hello/greetings"
	"log"
)

func main() {
	message, err := greetings.Hello("eliseekn")
	if err != nil {
		log.Fatal(err)
	}

	messages, err := greetings.MultipleHello([]string{"eliseekn", "wrh1d3"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	for _, message := range messages {
		fmt.Println(message)
	}
}
