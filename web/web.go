package web

import (
	"sync"
	apiV1 "tmp/api"
	"tmp/docs"
	"tmp/log"
	"tmp/middleware"
	"tmp/pkg/server/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	isMaintenanceMode bool       // Maintenance mode flag
	mutex             sync.Mutex // Mutex to protect the flag
)

func GetVm(ctx *gin.Context) {
	vmid := ctx.Param("id")
	apiV1.HandleSuccess(ctx, vmid+": I am OK")
}

func NewHTTPServer(
	logger *log.Logger,

) *http.Server {
	//gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost("0.0.0.0"),
		http.WithServerPort(7070),
	)

	// // swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	s.Use(
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)

	s.GET("/", func(ctx *gin.Context) {
		//logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, "Thank you for using test-cloud-api which is powered by Proxmox Cluster")
	})
	s.GET("/health", func(ctx *gin.Context) {
		apiV1.HandleSuccess(ctx, "I am 32K")
	})
	s.GET("/readiness", func(ctx *gin.Context) {
		apiV1.HandleSuccess(ctx, "I am Ready")
	})

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.GET("/vm/:id", GetVm)
			// noAuthRouter.GET("/vm/:id", vmHandler.GetVm)
			// noAuthRouter.DELETE("/vm/:id", vmHandler.DestroyVm)
			// noAuthRouter.POST("/vm", vmHandler.RequestCreateVm)

			// noAuthRouter.GET("/network/available", networkHandler.GetAvailableNetworks)
			// noAuthRouter.GET("/network/:id", networkHandler.GetNetwork)
			// noAuthRouter.POST("/network", networkHandler.CreateNetwork)
			// noAuthRouter.DELETE("/network/:id", networkHandler.DeleteNetwork)
			// noAuthRouter.PUT("/network/:id", networkHandler.UpdateNetwork)

			// noAuthRouter.GET("/image", imageHandler.ListImages)
			// noAuthRouter.GET("/image/filter-options", imageHandler.GetFilterOptions)
			// noAuthRouter.GET("/image/:id", imageHandler.GetImage)
			// noAuthRouter.POST("/image", imageHandler.RegisterImage)
			// noAuthRouter.DELETE("/image/:id", imageHandler.UnRegisterImage)
			// noAuthRouter.PUT("/image/:id", imageHandler.UpdateImage)
		}

	}

	return s
}
