package model

/**
CREATE TABLE `system_city`  (
  `id` int(11) NOT NULL,
  `city_id` int(11) NOT NULL DEFAULT 0 COMMENT '城市id',
  `level` int(11) NOT NULL DEFAULT 0 COMMENT '省市级别',
  `parent_id` int(11) NOT NULL DEFAULT 0 COMMENT '父级id',
  `area_code` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '区号',
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `merger_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '合并名称',
  `lng` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '经度',
  `lat` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '纬度',
  `is_show` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否展示',
  `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `update_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '城市表' ROW_FORMAT = Dynamic;
*/

type SystemCity struct {
	Id         uint   `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;NOT NULL" json:"id"`
	CityId     uint   `gorm:"column:city_id;type:int(11);NOT NULL" json:"city_id"`
	Level      uint   `gorm:"column:level;type:int(11);NOT NULL" json:"level"`
	ParentId   uint   `gorm:"column:parent_id;type:int(11);NOT NULL" json:"parent_id"`
	AreaCode   string `gorm:"column:area_code;type:varchar(30);NOT NULL" json:"area_code"`
	Name       string `gorm:"column:name;type:varchar(100);NOT NULL" json:"name"`
	MergerName string `gorm:"column:merger_name;type:varchar(255);NOT NULL" json:"merger_name"`
	Lng        string `gorm:"column:lng;type:varchar(50);NOT NULL" json:"lng"`
	Lat        string `gorm:"column:lat;type:varchar(50);NOT NULL" json:"lat"`
	IsShow     uint   `gorm:"column:is_show;type:tinyint(1);NOT NULL" json:"is_show"`
	CreateTime string `gorm:"column:create_time;type:timestamp(0);NOT NULL" json:"create_time"`
	UpdateTime string `gorm:"column:update_time;type:timestamp(0);NOT NULL" json:"update_time"`
}
