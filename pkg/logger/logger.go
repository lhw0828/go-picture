package logger

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	log *zap.Logger
}

// 添加日志配置选项
type Options struct {
	Level      string
	Filename   string
	MaxSize    int  // 每个日志文件的最大大小（MB）
	MaxBackups int  // 保留的旧日志文件的最大数量
	MaxAge     int  // 保留旧日志文件的最大天数
	Compress   bool // 是否压缩旧日志文件
}

func NewLoggerWithOptions(opts Options) (*Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.EncoderConfig.StacktraceKey = "stacktrace"
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// 设置日志级别
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		return nil, err
	}
	config.Level = zap.NewAtomicLevelAt(zapLevel)

	// 配置日志输出
	if opts.Filename != "" {
		config.OutputPaths = []string{opts.Filename}
		config.ErrorOutputPaths = []string{opts.Filename}
	}

	logger, err := config.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		return nil, err
	}

	return &Logger{log: logger}, nil
}

// 添加结构化字段
func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{log: l.log.With(fields...)}
}

// 添加带错误的日志
func (l *Logger) ErrorWithErr(ctx context.Context, msg string, err error, fields ...zap.Field) {
	allFields := append(fields, zap.Error(err))
	l.Error(ctx, msg, allFields...)
}

// 添加带持续时间的日志
func (l *Logger) InfoWithDuration(ctx context.Context, msg string, duration time.Duration, fields ...zap.Field) {
	allFields := append(fields, zap.Duration("duration", duration))
	l.Info(ctx, msg, allFields...)
}

// 添加带请求信息的日志
func (l *Logger) InfoWithRequest(ctx context.Context, msg string, method, path string, fields ...zap.Field) {
	allFields := append(fields,
		zap.String("method", method),
		zap.String("path", path),
	)
	l.Info(ctx, msg, allFields...)
}

// 优雅关闭
func (l *Logger) Sync() error {
	return l.log.Sync()
}

// 从上下文获取跟踪ID
func getTraceID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if id, ok := ctx.Value("trace_id").(string); ok {
		return id
	}
	return ""
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	l.log.Info(msg, append(fields, zap.String("trace_id", getTraceID(ctx)))...)
}

func (l *Logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	l.log.Error(msg, append(fields, zap.String("trace_id", getTraceID(ctx)))...)
}

func (l *Logger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	l.log.Debug(msg, append(fields, zap.String("trace_id", getTraceID(ctx)))...)
}

func (l *Logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	l.log.Warn(msg, append(fields, zap.String("trace_id", getTraceID(ctx)))...)
}