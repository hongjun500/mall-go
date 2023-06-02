package main

import (
	"fmt"
	"github.com/hongjun500/mall-go/internal/initialize"
)

func main() {
	fmt.Println("hello, mall-go")

	ginEngine := initialize.StartUp()
	// 启动 gin 引擎并监听在 8080 端口
	ginEngine.Run(":8080")
}
