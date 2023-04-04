package main

import "goxs/util"

func main() {
	// 初始化日志
	util.GetLogger()
	defer util.SyncLog()

	// 启动 Gin 引擎
	util.StartServer(":8080")
}
