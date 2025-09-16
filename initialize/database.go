package initialize

import (
	"cangqiong/global"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

// 初始化数据库
func InitDB() {
	// 初始化数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.Configer.MySql.Username,
		global.Configer.MySql.Password,
		global.Configer.MySql.Host,
		global.Configer.MySql.Port,
		global.Configer.MySql.Database,
	)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.SugarLogger.Fatalf("数据库连接失败: %v", err)
	}
}
