```shell
├── conf                    #项目配置文件目录
│   └── config.toml         #大家可以选择自己熟悉的配置文件管理工具包例如：toml、xml等等
├── controllers             #控制器目录，按模块存放控制器（或者叫控制器函数），必要的时候可以继续划分子目录。
│   └── user.go
├── models                  #模型目录，负责项目的数据存储部分，例如各个模块的Mysql表的读写模型。
│   ├── food.go
│   └── user.go
├── static                  #静态资源目录，包括Js，css，jpg等等，可以通过Gin框架配置，直接让用户访问。
│   ├── css
│   ├── images
│   └── js
├── logs                    #日志文件目录，主要保存项目运行过程中产生的日志。
└── views                   #视图模板目录，存放各个模块的视图模板，当然有些项目只有api，是不需要视图部分，可以忽略这个目录
│   └── index.html
├── main.go                 #项目入口，这里负责Gin框架的初始化，注册路由信息，关联控制器函数等。
```


> 安装 https://github.com/codegangsta/gin 这一个golang库
>
> `gin -p 6666 -a 8080 -i run main.go` 对gin项目进行热更新开发
>
> 其中的 `-p` 指定的端口是热更新插件使用的端口号
>
> `-a` 是项目运行的端口号
