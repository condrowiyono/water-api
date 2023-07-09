package routers

import (
	"mini-bank/controllers/example"
	"mini-bank/controllers/permission"
	"mini-bank/controllers/role"
	"mini-bank/controllers/user"
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

	route.GET("/permissions", permission.GetAll)
	route.GET("/permissions/:id", permission.GetByID)
	route.POST("/permissions", permission.Create)
	route.PUT("/permissions/:id", permission.Update)
	route.DELETE("/permissions/:id", permission.Delete)

	route.GET("/roles", role.GetAll)
	route.GET("/roles/:id", role.GetByID)
	route.POST("/roles/:id/add-permissions", role.AddPermissions)
	route.POST("/roles", role.Create)
	route.PUT("/roles/:id", role.Update)
	route.DELETE("/roles/:id", role.Delete)

	route.GET("/users", user.GetAll)
	route.GET("/users/:id", user.GetByID)
	route.POST("/users", user.Create)
	route.PUT("/users/:id", user.Update)
	route.DELETE("/users/:id", user.Delete)
}
