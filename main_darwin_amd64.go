package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("This is darwin AMD 64")
	fmt.Println(runtime.GOARCH)
}
