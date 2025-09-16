package initialize

import (
	"cangqiong/global"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	global.Configer = &global.Configers{}
	getConf()
}

func getConf() {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		global.SugarLogger.Fatalf("获取项目执行路径失败，err：%v", err)
	}

	vip := viper.New()
	vip.AddConfigPath(path)     //设置读取的文件路径
	vip.SetConfigName("config") //设置读取的文件名
	vip.SetConfigType("yaml")   //设置文件的类型
	//尝试进行配置读取
	if err = vip.ReadInConfig(); err != nil {
		global.SugarLogger.Fatalf("读取配置文件失败，err：%v", err)
	}

	err = vip.Unmarshal(global.Configer)
	if err != nil {
		global.SugarLogger.Fatalf("解析配置文件失败，err：%v", err)
	}
}
