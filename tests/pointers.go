package main

import "fmt"

func Pointers() {
	i := 2
	fmt.Println(i)
	p := &i
	fmt.Println(*p)
	*p = 45
	fmt.Println(*p)
}
