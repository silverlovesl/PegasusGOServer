package main

import (
	"fmt"
	"pegasus/server"
	"pegasus/utils"

	"gopkg.in/tylerb/graceful.v1"
)

func main() {
	fmt.Println("Start Server...")
	initServer()
}

func initServer() {

	///初始化服务器
	srv := server.InitServer()

	///开始运行
	fmt.Printf("The web service is working on port %s...\n", utils.PORT)
	graceful.Run(":"+utils.PORT, 0, srv)
}
