package logger

import (
	"context"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap/zapcore"
)

type EsLogger struct {
	Module string `json:"module"`
	Level  string `json:"level"`
	Data   string `json:"data"`
	Caller string `json:"caller"`
	Msg    string `json:"msg"`
}

type EsCore struct {
	client     *elastic.Client
	index      string
	module     string
	level      zapcore.Level
	loggerList []*EsLogger
}

func (e EsCore) Enabled(level zapcore.Level) bool {
	return level >= e.level
}

func (e EsCore) With(fields []zapcore.Field) zapcore.Core {
	return e
}

func (e EsCore) Check(entry zapcore.Entry, entry2 *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if !e.Enabled(entry.Level) {
		return entry2
	}
	return entry2.AddCore(entry, e)
}

func (e EsCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	esLogger := EsLogger{
		Module: e.module,
		Level:  entry.Level.String(),
		Data:   entry.Time.Format("2006-01-02 15:04:05"),
		Caller: entry.Caller.String(),
		Msg:    entry.Message,
	}
	e.client.Index().
		Index(e.index).
		BodyJson(&esLogger).
		Do(context.TODO())
	return nil
}

func (e EsCore) Sync() error {
	panic("implement me")
}

func newEsClient(addr []string) (*elastic.Client, error) {
	return elastic.NewClient(elastic.SetURL(addr...))
}
