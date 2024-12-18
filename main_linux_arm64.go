package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("This is linuux ARM 64")
	fmt.Println(runtime.GOARCH)
}
