package common

import (
	"fmt"
	"log"
	"net/url"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("database.default")
	host := viper.GetString(fmt.Sprintf("database.%s.host", driverName))
	port := viper.GetString(fmt.Sprintf("database.%s.port", driverName))
	database := viper.GetString(fmt.Sprintf("database.%s.database", driverName))
	username := viper.GetString(fmt.Sprintf("database.%s.username", driverName))
	password := viper.GetString(fmt.Sprintf("database.%s.password", driverName))
	charset := viper.GetString(fmt.Sprintf("database.%s.charset", driverName))
	loc := viper.GetString(fmt.Sprintf("database.%s.loc", driverName))
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	// fmt.Println(args)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{ //建立连接时指定打印info级别的sql
		Logger: logger.Default.LogMode(logger.Info), //配置日志级别，打印出所有的sql
		NamingStrategy: schema.NamingStrategy{ // https://juejin.cn/post/7054586954901356558
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是 `user`
		},
	})
	if err != nil {
		panic("fail to connect database, err: " + err.Error())
	}

	DB = db
	log.Println("数据库连接成功")
	return db
}

func GetDB() *gorm.DB {
	return DB
}
