package main

import (
	"fmt"
	"github.com/hongjun500/mall-go/internal/conf"
)

func main() {
	fmt.Println("hello, mall-go")

	_, _ = conf.InitMySQLConn()

}
