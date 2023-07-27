package main

import (
	"BigScreen_Gin/common"
	"BigScreen_Gin/route"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var serverPort = 8080
var distPath = "./view/dist/"

// dist目录下需要开放访问的静态资源的文件夹、文件
var assetsPath = []string{"assets", "images"}
var assetsFile = []string{"vite.svg"}

// 将以上两个整理成一个对象
var assets = map[string][]string{
	"path":  assetsPath,
	"files": assetsFile,
}

func main() {
	// 初始化配置
	InitConfig()
	// 连接数据库
	db := common.InitDB()

	// 新版本中没有Close方法关闭连接，有数据库连接池维护连接信息，也可以利用通用的数据库对象关掉连接
	// https://blog.csdn.net/dl962454/article/details/124109828
	defer func() {
		log.Println("关闭数据库连接")
		sqlDB, dbErr := db.DB() //获取通用数据库对象
		if dbErr != nil {
			panic(dbErr)
		}
		// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)
		dbErr = sqlDB.Close() //常规数据库接口 sql.DB关闭
		if dbErr != nil {
			panic(dbErr)
		}
	}()

	router := gin.Default()
	// 直接一个星号的话 这个dist文件夹下面禁止存在子文件夹  否则会报错||而下面的写法则规避了这种问题
	router.LoadHTMLGlob(distPath + "/*.html")
	router.GET("/", func(context *gin.Context) {
		//time.Sleep(5 * time.Second)
		//context.String(http.StatusOK, "Welcome Gin Server")
		context.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/api/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.Load(router)
	// 遍历assets对象 循环挂载静态资源路由
	for key, item := range assets {
		if key == "path" {
			// 遍历path数组
			for _, path := range item {
				router.Static("/"+path, distPath+path)
			}
		} else if key == "files" {
			// 遍历files数组
			for _, file := range item {
				router.StaticFile("/"+file, distPath+file)
			}
		}
	}
	//_ = browser.OpenURL("http://127.0.0.1:" + strconv.Itoa(serverPort))
	//err := router.Run(":" + strconv.Itoa(serverPort))
	//if err != nil {
	//	return
	//} // 监听并在 0.0.0.0:8080 上启动服务
	log.Println("运行于:", "http://127.0.0.1:"+strconv.Itoa(serverPort))
	// 优雅地关闭server: 当退出这个软件时或者命令行执行ctrl+c终止进程时，会执行以下的退出提示
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(serverPort),
		Handler: router,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("服务优雅关闭 ...")
	// 如果服务器无法正常关闭，则会在5秒钟后强制关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("服务优雅关闭遇到了状况: ", err)
	}
	log.Println("服务已优雅退出")
}

func InitConfig() {
	workDir, _ := os.Getwd()
	// 读取toml配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	// 替换默认的配置
	if viper.IsSet("server.port") {
		serverPort = viper.GetInt("server.port")
	}
	if viper.IsSet("assets") {
		assets = viper.GetStringMapStringSlice("assets")
	}
	log.Println("配置文件读取成功")
}
