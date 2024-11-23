package routes

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"login/configs"
	"login/models"
	"net/http"
	"time"
)

// LoginRequest 请求体结构
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterUserRoutes 注册用户相关的路由
func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) *gin.Engine {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", func(c *gin.Context) {
			var loginReq LoginRequest
			if err := c.ShouldBindJSON(&loginReq); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "无效的输入"})
				return
			}

			var user models.User
			if err := db.Where("username = ? AND password = ?", loginReq.Username, loginReq.Password).First(&user).Error; err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
				return
			}

			// 生成JWT token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": user.Username,
				"exp":      time.Now().Add(time.Hour * 24).Unix(),
			})

			tokenString, err := token.SignedString([]byte(configs.Conf.JWT.Secret))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "无法生成JWT token"})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "登录成功",
				"token":   tokenString,
			})
		})
	}
	return r
}
