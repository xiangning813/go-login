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
func InitDB() (*gorm.DB, error) {
	cfg := configs.Conf.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Dbname, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	return db, nil
}

// InitConfig 初始化配置
func InitConfig() {
	configs.InitConfig()
}

// ProviderSet 定义 Provider 集合
var ProviderSet = wire.NewSet(
	InitGinEngine,
	InitConfig,
	InitDB,
)

// InitApp 初始化应用
func InitApp(engine *gin.Engine, db *gorm.DB) *gin.Engine {

	fmt.Println(engine)
	fmt.Println(db)

	routes.RegisterUserRoutes(engine, db)
	/* 404 */
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
