package main

import (
	"fmt"
)

func returnFunc() func() {
	return func() {
		fmt.Printf("I'm a function")
	}
}

func main() {
	f := returnFunc()
	f()
}