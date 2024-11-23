package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server   ServerConfig
	JWT      JWTConfig
	Database DatabaseConfig // 添加 DatabaseConfig 字段
}

type ServerConfig struct {
	Port int
}

type JWTConfig struct {
	Secret string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Dbname   string
	Sslmode  string
	Password string
}

var Conf *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // 还可以设置多个搜索路径

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	Conf = &Config{}
	if err := viper.Unmarshal(Conf); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
}
