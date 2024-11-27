package wire

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"login/configs"
	"login/models"
	"login/routes"
)

// InitGinEngine 初始化 Gin 引擎
func InitGinEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}

// InitDB 初始化数据库连接
func InitDB(cfg configs.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Dbname, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

// InitConfig 初始化配置
func InitConfig() *configs.Config {
	configs.InitConfig()
	return configs.Conf
}

// ProviderSet 定义 Provider 集合
var ProviderSet = wire.NewSet(
	InitGinEngine,
	InitConfig,
	InitDB,
)

// InitApp 初始化应用
func InitApp(engine *gin.Engine, db *gorm.DB) *gin.Engine {
	// 执行数据库迁移
	db.AutoMigrate(&models.User{})

	// 注册路由
	routes.RegisterUserRoutes(engine, db)

	// 设置404路由
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Url!"})
	})

	return engine
}

// Initialize 生成的 Provider 初始化函数
func Initialize() (*gin.Engine, error) {
	wire.Build(
		ProviderSet,
		InitApp,
	)
	return &gin.Engine{}, nil
}
