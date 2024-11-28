package user

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"login/configs"
	"login/models"
	"time"
)

// UserService 定义用户服务
type UserService struct {
	DB *gorm.DB
}

// NewUserService 创建一个新的用户服务实例
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

// LoginRequest 请求体结构
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login 实现用户登录逻辑
func (s *UserService) Login(loginReq *LoginRequest) (string, error) {
	var user models.User
	if err := s.DB.Where("username = ? AND password = ?", loginReq.Username, loginReq.Password).First(&user).Error; err != nil {
		return "", errors.New("用户名或密码错误")
	}

	// 生成JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间为24小时
	})

	// 使用密钥签名并生成token
	tokenString, err := token.SignedString([]byte(configs.Conf.JWT.Secret))
	if err != nil {
		return "", errors.New("无法生成JWT token")
	}

	return tokenString, nil
}
