package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"login/configs"
	"login/models"
	"login/routes"
)

// InitConfig 初始化配置
func InitConfig() *configs.Config {
	configs.InitConfig()
	return configs.Conf
}

// InitDB 初始化数据库连接
func InitDB(cfg configs.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Port)

	fmt.Println("数据库字符串：【%s】:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

// InitGinEngine 初始化 Gin 引擎
func InitGinEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}

// InitApp 初始化应用
func InitApp(engine *gin.Engine, db *gorm.DB) *gin.Engine {
	// 执行数据库迁移
	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic(fmt.Sprintf("数据库迁移失败: %v", err))
	}

	// 注册路由
	routes.RegisterUserRoutes(engine, db)

	// 设置404路由
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Url!"})
	})

	fmt.Println("service[port:8999].......")

	return engine
}

func main() {
	// 初始化配置
	config := InitConfig()

	// 初始化数据库
	db, err := InitDB(config.Database)
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}

	// 初始化Gin引擎
	engine := InitGinEngine()

	// 初始化应用
	engine = InitApp(engine, db)

	// 启动Gin服务
	if err := engine.Run(":8999"); err != nil {
		fmt.Println("启动服务器失败:", err)
	}
}
