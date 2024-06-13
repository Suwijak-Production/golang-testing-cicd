package router

import (
	admin "myapp/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// r.Use(middleware.AuthMiddleware()) // Add this line to use the middleware

	api := r.Group("/api") // กำหนดกลุ่มสำหรับเส้นทาง
	{
		api.GET("/ping/:id", admin.PingHandler) // กำหนดเส้นทาง /api/ping/:id ภายใต้กลุ่ม api
		api.GET("/users", admin.GetAllUser)
		api.GET("/admins", admin.GetAllAdmin)
		api.POST("/createadmin", admin.CreateAdmin)
		api.PUT("/updateadmins/:id", admin.UpdateAdmin)
	}
}
