package main

import (
	"log"
	"login/configs"
	"login/wire"
	"strconv"
)

func main() {
	// 初始化配置
	configs.InitConfig()

	// 初始化应用
	r, err := wire.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize: %v", err)
	}

	// 从配置中获取端口号并启动服务器
	portStr := strconv.Itoa(configs.Conf.Server.Port)
	if err := r.Run(":" + portStr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
