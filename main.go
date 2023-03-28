package main

import (
	"fpt/bootstrap"
)

func main() {
	//初始化配置
	bootstrap.InitializeConfig()

	//启动服务器
	bootstrap.RunServer()
}
