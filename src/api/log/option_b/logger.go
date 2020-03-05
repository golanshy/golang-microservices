package option_b

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",
			LevelKey:   "level",
			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeCaller:zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout", "/tmp/logs"},
		ErrorOutputPaths: nil,
		InitialFields:    nil,
	}

	var err error
	Log, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}

func Field(key string, valur interface{}) zap.Field {
	return zap.Any(key, valur)
}

func Debug(msg string, tags ...zap.Field) {
	Log.Debug(msg, tags...)
	Log.Sync()
}

func Info(msg string, tags ...zap.Field) {
	Log.Info(msg, tags...)
	Log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	msg = fmt.Sprintf("%s - ERROR - %v", msg, err)
	Log.Error(msg, tags...)
	Log.Sync()
}
