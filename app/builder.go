package main

import (
	"os"

	"github.com/rock2z/tradebot/internal/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
	level := zap.DebugLevel
	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = zapcore.TimeEncoderOfLayout(util.DefaultLogLayout)
	fileEncoder := zapcore.NewJSONEncoder(pe)
	consoleEncoder := zapcore.NewConsoleEncoder(pe)
	f, _ := os.Create("log/data.log")

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(f), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
}
