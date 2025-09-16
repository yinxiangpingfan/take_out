package initialize

import (
	"cangqiong/global"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 初始化日志

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.ErrorLevel)

	logger := zap.New(core)
	global.SugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./log/cangqiong.log")
	return zapcore.AddSync(file)
}
