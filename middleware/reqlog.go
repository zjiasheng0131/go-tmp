package middleware

import (
	"bytes"
	"io"
	"time"
	"tmp/log"

	"fmt"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/random"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RequestLogMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// The configuration is initialized once per request
		uuid, err := random.UUIdV4()
		if err != nil {
			return
		}
		trace := cryptor.Md5String(uuid)
		fmt.Println(111111)
		logger.WithValue(ctx, zap.String("trace", trace))
		fmt.Println(111112)
		logger.WithValue(ctx, zap.String("request_method", ctx.Request.Method))
		fmt.Println(111113)
		logger.WithValue(ctx, zap.Any("request_headers", ctx.Request.Header))
		fmt.Println(111114)
		logger.WithValue(ctx, zap.String("request_url", ctx.Request.URL.String()))
		fmt.Println(111115)
		if ctx.Request.Body != nil {
			bodyBytes, _ := ctx.GetRawData()
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 关键点
			logger.WithValue(ctx, zap.String("request_params", string(bodyBytes)))
		}
		fmt.Println(111116)
		logger.WithContext(ctx).Info("Request")
		fmt.Println(111117)
		ctx.Next()
	}
}
func ResponseLogMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(111118)
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		startTime := time.Now()
		ctx.Next()
		duration := time.Since(startTime).String()
		ctx.Header("X-Response-Time", duration)
		logger.WithContext(ctx).Info("Response", zap.Any("response_body", blw.body.String()), zap.Any("time", duration))
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
