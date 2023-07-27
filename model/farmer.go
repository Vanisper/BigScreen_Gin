package model

import (
	"database/sql"
	"log"

	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

/*
*
CREATE TABLE `farmer`  (

	`id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
	`entity` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '养殖主体',
	`brand` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '产地、地理标志（eg：固城湖、兴化）',
	`pond_area` float(10, 2) NULL DEFAULT NULL COMMENT '塘口面积',
	`pond_num` int(10) NULL DEFAULT NULL COMMENT '塘口数量',
	`exp` float(10, 2) NULL DEFAULT NULL COMMENT '养殖经验（多少年）',
	`contact_man` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系人',
	`phone` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系电话',
	`entity_type` tinyint(4) UNSIGNED ZEROFILL NULL DEFAULT NULL COMMENT '主体类型：0个人，1家庭农场，2合作社，3公司',
	`address_province` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址-省',
	`address_city` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址-市',
	`address_city_id` int(11) NULL DEFAULT NULL COMMENT '地址-市id',
	`address_district` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址-区',
	`address_detail` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址详情',
	`uid` int(10) NULL DEFAULT NULL COMMENT '用户id',
	`pic` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图片',
	`create_time` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
	`update_time` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `uid`(`uid`) USING BTREE,
	INDEX `area`(`pond_area`) USING BTREE,
	INDEX `experience`(`exp`) USING BTREE,
	INDEX `type`(`entity_type`) USING BTREE,
	INDEX `province`(`address_province`) USING BTREE,
	INDEX `city`(`address_province`, `address_city`) USING BTREE,
	INDEX `district`(`address_province`, `address_city`, `address_district`) USING BTREE,
	INDEX `mobile`(`phone`) USING BTREE

) ENGINE = InnoDB AUTO_INCREMENT = 3799 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
*/
type Farmer struct {
	Id       int     `gorm:"type:int(11) unsigned;primary_key;not null;comment:自增主键" json:"id"`
	Entity   string  `gorm:"type:varchar(255);not null;comment:养殖主体" json:"entity"`
	Brand    string  `gorm:"type:varchar(100);comment:产地、地理标志（eg：固城湖、兴化）" json:"brand"`
	PondArea float64 `gorm:"type:float(10,2);index:area;comment:塘口面积" json:"pond_area"`
	// 投保面积
	InsuredArea     sql.NullFloat64 `gorm:"type:float(10,2);index:insured_area;default:null;comment:投保面积" json:"insured_area"`
	PondNum         uint            `gorm:"type:int(10);comment:塘口数量" json:"pond_num"`
	Exp             float64         `gorm:"type:float(10,2);index:experience;comment:养殖经验（多少年）" json:"exp"`
	ContactMan      string          `gorm:"type:varchar(64);comment:联系人" json:"contact_man"`
	Phone           string          `gorm:"type:varchar(50);index:mobile;comment:联系电话" json:"phone"`
	EntityType      uint            `gorm:"type:tinyint(4) unsigned zerofill;index:type;comment:主体类型：0个人，1家庭农场，2合作社，3公司" json:"entity_type"`
	AddressProvince string          `gorm:"type:varchar(64);index:province;index:city;index:district;comment:地址-省" json:"address_province"`
	AddressCity     string          `gorm:"type:varchar(64);index:city;index:district;comment:地址-市" json:"address_city"`
	AddressCityId   uint            `gorm:"type:int(11);comment:地址-市id" json:"address_city_id"`
	AddressDistrict string          `gorm:"type:varchar(64);index:district;comment:地址-区" json:"address_district"`
	AddressDetail   string          `gorm:"type:varchar(255);comment:地址详情" json:"address_detail"`
	Uid             uuid.UUID       `gorm:"type:char(36);uniqueIndex:uid;comment:用户id" json:"uid"`
	Pic             string          `gorm:"type:varchar(255);comment:图片" json:"pic"`
	CreatedAt       Time            `gorm:"type:timestamp default CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt       Time            `gorm:"type:timestamp default CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

// 注意：必须实现这个方法，才能正确调用工具包转换方法
func (farmer Farmer) GetStructData() interface{} {
	return farmer
}

func (farmer *Farmer) BeforeCreate(db *gorm.DB) error {
	farmer.Uid = uuid.NewV4()

	log.Println("创建用户前生成用户id:", farmer.Uid)
	return nil
}

func (farmer *Farmer) BeforeUpdate(db *gorm.DB) error {
	log.Println("更新用户数据时间：", Time(time.Now()))
	return nil
}
