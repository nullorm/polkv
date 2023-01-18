package app

import (
	"github.com/rs/xid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type CTX struct {
	ReqID string
}

func (ctx CTX) LogDebug(msg string, fields ...zapcore.Field) {
	fields = append(fields, zap.String("reqID", ctx.ReqID))
	Log.Debug(msg, fields...)
}

func (ctx CTX) LogInfo(msg string, fields ...zapcore.Field) {
	fields = append(fields, zap.String("reqID", ctx.ReqID))
	Log.Info(msg, fields...)
}
func (ctx CTX) LogWarn(msg string, fields ...zapcore.Field) {
	fields = append(fields, zap.String("reqID", ctx.ReqID))
	Log.Warn(msg, fields...)
}
func (ctx CTX) LogError(msg string, fields ...zapcore.Field) {
	fields = append(fields, zap.String("reqID", ctx.ReqID))
	Log.Error(msg, fields...)
}
func (ctx CTX) LogFatal(msg string, fields ...zapcore.Field) {
	fields = append(fields, zap.String("reqID", ctx.ReqID))
	Log.Fatal(msg, fields...)
}

func NewCtx() CTX {
	return CTX{
		ReqID: xid.New().String(),
	}
}
