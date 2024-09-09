package main

import (
	"fmt"
	"hello/greetings"
	"log"
)

func main() {
	//	message, err := greetings.Hello("eliseekn")
	names := []string{"eliseekn", "wrh1d3"}
	messages, err := greetings.MultipleHello(names)
	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
