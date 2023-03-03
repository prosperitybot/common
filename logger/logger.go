package logger

import (
	"context"

	golog "log"

	"github.com/prosperitybot/common/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	if err := Init(); err != nil {
		golog.Fatal(err)
	}
}

func Init() error {
	var (
		err    error
		config = zap.NewProductionConfig()
	)

	config.Development = true
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderCfg

	log, err = config.Build(zap.AddCallerSkip(1))

	return err
}

func GetLogger() *zap.Logger {
	return log
}

type Field = zap.Field

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	log.Debug(msg, addContextInfo(ctx, fields)...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	log.Info(msg, addContextInfo(ctx, fields)...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	log.Warn(msg, addContextInfo(ctx, fields)...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	log.Error(msg, addContextInfo(ctx, fields)...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	log.Fatal(msg, addContextInfo(ctx, fields)...)
}

// Default returns pointer to the default log.
func Default() *zap.Logger {
	return log
}

func addContextInfo(ctx context.Context, fields []zapcore.Field) []zapcore.Field {
	if ctx.Value(utils.GuildIdContextKey) != nil {
		fields = append(fields, zap.String("guild_id", ctx.Value(utils.GuildIdContextKey).(string)))
	}

	if ctx.Value(utils.UserIdContextKey) != nil {
		fields = append(fields, zap.String("user_id", ctx.Value(utils.UserIdContextKey).(string)))
	}

	if ctx.Value(utils.ChannelIdContextKey) != nil {
		fields = append(fields, zap.String("channel_id", ctx.Value(utils.ChannelIdContextKey).(string)))
	}

	if ctx.Value(utils.BotIdContextKey) != nil {
		fields = append(fields, zap.String("bot_id", ctx.Value(utils.BotIdContextKey).(string)))
	}

	return fields
}
