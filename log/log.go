package log

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const ctxLoggerKey = "zapLogger"

func init() {

	var name, sex = "pprof.cn", 1
	fmt.Println("llllloooooggg", name, sex)
}

type Logger struct {
	*zap.Logger
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	//enc.AppendString(t.Format("2006-01-02 15:04:05"))
	enc.AppendString(t.Format("2006-01-02 15:04:05.000000000"))
}

func NewLog() *Logger {
	// log address "out.log" User-defined
	lp := "log.log"
	lv := "info"
	var level zapcore.Level
	//debug<info<warn<error<fatal<panic
	switch lv {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	hook := lumberjack.Logger{
		Filename:   lp,    // Log file path
		MaxSize:    10,    // Maximum size unit for each log file: M
		MaxBackups: 2,     // The maximum number of backups that can be saved for log files
		MaxAge:     1,     // Maximum number of days the file can be saved
		Compress:   false, // Compression or not
	}

	var encoder zapcore.Encoder
	if true {
		encoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "Logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
			EncodeTime:     timeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
		})
	}
	// } else {
	// 	encoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
	// 		TimeKey:        "ts",
	// 		LevelKey:       "level",
	// 		NameKey:        "logger",
	// 		CallerKey:      "caller",
	// 		FunctionKey:    zapcore.OmitKey,
	// 		MessageKey:     "msg",
	// 		StacktraceKey:  "stacktrace",
	// 		LineEnding:     zapcore.DefaultLineEnding,
	// 		EncodeLevel:    zapcore.LowercaseLevelEncoder,
	// 		EncodeTime:     zapcore.EpochTimeEncoder,
	// 		EncodeDuration: zapcore.SecondsDurationEncoder,
	// 		EncodeCaller:   zapcore.ShortCallerEncoder,
	// 	})
	// }
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // Print to console and file
		level,
	)

	//return &Logger{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
	//return &Logger{zap.New(core, zap.AddCaller())}
	return &Logger{zap.New(core, zap.AddStacktrace(zap.ErrorLevel))}
}

// WithValue Adds a field to the specified context
func (l *Logger) WithValue(ctx context.Context, fields ...zapcore.Field) context.Context {
	//fmt.Println(222220, ctx)
	fmt.Printf("222220  %p\n", ctx)
	if c, ok := ctx.(*gin.Context); ok {
		fmt.Printf("222221  %p\n", ctx)
		fmt.Printf("222226  %p\n", c)
		//fmt.Println(222226, c)
		ctx = c.Request.Context()
		//fmt.Println(222225, ctx)
		fmt.Printf("222225  %v\n", ctx)
		c.Request = c.Request.WithContext(context.WithValue(ctx, ctxLoggerKey, l.WithContext(ctx).With(fields...)))
		return c
	}
	//fmt.Println(222223, &ctx)
	fmt.Printf("222223  %p\n", ctx)
	return context.WithValue(ctx, ctxLoggerKey, l.WithContext(ctx).With(fields...))
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	//	fmt.Println(333331, ctx)
	fmt.Printf("333331  %p\n", ctx)
	if c, ok := ctx.(*gin.Context); ok {
		fmt.Printf("333332  %p\n", ctx)
		//	fmt.Println(333332, ctx)
		ctx = c.Request.Context()
	}
	zl := ctx.Value(ctxLoggerKey)
	ctxLogger, ok := zl.(*zap.Logger)
	if ok {
		fmt.Printf("333334  %p\n", &ctxLogger)
		fmt.Printf("3333342  %p\n", ctxLogger)
		//	fmt.Println(333334, &ctxLogger)
		a2 := &Logger{ctxLogger}
		fmt.Printf("3333345  %p\n", a2)
		return a2
	}
	fmt.Printf("333335  %p\n", &l)
	fmt.Printf("333336  %p\n", l)
	//	fmt.Println(333335, &l)
	return l
}
