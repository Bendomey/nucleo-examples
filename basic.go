package main

import (
	"fmt"

	"github.com/Bendomey/nucleo-go/broker"
	"github.com/Bendomey/nucleo-go/nucleo"
)

func main() {
	fmt.Println("Hello World")
	nucleoBroker := broker.New(nucleo.Config{
		Namespace: "basic-example",
	})

	nucleoBroker.PublishServices()

	nucleoBroker.Start()

	nucleoBroker.Stop()
}
