package controller

import (
	"BigScreen_Gin/common"
	"BigScreen_Gin/model"
	"BigScreen_Gin/response"
	utils "BigScreen_Gin/util"
	"BigScreen_Gin/validator"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IFarmerController interface {
	RestController
	GetFarmerList(ctx *gin.Context)
}

type FarmerController struct {
	DB *gorm.DB
}

func NewFarmerController() IFarmerController {
	db := common.GetDB()
	db.AutoMigrate(model.Farmer{})
	return FarmerController{DB: db}
}

func (fc FarmerController) Create(ctx *gin.Context) {
	var requestFarmer validator.CreateFarmerRequest
	// 数据验证
	if err := ctx.ShouldBind(&requestFarmer); err != nil {
		response.Fail(ctx, nil, fmt.Sprintf("数据验证错误: %v", err))
		return
	}

	// 先将省市区中的空格、“省”、“市”、“区”去掉
	requestFarmer.AddressProvince = strings.TrimSuffix(strings.ReplaceAll(requestFarmer.AddressProvince, " ", ""), "省")
	requestFarmer.AddressCity = strings.TrimSuffix(strings.ReplaceAll(requestFarmer.AddressCity, " ", ""), "市")
	requestFarmer.AddressDistrict = strings.TrimSuffix(strings.ReplaceAll(requestFarmer.AddressDistrict, " ", ""), "区")
	// 拼接mergerName
	mergerName := requestFarmer.AddressProvince + "," + requestFarmer.AddressCity + "," + requestFarmer.AddressDistrict
	// 查询城市信息：调用自己的接口
	var systemCity model.SystemCity
	if err := fc.DB.Where("merger_name = ?", mergerName).First(&systemCity).Error; err != nil {
		response.Fail(ctx, nil, fmt.Sprintf("查询城市信息失败(检查创建用户的省市区正确与否): %v", err))
		return
	}
	log.Println("用户城市ID:", systemCity.CityId)

	// 创建农户
	farmer := model.Farmer{
		Entity:          requestFarmer.Entity,
		Brand:           requestFarmer.Brand,
		PondArea:        requestFarmer.PondArea,
		PondNum:         requestFarmer.PondNum,
		Exp:             requestFarmer.Exp,
		ContactMan:      requestFarmer.ContactMan,
		Phone:           requestFarmer.Phone,
		EntityType:      requestFarmer.EntityType,
		AddressProvince: requestFarmer.AddressProvince,
		AddressCity:     requestFarmer.AddressCity,
		AddressDistrict: requestFarmer.AddressDistrict,
		AddressDetail:   requestFarmer.AddressDetail,
		Pic:             requestFarmer.Pic,
		AddressCityId:   systemCity.CityId,
	}

	// 插入数据
	if err := fc.DB.Create(&farmer).Error; err != nil {
		response.Fail(ctx, nil, fmt.Sprintf("创建失败: %v", err))
		return
	}
	// 返回结果
	response.Success(ctx, nil, "创建成功")
}

func (fc FarmerController) Update(ctx *gin.Context) {
	// 获取ctx的body
	body, _ := io.ReadAll(ctx.Request.Body)
	var bodyObj map[string]interface{}

	if err := json.Unmarshal(body, &bodyObj); err != nil {
		response.Fail(ctx, nil, fmt.Sprintf("解析 JSON 失败(请确认body为json格式): %v", err))
		return
	}
	// fmt.Printf("req.body=%s\n, content-type=%v\n", body, ctx.ContentType())
	// 这点很重要，把字节流重新放回 body 中
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	// 获取路径中的参数
	farmerUid := ctx.Param("uid")
	log.Println("farmerUid:", farmerUid)
	// 查询农户是否存在
	var farmer model.Farmer
	if err := fc.DB.Where("uid = ?", farmerUid).First(&farmer).Error; err != nil {
		response.Fail(ctx, nil, "农户(uid)不存在")
		return
	}
	// 拷贝一份farmer，用于更新之后的检查
	originFarmer := farmer

	// 将bodyObj中的非空数据更新到farmer中
	vfarmer := reflect.ValueOf(&farmer).Elem()
	for i := 0; i < vfarmer.NumField(); i++ {
		field := vfarmer.Type().Field(i)
		jsonTag := field.Tag.Get("json")
		// 如果对应的jsonTag在body中存在，就更新到farmer中
		if tagValue := bodyObj[jsonTag]; tagValue != nil {
			// 获取当前字段的类型
			fieldType := vfarmer.FieldByName(field.Name).Type().String()
			// 判断是否是sql.Float64类型
			if fieldType == "sql.NullFloat64" {
				vfarmer.FieldByName(field.Name).Set(reflect.ValueOf(
					sql.NullFloat64{
						Float64: tagValue.(float64),
						Valid:   true,
					},
				))
			} else {
				vfarmer.FieldByName(field.Name).Set(reflect.ValueOf(tagValue))
			}
		}
	}

	// 对比数据是否有变化
	if reflect.DeepEqual(originFarmer, farmer) {
		response.Fail(ctx, nil, "数据无变化(请确认是否有数据更新，或者待更新字段是否正确)")
		return
	}

	// 更新城市id
	// 先将省市区中的空格、“省”、“市”、“区”去掉
	farmer.AddressProvince = strings.TrimSuffix(strings.ReplaceAll(farmer.AddressProvince, " ", ""), "省")
	farmer.AddressCity = strings.TrimSuffix(strings.ReplaceAll(farmer.AddressCity, " ", ""), "市")
	farmer.AddressDistrict = strings.TrimSuffix(strings.ReplaceAll(farmer.AddressDistrict, " ", ""), "区")
	// 拼接mergerName
	mergerName := farmer.AddressProvince + "," + farmer.AddressCity + "," + farmer.AddressDistrict
	// 查询城市信息：调用自己的接口
	var systemCity model.SystemCity
	if err := fc.DB.Where("merger_name = ?", mergerName).First(&systemCity).Error; err != nil {
		response.Fail(ctx, nil, fmt.Sprintf("查询城市信息失败(检查创建用户的省市区正确与否): %v", err))
		return
	}
	farmer.AddressCityId = systemCity.CityId
	farmer.UpdatedAt = model.Time(time.Now())
	// 更新农户信息(stuct转map确保0值也能更新)
	if err := fc.DB.Model(&farmer).Updates(utils.StructToMap(farmer)).Error; err != nil {
		response.Fail(ctx, nil, fmt.Sprintf("更新失败: %v", err))
		return
	}

	response.Success(ctx, gin.H{"data": farmer}, "更新成功")
}

func (fc FarmerController) Show(ctx *gin.Context) {
	// 获取路径中的参数——UID
	farmerUid := ctx.Param("uid")

	// 查询数据
	var farmer model.Farmer
	if err := fc.DB.Where("uid = ?", farmerUid).First(&farmer).Error; err != nil {
		response.Fail(ctx, nil, "农户(uid)不存在")
		return
	}
	// 返回结果
	response.Success(ctx, gin.H{"farmer": farmer}, "查询成功")
}

func (fc FarmerController) Delete(ctx *gin.Context) {
	// 获取路径中的参数——UID
	farmerUid := ctx.Param("uid")

	// 查询数据
	var farmer model.Farmer
	if err := fc.DB.Where("uid = ?", farmerUid).First(&farmer).Error; err != nil {
		response.Fail(ctx, nil, "农户(uid)不存在")
		return
	}
	// 删除数据
	fc.DB.Delete(&farmer)
	// 返回结果
	response.Success(ctx, nil, "删除成功")
}

func (fc FarmerController) GetFarmerList(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))
	orderField := ctx.DefaultQuery("orderField", "id") // 排序字段
	orderType := ctx.DefaultQuery("orderType", "asc")  // 排序方式: asc | desc
	if orderType != "asc" && orderType != "desc" {
		orderType = "asc"
	}

	// 分页
	var posts []model.Farmer
	fc.DB.Order(fmt.Sprintf("%s %s", orderField, orderType)).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 记录的总条数
	var total int64
	fc.DB.Model(model.Farmer{}).Count(&total)

	// 返回数据
	response.Success(ctx, gin.H{"data": posts, "total": total}, "成功")
}
