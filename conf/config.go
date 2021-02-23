package conf

import "gopkg.in/ini.v1"

// Config 配置类
type Config struct {
	App AppConfig     `ini:"app"`
	DB  PostgreConfig `ini:"postgre"`
}

// AppConfig App 配置
type AppConfig struct {
	Port string `ini:"port"`
}

// PostgreConfig 数据库配置
type PostgreConfig struct {
	Host        string `ini:"host"`
	Port        int    `ini:"port"`
	User        string `ini:"user"`
	Password    string `ini:"password"`
	Database    string `ini:"database"`
	Pooling     bool   `ini:"pooling"`
	MinPoolSize int    `ini:"min_pool_size"`
	MaxPoolSize int    `ini:"max_pool_size"`
	LifeTime    int    `ini:"lifetime"`
}

// Conf 配置文件
var Conf = new(Config)

// InitConfig 初始化
func InitConfig(file string) error {
	return ini.MapTo(Conf, file)
}
