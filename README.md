# ToolBox

[![standard-readme compliant](https://camo.githubusercontent.com/f116695412df39ab3c98d8291befdb93af123f56aecc79fff4b20c410a5b54c7/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f726561646d652532307374796c652d7374616e646172642d627269676874677265656e2e7376673f7374796c653d666c61742d737175617265)](https://github.com/RichardLitt/standard-readme)

[接口文档ApiFox](https://www.apifox.cn/apidoc/shared-c8e90d57-35a1-4ca2-9412-e3b341d4e1b1/api-65866014)

[项目地址](https://github.com/writiger/ToolBox)

## 运行截图

### 接口



## 背景

这是一个用来解决个人烦恼的工具箱

* 曾经使用番茄钟尝试自律 因收费被迫停止（
* 账号密码经常记不住 全写在手机备忘录上
* 多次尝试RESTful风接口未果
* 总是突然有什么想玩的游戏过两天就忘
* 记录一些专业的词意
* 制作时间表 打卡 
* 消费记录j
* （想到了再补充）



## 使用说明

安装Go [Golang安装包下载地址](https://go.dev/dl/)

安装Redis  [Eedis镜像库](https://hub.docker.com/_/redis?tab=tags)

安装MySQL [MySQL下载地址](https://dev.mysql.com/downloads/mysql/)

下载依赖  `go mod download`

进入server目录修改conf.ini

启动server `go build main.go`



# 技术选型

## 后端

* gin
* xorm
* redis
* cron

# 技术难点

## 注册时验证身份

​	在这个项目中身份验证采用邮箱验证，生成验证码后保存于redis，并**设置过期时间**。过期时间能在一定程度上保证验证码的安全性。在提交注册表单时需要携带注册邮箱的验证码。使用`github.com/jordan-wright/email`库进行发送邮件。

## 跨域访问

​	使用中间件修改报头，在每次请求是添加上 `Access-Control-Allow-Origin` 头部

~~~ go
func Cors() gin.HandlerFunc {
	c := conf.GetConf()
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", c.DomainName) // 设置允许访问所有域
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		context.Header("Access-Control-Max-Age", "172800")
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Set("content-type", "application/json")
		// 前端向后端发送的OPTIONS测试请求 服务器只需回应200
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, response.OK)
		}

		//处理请求
		context.Next()
	}
}
~~~

​	并在注册路由前使用中间件

~~~ go
// 使用中间件配置跨域
r.Use(middleware.Cors())
~~~

##  身份验证

​	使用中间件在报文中读取`Authorization`信息并进行验证。

需要注意：

* 项目中的密文统一使用`Bearer `开头，在解密时需要注意删除前7位(包含空格)

~~~ go
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
~~~



​	将翻译出来的claims中包含的**user信息**与数据库中的信息进行对比。

在需要进行身份的路由上直接调用中间件即可进行身份验证

~~~ go
// 修改余额信息
r.PUT(userApiURI+"/pay", middleware.AuthMiddleWare(), api.UserChangePay)
~~~

## 密码加密解密

​	考虑到密码的重要性，加密时使用**RSA加密算法**，在注册用户时返回私钥，由服务器的Redis保存公钥用于加密。

## 定时任务

​	当前版本仅每月收入，会员续租模块需要使用定时任务，且没有分布式考虑，因此选择了较为简单的适用于单体项目的`cron`库。

​	该库的时间表达式十分易懂，了解crontab的人应该不会陌生

| 字段名              | 是否必须 | 允许的值        | 允许的特殊字符 |
| ------------------- | -------- | --------------- | -------------- |
| 秒（Second）        | 是       | 0-59            | * / , –        |
| 分（Minutes）       | 是       | 0-59            | * / , –        |
| 时（Hour）          | 是       | 0-23            | * / , –        |
| 日（Day of month）  | 是       | 1-31            | * / , – ?      |
| 月（Month）         | 是       | 1-12 or JAN-DEC | * / , –        |
| 星期（Day of week） | 是       | 0-6 or SUN-SAT  | * / , – ?      |



# 接口开发文档

[接口文档ApiFox](https://www.apifox.cn/apidoc/shared-c8e90d57-35a1-4ca2-9412-e3b341d4e1b1/api-65866014)

