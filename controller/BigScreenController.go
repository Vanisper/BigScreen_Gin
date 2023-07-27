package controller

import (
	"BigScreen_Gin/common"
	"BigScreen_Gin/model"
	"BigScreen_Gin/response"
	utils "BigScreen_Gin/util"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IBigScreenController interface {
	MainData(ctx *gin.Context)
	GetHeaderData(ctx *gin.Context)
	GetBreedingMode(ctx *gin.Context)
}

type BigScreenController struct {
	DB *gorm.DB
}

func NewBigScreenController() IBigScreenController {
	db := common.GetDB()
	return BigScreenController{DB: db}
}

type 公母养殖模式结构 struct {
	Title []string       `json:"title"`
	Datas []公母养殖模式Data结构 `json:"datas"`
}

type 公母养殖模式Data结构 struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

func (bsc BigScreenController) GetBreedingMode(ctx *gin.Context) {
	BreedingMode := 公母养殖模式结构{
		Title: []string{"养殖面积"},
		Datas: []公母养殖模式Data结构{
			{
				Name:   "♀️纯母养殖",
				Values: []string{fmt.Sprintf("%v亩", 0)},
			},
			{
				Name:   "♂️纯公养殖",
				Values: []string{fmt.Sprintf("%v亩", 0)},
			},
			{
				Name:   "♂️♀️混合养殖",
				Values: []string{fmt.Sprintf("%v亩", 0)},
			},
		},
	}

	response.Success(ctx, gin.H{"data": BreedingMode}, "获取数据成功")
}

type 头部数据结构 struct {
	Rate            头部子数据结构 `json:"率"`
	Overview        头部子数据结构 `json:"概览"`
	ProsperityIndex 头部子数据结构 `json:"景气指数"`
}

type 头部子数据结构 struct {
	Color string        `json:"color"`
	Datas []头部子数据Data结构 `json:"datas"`
}

type 头部子数据Data结构 struct {
	Name  string  `json:"name"`
	Value string  `json:"value"`
	Color *string `json:"color"`
}

func (bsc BigScreenController) GetHeaderData(ctx *gin.Context) {
	//.............................. 数据查询
	// 查询详情地址这一列
	var addressDetails []string
	if err := bsc.DB.Model(&model.Farmer{}).Select("address_detail").Find(&addressDetails).Error; err != nil {
		response.Fail(ctx, nil, fmt.Sprintf("查询address_detail失败: %v", err))
		return
	}
	// 去重
	addressDetails = utils.RemoveDuplicateElement(addressDetails)

	// 查询塘口面积这一列
	var pondAreas []float64
	if err := bsc.DB.Model(&model.Farmer{}).Select("pond_area").Find(&pondAreas).Error; err != nil {
		response.Fail(ctx, nil, fmt.Sprintf("查询pond_area失败: %v", err))
		return
	}

	// 查询塘口数量这一列
	var pondNums []int
	if err := bsc.DB.Model(&model.Farmer{}).Select("pond_num").Find(&pondNums).Error; err != nil {
		response.Fail(ctx, nil, fmt.Sprintf("查询pond_num失败: %v", err))
		return
	}

	// 查询投保面积这一列
	var insuredAreasNullFloat64 []sql.NullFloat64
	if err := bsc.DB.Model(&model.Farmer{}).Select("insured_area").Find(&insuredAreasNullFloat64).Error; err != nil {
		response.Fail(ctx, nil, fmt.Sprintf("查询insured_area失败: %v", err))
		return
	}
	// 将sql.NullFloat64类型转换为float64类型
	var insuredAreas []float64
	for _, insuredAreaNullFloat64 := range insuredAreasNullFloat64 {
		insuredAreas = append(insuredAreas, insuredAreaNullFloat64.Float64)
	}
	//.............................. 数据处理
	// 计算总户数
	totalNum := len(pondNums)
	// 计算总面积
	var totalArea float64
	for _, pondArea := range pondAreas {
		totalArea += pondArea
	}
	// 计算总投保面积
	var totalInsuredArea float64
	for _, insuredArea := range insuredAreas {
		totalInsuredArea += insuredArea
	}
	// 计算总投保率
	var totalInsuredRate float64
	if totalArea != 0 {
		totalInsuredRate = totalInsuredArea / totalArea
	}
	totalInsuredRateString := fmt.Sprintf("%.2f", totalInsuredRate*100) + "%"
	// 精品率
	totalFineRate := 0.5
	totalFineRateString := fmt.Sprintf("%.2f", totalFineRate*100) + "%"

	// 养殖面积同比增减
	pondAreaIncrease := 0.123
	pondAreaIncreaseString := fmt.Sprintf("%.2f", pondAreaIncrease*100) + "%"
	// 产能同比增减
	capacityIncrease := 0.186
	capacityIncreaseString := fmt.Sprintf("%.2f", capacityIncrease*100) + "%"
	// 头部数据
	// 头部数据第一组：率
	rate := 头部子数据结构{
		Color: "#ffff43",
		Datas: []头部子数据Data结构{
			{
				Name:  "总投保率",
				Value: totalInsuredRateString,
			},
			{
				Name:  "精品率",
				Value: totalFineRateString,
			}, {
				Name:  "精品率",
				Value: totalFineRateString,
			},
		},
	}
	// 头部数据第二组：概览
	overview := 头部子数据结构{
		Color: "#f74d52",
		Datas: []头部子数据Data结构{
			{
				Name:  "养殖户总数量",
				Value: fmt.Sprintf("%d", totalNum),
			},
			{
				Name:  "养殖总面积（亩）",
				Value: fmt.Sprintf("%.2f", totalArea),
			},
			{
				Name:  "塘口总数量",
				Value: fmt.Sprintf("%d", totalNum),
			},
		},
	}
	// 头部数据第三组：景气指数
	prosperityIndex := 头部子数据结构{
		Color: "#24f3e5",
		Datas: []头部子数据Data结构{
			{
				Name:  "养殖面积跟同比增减",
				Value: pondAreaIncreaseString,
			},
			{
				Name:  "产能跟同比增减",
				Value: capacityIncreaseString,
			},
		},
	}
	// 头部数据
	头部数据 := 头部数据结构{
		Rate:            rate,
		Overview:        overview,
		ProsperityIndex: prosperityIndex,
	}
	//.............................. 数据返回
	response.Success(ctx, gin.H{
		"data": 头部数据,
	}, "获取数据成功")
}

func (bsc BigScreenController) MainData(ctx *gin.Context) {
	//.............................. 数据查询
	// 查询详情地址这一列
	var addressDetails []string
	if err := bsc.DB.Model(&model.Farmer{}).Select("address_detail").Find(&addressDetails).Error; err != nil {
		response.Fail(ctx, nil, fmt.Sprintf("查询address_detail失败: %v", err))
		return
	}
	// 去重
	addressDetails = utils.RemoveDuplicateElement(addressDetails)

	// 查询塘口面积这一列
	var pondAreas []float64
	if err := bsc.DB.Model(&model.Farmer{}).Select("pond_area").Find(&pondAreas).Error; err != nil {
		response.Fail(ctx, nil, fmt.Sprintf("查询pond_area失败: %v", err))
		return
	}

	// 查询塘口数量这一列
	var pondNums []int
	if err := bsc.DB.Model(&model.Farmer{}).Select("pond_num").Find(&pondNums).Error; err != nil {
		response.Fail(ctx, nil, fmt.Sprintf("查询pond_num失败: %v", err))
		return
	}

	// 查询投保面积这一列
	var insuredAreasNullFloat64 []sql.NullFloat64
	if err := bsc.DB.Model(&model.Farmer{}).Select("insured_area").Find(&insuredAreasNullFloat64).Error; err != nil {
		response.Fail(ctx, nil, fmt.Sprintf("查询insured_area失败: %v", err))
		return
	}
	// 将sql.NullFloat64类型转换为float64类型
	var insuredAreas []float64
	for _, insuredAreaNullFloat64 := range insuredAreasNullFloat64 {
		insuredAreas = append(insuredAreas, insuredAreaNullFloat64.Float64)
	}
	//.............................. 数据处理
	// 计算总户数
	totalNum := len(pondNums)
	// 计算总面积
	var totalArea float64
	for _, pondArea := range pondAreas {
		totalArea += pondArea
	}
	// 计算总投保面积
	var totalInsuredArea float64
	for _, insuredArea := range insuredAreas {
		totalInsuredArea += insuredArea
	}
	// 计算总投保率
	var totalInsuredRate float64
	if totalArea != 0 {
		totalInsuredRate = totalInsuredArea / totalArea
	}
	totalInsuredRateString := fmt.Sprintf("%.2f", totalInsuredRate*100) + "%"
	// 精品率
	totalFineRate := 0.5
	totalFineRateString := fmt.Sprintf("%.2f", totalFineRate*100) + "%"

	// 养殖面积同比增减
	pondAreaIncrease := 0.123
	pondAreaIncreaseString := fmt.Sprintf("%.2f", pondAreaIncrease*100) + "%"
	// 产能同比增减
	capacityIncrease := 0.186
	capacityIncreaseString := fmt.Sprintf("%.2f", capacityIncrease*100) + "%"

	// 头部数据
	// 头部数据第一组：率
	rate := 头部子数据结构{
		Color: "#ffff43",
		Datas: []头部子数据Data结构{
			{
				Name:  "总投保率",
				Value: totalInsuredRateString,
			},
			{
				Name:  "拟精品率",
				Value: totalFineRateString,
			},
		},
	}
	// 头部数据第二组：概览
	overview := 头部子数据结构{
		Color: "#f74d52",
		Datas: []头部子数据Data结构{
			{
				Name:  "养殖户总数量",
				Value: fmt.Sprintf("%d", totalNum),
			},
			{
				Name:  "养殖总面积（亩）",
				Value: fmt.Sprintf("%.2f", totalArea),
			},
			{
				Name:  "塘口总数量",
				Value: fmt.Sprintf("%d", totalNum),
			},
		},
	}
	// 头部数据第三组：景气指数
	prosperityIndex := 头部子数据结构{
		Color: "#24f3e5",
		Datas: []头部子数据Data结构{
			{
				Name:  "养殖面积跟同比增减",
				Value: pondAreaIncreaseString,
			},
			{
				Name:  "产能跟同比增减",
				Value: capacityIncreaseString,
			},
		},
	}
	// 头部数据
	头部数据 := 头部数据结构{
		Rate:            rate,
		Overview:        overview,
		ProsperityIndex: prosperityIndex,
	}

	response.Success(ctx, gin.H{
		"addressDetails": addressDetails,
		"pondAreas":      pondAreas,
		"pondNums":       pondNums,
		"insuredAreas":   insuredAreas,
		"头部数据":           头部数据,
	}, "获取address_detail成功")
}
