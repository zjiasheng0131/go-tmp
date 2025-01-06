package main

import (
	"context"
	"fmt"
	"tmp/log"
	"tmp/web"

	"go.uber.org/zap"
)

// @title						Fortinet Test Cloud API
// @version					    1.0.0
// @description				    The infra layer for Forti test Cloud
// @contact.name				Kunlun(Kevin) GUAN
// @contact.email				kguan@fortinet.com
// @host						localhost:8000
// @securityDefinitions.apiKey	Bearer
// @in							header
// @name						Authorization
func main() {

	logger := log.NewLog()

	srv := web.NewHTTPServer(logger)
	logger.Info("server start", zap.String("host", "127:0.0.1:7070"))

	logger.With(zap.String("url", fmt.Sprintf("www.test.com")))
	logger.With(zap.String("name", "jimmmyr"))
	// zap.String("name", "jimmmyr"),
	// zap.Int("age", 23),
	// zap.String("agradege", "no111-000222"),
	logger.Info("test info", zap.String("url", fmt.Sprintf("www.test.com")))

	err := srv.Start(context.Background())
	if err != nil {
		fmt.Printf("Server start err: %v", err)
	}

}
