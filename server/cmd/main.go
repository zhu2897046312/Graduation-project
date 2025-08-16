package main

import (
	"log"
	"server/config"
	"server/repository"
	"server/router"
	"server/service"
	"fmt"
	"github.com/gin-gonic/gin"
)
func main(){
	// 初始化配置
	if err := config.Init(); err != nil {
		log.Fatalf("配置初始化失败: %v", err)
	}
	config.InitMySQLDB()
	config.InitRedis()
	
	// 创建仓储工厂
    repoFactory := repository.NewRepositoryFactory(config.DB)

	// 创建服务工厂
    baseService := service.NewService(repoFactory)
    serviceFactory := service.NewServiceFactory(baseService,config.RedisClient)

	// 创建 Gin 引擎
    r := gin.Default()

	router.SetupRouter(serviceFactory)
	// 启动服务器
    port := config.GlobalConfig.Server.Port
    log.Printf("服务器启动在 :%d", port)
    log.Fatal(r.Run(fmt.Sprintf(":%d", port)))
}