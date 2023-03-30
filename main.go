package main

import (
	"fmt"
	"fpt/bootstrap"
	"fpt/global"
)

func main() {
	//初始化配置
	bootstrap.InitializeConfig()
	fmt.Println(global.Ftpserver)

	//初始化数据库
	bootstrap.InitializeDB()

	//程序结束前干的事情
	defer bootstrap.DeferThings()

	//启动服务器
	bootstrap.RunServer()
}
