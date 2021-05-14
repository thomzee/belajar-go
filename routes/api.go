package routes

import (
	"assignment_2/configs"
	"assignment_2/controllers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ApiRoute(e *echo.Echo, db *gorm.DB) {
	// Initialize base controller
	redis := configs.GetRedis()
	redisDb := configs.GetRedisDb()
	base := controllers.Controller{Db: db, Redis: redis, RedisDb: redisDb}

	// Prefixes
	apiRoute := e.Group("/api")
	versionRoute := apiRoute.Group("/v1")

	// OrderController
	orderRoute := versionRoute.Group("/order")

	orderController := controllers.OrderController{Base: base}
	orderRoute.GET("/index", orderController.Index)
	orderRoute.GET("/show/:id", orderController.Show)
	orderRoute.POST("/create", orderController.Create)
	orderRoute.PATCH("/edit/:id", orderController.Edit)
	orderRoute.DELETE("/delete/:id", orderController.Delete)
}
