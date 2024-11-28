package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"login/services/user"
	"net/http"
)

// RegisterUserRoutes 注册用户相关的路由
func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) *gin.Engine {
	user := user.NewUserService(db)
	userGroup := r.Group("/user")
	{

	}
	return r
}
