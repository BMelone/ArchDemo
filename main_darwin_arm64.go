package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("This is darwin ARM 64")
	fmt.Println(runtime.GOARCH)
}
