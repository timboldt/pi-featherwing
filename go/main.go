package main

import (
	"fmt"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	fmt.Println("hello")
	err := rpio.Open()
	fmt.Println(err)
}
