package main

import (
	"flag"
	"fmt"
)

var name string

func main() {
	helloFunc()
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}