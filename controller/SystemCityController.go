package controller

import (
	"BigScreen_Gin/common"
	"BigScreen_Gin/model"
	"BigScreen_Gin/response"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ISystemCityController interface {
	Show(ctx *gin.Context)
}

type SystemCityController struct {
	DB *gorm.DB
}

func NewSystemCityController() ISystemCityController {
	db := common.GetDB()
	db.AutoMigrate(model.SystemCity{})
	return SystemCityController{DB: db}
}

func (sc SystemCityController) Show(ctx *gin.Context) {
	// 获取参数：如，江苏,南京,溧水
	mergerName := ctx.Query("name")
	// 逗号分割
	mergerNameArr := strings.Split(mergerName, ",")
	// 遍历mergerNameArr，去除所有空格、"省"、"市"、"县"、"区"，最后拼接回去
	var mergerNameArrNew []string
	for _, v := range mergerNameArr {
		v = strings.ReplaceAll(v, " ", "")
		v = strings.TrimSuffix(v, "省")
		v = strings.TrimSuffix(v, "市")
		v = strings.TrimSuffix(v, "县")
		v = strings.TrimSuffix(v, "区")
		mergerNameArrNew = append(mergerNameArrNew, v)
	}
	mergerName = strings.Join(mergerNameArrNew, ",")

	var systemCity model.SystemCity
	// 查询数据
	if err := sc.DB.Where("merger_name = ?", mergerName).First(&systemCity).Error; err != nil {
		response.Fail(ctx, nil, "查询失败")
		return
	}
	// 返回结果
	response.Success(ctx, gin.H{"data": systemCity}, "查询成功")
}
