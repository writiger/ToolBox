# ToolBox

[![standard-readme compliant](https://camo.githubusercontent.com/f116695412df39ab3c98d8291befdb93af123f56aecc79fff4b20c410a5b54c7/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f726561646d652532307374796c652d7374616e646172642d627269676874677265656e2e7376673f7374796c653d666c61742d737175617265)](https://github.com/RichardLitt/standard-readme)

## 背景

这是一个用来解决个人烦恼的工具箱

* 曾经使用番茄钟尝试自律 因收费被迫停止（
* 账号密码经常记不住 全写在手机备忘录上
* 多次尝试RESTful风接口未果
* 总是突然有什么想玩的游戏过两天就忘
* 记录一些专业的词意
* 制作时间表 打卡 



## 使用说明

安装Go [Golang安装包下载地址](https://go.dev/dl/)

安装Redis  [Eedis镜像库](https://hub.docker.com/_/redis?tab=tags)

安装MySQL [MySQL下载地址](https://dev.mysql.com/downloads/mysql/)

下载依赖  `go mod download`

进入server目录修改conf.ini

启动server `go build main.go`



# 接口开发进度

| 接口路径        | 请求方式 | 功能                 |
| --------------- | -------- | -------------------- |
| /api/user/mail  | POST     | 请求发送邮箱验证码   |
| /api/user       | POST     | 注册用户             |
| /api/user/login | POST     | 登录                 |
| /api/memo       | GET      | 根据拥有者查询备忘录 |
| /api/memo       | DELETE   | 删除备忘录           |
| /api/memo       | POST     | 上传备忘录           |
| /api/memo       | PUT      | 更新备忘录           |
| /api/cipher     | POST     | 保存密码             |

