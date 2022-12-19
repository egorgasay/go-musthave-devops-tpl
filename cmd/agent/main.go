package main

import (
	"fmt"
	"runtime"
)

func main() {
	var mem runtime.MemStats
	fmt.Println(mem)
}
