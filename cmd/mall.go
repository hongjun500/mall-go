package main

import (
	"fmt"
	"github.com/hongjun500/mall-go/internal/initialize"
)

func main() {
	fmt.Println("hello, mall-go")

	initialize.StartUp()
	// initialize.GinEngine.GinEngine.Run(":8080")
}
