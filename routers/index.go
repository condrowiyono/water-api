package routers

import (
	"mini-bank/controllers/auth"
	"mini-bank/controllers/climate"
	"mini-bank/controllers/example"
	"mini-bank/controllers/permission"
	"mini-bank/controllers/rainfall"
	"mini-bank/controllers/river"
	"mini-bank/controllers/role"
	"mini-bank/controllers/user"
	"mini-bank/controllers/waterlevel"
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

	route.GET("/rivers", river.GetAllNoPagination)
	route.GET("/rivers/:id", river.GetByID)
	route.GET("/rivers-count", river.GetRiverCount)
	route.GET("/rainfalls/today/:river", rainfall.GetToday)
	route.GET("/waterlevels/today/:river", waterlevel.GetToday)
	route.GET("/climates/today/:river", climate.GetToday)

	meRoute := route.Group("/me")
	{
		meRoute.Use(middleware.AuthMiddleware())

		meRoute.GET("", auth.GetProfile)
	}

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

		adminRoute.GET("/rivers", river.GetAll)
		adminRoute.GET("/rivers/:id", river.GetByID)
		adminRoute.POST("/rivers", river.Create)
		adminRoute.PUT("/rivers/:id", river.Update)
		adminRoute.DELETE("/rivers/:id", river.Delete)

		adminRoute.GET("/waterlevels", waterlevel.GetAll)
		adminRoute.GET("/waterlevels/:id", waterlevel.GetByID)
		adminRoute.PUT("/waterlevels/:id", waterlevel.Update)
		adminRoute.DELETE("/waterlevels/:id", waterlevel.Delete)
		adminRoute.GET("/waterlevels/export/:river", waterlevel.ExportByID)

		adminRoute.GET("/rainfalls", rainfall.GetAll)
		adminRoute.GET("/rainfalls/:id", rainfall.GetByID)
		adminRoute.PUT("/rainfalls/:id", rainfall.Update)
		adminRoute.DELETE("/rainfalls/:id", rainfall.Delete)
		adminRoute.GET("/rainfalls/export/:river", rainfall.ExportByID)

		adminRoute.GET("/climates", climate.GetAll)
		adminRoute.GET("/climates/:id", climate.GetByID)
		adminRoute.PUT("/climates/:id", climate.Update)
		adminRoute.DELETE("/climates/:id", climate.Delete)
		adminRoute.GET("/climates/export/:river", climate.ExportByID)
	}

	mobileRoute := route.Group("/mobile")
	{
		mobileRoute.Use(middleware.AuthMiddleware())

		mobileRoute.POST("/waterlevels", waterlevel.Create)
		mobileRoute.POST("/rainfalls", rainfall.Create)
		mobileRoute.POST("/climates", climate.Create)
	}

}
