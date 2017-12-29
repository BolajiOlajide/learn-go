package main

import (
	"fmt"
)

func main() {
	emp := enhancedMessagePrinter{} // or enhancedMessagePrinter{messagePrinter{"foo"}} - not advisable though
	// because it requires user of the base class to understand the internal workings of the base class
	emp.message = "foo"
	emp.printMessage()
}

type messagePrinter struct {
	message string
}

func (message *messagePrinter) printMessage() {
	fmt.Println(message.message)
}

type enhancedMessagePrinter struct {
	messagePrinter
}
