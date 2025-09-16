package initialize

import (
	"cangqiong/global"
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	// 先初始化日志系统
	InitLogger()
	defer global.SugarLogger.Sync()

	// 再初始化配置
	InitConfig()
	fmt.Println(global.Configer.MySql.Host)
}
