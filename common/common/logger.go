package common

import (
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 全局日志对象
var Logger *zap.SugaredLogger

func init() {
	// 自动初始化执行
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		// 时间格式
		EncodeTime: zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		// 彩色日志级别
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),        // 日志从终端输出格式，适合人阅读
		zapcore.AddSync(colorable.NewColorableStdout()), // 输出到标准输出，并支持彩色终端显示；`AddSync` 将输出对象适配成 Zap 所需的接口。
		zapcore.DebugLevel,                              // 最低输出级别为 Debug, Debug < Info < Warn < Error < Fatal
	)

	// 创建日志记录器
	// zap.AddCaller()：显示打印日志的文件名和行号。
	// zap.AddCallerSkip(1)：额外跳过一层调用，避免显示当前函数的调用信息
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Logger = logger.Sugar()
}

// 对外封装日志函数
func Info(args ...any) {
	Logger.Info(args...)
}

func Infof(format string, args ...any) {
	Logger.Infof(format, args...)
}

func Error(args ...any) {
	Logger.Error(args...)
}

func Errorf(format string, args ...any) {
	Logger.Errorf(format, args...)
}

func Warn(args ...any) {
	Logger.Warn(args...)
}

func Warnf(format string, args ...any) {
	Logger.Warnf(format, args...)
}

func Debug(args ...any) {
	Logger.Debug(args...)
}

func Debugf(format string, args ...any) {
	Logger.Debugf(format, args...)
}

func Fatal(args ...any) {
	Logger.Fatal(args...)
}

func Fatalf(format string, args ...any) {
	Logger.Fatalf(format, args...)
}
