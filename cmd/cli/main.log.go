package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Failed to fetch URL.")

	// logger := zap.NewExample()
	// logger.Info("Failed to fetch URL.")
	// // development
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Failed to fetch URL.")
	// // production
	// logger, _ = zap.NewProduction()
	// logger.Info("Failed to fetch URL.")

	encode := getEncoderLog()
	sync := getWriteSyncerLog()
	core := zapcore.NewCore(encode, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Infor log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))

}
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSyncerLog() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_RDONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
