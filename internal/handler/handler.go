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
		minio := api.Group("/minio")
		{
			minio.POST("/upload-image", h.uploadImage)
		}

		users := api.Group("/users")
		{
			users.POST("", h.createUser)
			users.GET("", h.getListUser)
			// users.GET("/:id", h.getUserById)
		}

		roles := api.Group("/roles")
		{
			// roles.POST("", h.createRole)
			roles.GET("", h.getListRole)
			// roles.GET("/:id", h.getRoleById)
			// roles.PUT("/:id", h.updateRole)
			// roles.DELETE("/:id", h.deleteRole)
		}

		courses := api.Group("/courses")
		{
			courses.POST("", h.createCourse)
			courses.GET("", h.getListCourse)
			courses.GET("/:id", h.getCourseById)
			courses.PUT("/:id", h.updateCourse)
			courses.DELETE("/:id", h.deleteCourse)

			examTypes := courses.Group("/:id/exam-types")
			{
				examTypes.POST("", h.createCourseExamType)
				examTypes.GET("", h.getListCourseExamType)
				examTypes.GET("/:examTypeId", h.getCourseExamType)
				examTypes.PUT("/:examTypeId", h.updateCourseExamType)
				examTypes.DELETE("/:examTypeId", h.deleteCourseExamType)
			}
		}

		lids := api.Group("/lids")
		{
			lids.POST("", h.createLid)
			lids.GET("", h.getListLid)
			lids.GET("/:id", h.getLidById)
			lids.PUT("/:id", h.updateLid)
			lids.DELETE("/:id", h.deleteLid)
		}

		groups := api.Group("/groups")
		{
			groups.POST("", h.createGroup)
			groups.GET("", h.getListGroup)
			groups.GET("/:id", h.getGroupById)
			groups.PUT("/:id", h.updateGroup)
			groups.DELETE("/:id", h.deleteGroup)
		}

		students := api.Group("/students")
		{
			students.POST("", h.createStudent)
			students.GET("", h.getListStudent)
			students.GET("/:id", h.getStudentById)
			students.PUT("/:id", h.updateStudent)
			students.DELETE("/:id", h.deleteStudent)
		}
	}

	return router
}
