package routers

import (
	"mini-bank/controllers/auth"
	"mini-bank/controllers/example"
	"mini-bank/controllers/permission"
	"mini-bank/controllers/role"
	"mini-bank/controllers/user"
	"mini-bank/routers/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	route.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })

	route.GET("/example", example.GetData)
	route.POST("/login", auth.Login)
	route.POST("/register", auth.Register)

	route.Use(middleware.AuthMiddleware()).GET("/me", auth.GetProfile)

	adminRoute := route.Group("/admin")
	{
		adminRoute.Use(middleware.AuthMiddleware())

		adminRoute.GET("/permissions", permission.GetAll)
		adminRoute.GET("/permissions/:id", permission.GetByID)
		adminRoute.POST("/permissions", permission.Create)
		adminRoute.PUT("/permissions/:id", permission.Update)
		adminRoute.DELETE("/permissions/:id", permission.Delete)

		adminRoute.GET("/roles", role.GetAll)
		adminRoute.GET("/roles/:id", role.GetByID)
		adminRoute.POST("/roles/:id/add-permissions", role.AddPermissions)
		adminRoute.POST("/roles", role.Create)
		adminRoute.PUT("/roles/:id", role.Update)
		adminRoute.DELETE("/roles/:id", role.Delete)

		adminRoute.GET("/users", user.GetAll)
		adminRoute.GET("/users/:id", user.GetByID)
		adminRoute.POST("/users/:id/attach-roles", user.AttachRole)
		adminRoute.POST("/users", user.Create)
		adminRoute.PUT("/users/:id", user.Update)
		adminRoute.DELETE("/users/:id", user.Delete)
	}

}
