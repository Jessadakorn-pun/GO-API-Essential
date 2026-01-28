package main

import (
	"fmt"

	"github.com/Jessadakorn-pun/go-example/mikelopster"
	"github.com/google/uuid"
)

func main() {

	id := uuid.New()
	fmt.Println("Hello World")
	sayHello()
	fmt.Printf("%s \n", id)

	// imported function from other package
	mikelopster.SayHelloMikelopster()
	mikelopster.SayHelloMikelopster()
	mikelopster.SayTest()
}

func sayHello() {
	fmt.Printf("Hello")
}
