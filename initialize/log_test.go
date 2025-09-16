package initialize

import (
	"cangqiong/global"
	"testing"
)

func TestInitLogger(t *testing.T) {
	InitLogger()
	defer global.SugarLogger.Sync()
	global.SugarLogger.Debugf("test debug")
	global.SugarLogger.Infof("test info")
	global.SugarLogger.Errorf("test error")
	global.SugarLogger.Warnf("test warn")
	global.SugarLogger.Fatalf("test fatal")
	global.SugarLogger.Panicf("test panic")
}
