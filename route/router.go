package route

import (
	"BigScreen_Gin/controller"
	"BigScreen_Gin/middleware"

	"github.com/gin-gonic/gin"
)

func Load(router *gin.Engine) {
	router.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())

	// --------------------- 养殖户信息路由组
	// 创建控制器实例
	farmerController := controller.NewFarmerController()
	// 定义路由组
	farmerRoutes := router.Group("/api/v1/farmer")
	farmerRoutes.Use() // 可以在这里添加中间件
	{
		// 定义路由
		farmerRoutes.POST("", farmerController.Create)
		farmerRoutes.PUT("/:uid", farmerController.Update)
		farmerRoutes.GET("/:uid", farmerController.Show)
		farmerRoutes.GET("/list", farmerController.GetFarmerList)
		farmerRoutes.DELETE("/:uid", farmerController.Delete)
	}

	// --------------------- 城市信息路由组(预置数据只读，不可增删改)
	systemCityController := controller.NewSystemCityController()
	systemCityRoutes := router.Group("/api/v1/city")
	systemCityRoutes.Use()
	{
		systemCityRoutes.GET("", systemCityController.Show)
	}

	// --------------------- 大屏数据路由组
	bigScreenController := controller.NewBigScreenController()
	bigScreenRoutes := router.Group("/api/v1/bigscreen")
	bigScreenRoutes.Use()
	{
		bigScreenRoutes.GET("/maindata", bigScreenController.MainData)
		// 头部数据
		bigScreenRoutes.GET("/getHeaderData", bigScreenController.GetHeaderData)
		// 公母养殖模式
		bigScreenRoutes.GET("/breedingMode", bigScreenController.GetBreedingMode)
	}
}
