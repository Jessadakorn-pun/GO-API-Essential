package mikelopster

import "fmt"

func SayTest() {
	saytest()
	fmt.Println("TEST")
}

func saytest() { // private function
	fmt.Println("test")
}
