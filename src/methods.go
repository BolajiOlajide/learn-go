package main

import (
	"fmt"
)

func main() {
	mp := messagePrinter{"foo"}
	mp.printMessage()
}

type messagePrinter struct {
	message string
}

func (message *messagePrinter) printMessage() {
	fmt.Println(message.message)
}
