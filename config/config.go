package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// AppConfig 应用程序配置
type AppConfig struct {
	Name    string
	Version string
	Port    int
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string
	Format     string
	Output     string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

// Config 全局配置
type Config struct {
	App AppConfig
	Log LogConfig
}

// GlobalConfig 全局配置实例
var GlobalConfig Config

func init() {
	// 设置配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 设置默认配置
	setDefaults()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s. Using default configuration.\n", err)
	}

	// 解析配置到结构体
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		fmt.Printf("Unable to decode into struct, %v. Using default configuration.\n", err)
	}
}

func setDefaults() {
	viper.SetDefault("app.name", "gin-app")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.port", 9000)
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.format", "json")
	viper.SetDefault("log.output", "stdout")
	viper.SetDefault("log.filename", "app.log")
	viper.SetDefault("log.maxSize", 10)
	viper.SetDefault("log.maxBackups", 3)
	viper.SetDefault("log.maxAge", 28)
	viper.SetDefault("log.compress", true)
}
