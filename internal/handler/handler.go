package handler

import (
	"edusync/internal/service"

	"edusync/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.HandleMethodNotAllowed = true
	router.Use(corsMiddleware())

	//swagger settings
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler), func(c *gin.Context) {
		docs.SwaggerInfo.Host = c.Request.Host
		if c.Request.TLS != nil {
			docs.SwaggerInfo.Schemes = []string{"https"}
		}
	})

	// AUTH
	router.POST("/auth/login", h.login)
	router.POST("/auth/refresh", h.refresh)

	api := router.Group("/api", h.userIdentity)
	{
		// minio := api.Group("/minio")
		// {
		// 	minio.POST("/upload-image", h.uploadImage)
		// }

		users := api.Group("/users")
		{
			users.POST("", h.createUser)
			// 	users.GET("/:id", h.getUserById)
		}

		roles := api.Group("/roles")
		{
			// roles.POST("", h.createRole)
			roles.GET("", h.getListRole)
			// roles.GET("/:id", h.getRoleById)
			// roles.PUT("/:id", h.updateRole)
			// roles.DELETE("/:id", h.deleteRole)
		}

		// courses := api.Group("/courses")
		// {
		// 	courses.POST("", h.createCourse)
		// }
	}

	return router
}
