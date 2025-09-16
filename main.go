package main

import (
	"cangqiong/global"
	"cangqiong/initialize"
)

func main() {
	initialize.InitLogger() //初始化日志
	defer global.SugarLogger.Sync()
	initialize.InitConfig() //初始化配置
	initialize.InitDB()     //初始化数据库
}
