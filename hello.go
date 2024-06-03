package main

import (
    "fmt"

    "example/greetings"
)

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
	message := greetings.Hello("Gladys")
    fmt.Println(message)
}