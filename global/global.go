package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var SugarLogger *zap.SugaredLogger //日志
var Configer *Configers            //配置
var DB *gorm.DB                    //数据库

// 定义配置结构体类型
type Configers struct {
	AppName string      `yaml:"appName"`
	MySql   MySqlConfig `yaml:"mySql"`
}

type MySqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
