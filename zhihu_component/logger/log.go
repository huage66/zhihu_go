package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type Config struct {
	Level    string   `json:"level"`
	Path     string   `json:"path"`
	MaxAge   string   `json:"max_age"` // 最大保存多少天, 天数为单位
	EsConfig EsConfig `json:"es_config"`
}

type EsConfig struct {
	Addr  []string `json:"addr"`
	Index string   `json:"index"`
}

var (
	logger   *zap.Logger
	levelMap = map[string]zapcore.Level{
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
		"panic": zapcore.PanicLevel,
	}
)

func getLevel(level string) zapcore.Level {
	if item, ok := levelMap[level]; ok {
		return item
	}
	return zapcore.InfoLevel
}

func Use(config Config) {
	var syncWriters []zapcore.WriteSyncer
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05.000000"))
	}
	syncWriters = append(syncWriters, zapcore.AddSync(os.Stdout))
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		zapcore.NewMultiWriteSyncer(syncWriters...),
		zap.NewAtomicLevelAt(getLevel(config.Level)))
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func Info(args ...interface{}) {
	logger.Sugar().Info(args)
}

func InfoF(template string, args ...interface{}) {
	logger.Sugar().Infof(template, args...)
}

func Error(args ...interface{}) {
	logger.Sugar().Error(args)
}

func ErrorF(template string, args ...interface{}) {
	logger.Sugar().Errorf(template, args...)
}

func Panic(args ...interface{}) {
	logger.Sugar().Panic(args)
}

func PanicF(template string, args ...interface{}) {
	logger.Sugar().Panicf(template, args)
}

func Warn(args ...interface{}) {
	logger.Sugar().Warn(args)
}

func WarnF(template string, args ...interface{}) {
	logger.Sugar().Warnf(template, args)
}
