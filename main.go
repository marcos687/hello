package main

import (
	"fmt"

	"github.com/marcos687/hello/mypkg"
)

func main() {
	fmt.Println("hello, modules")

	mypkg.PrintHello()
}
