package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("This is linux AMD 64")
	fmt.Println(runtime.GOARCH)
}
