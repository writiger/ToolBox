package verification

import (
	"log"
	"net/smtp"
	"server/utils"
	"strconv"

	"github.com/jordan-wright/email"
)

func SendMail(receiver string, registercode int) error {
	c := utils.GetConf()
	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱
	em.From = c.EmailAccount
	em.To = []string{receiver}
	// 设置主题
	em.Subject = "Verify your email address --ToolBoxW"
	em.HTML = []byte(`
		<soan> 验证码具有时效性 请及时验证</span>
		<span>要验证您的电子邮件地址，请使用安全代码:</span>
		<span style="background:#64a0e4;font-size:32px;color:#ffffff;">` + strconv.Itoa(registercode) + `</span>
		<h1>这是一个简单的开源练手项目</h1>
		<a style="font-size:32px;text-decoration:none;" href="https://github.com/writiger/ToolBox">项目地址</a>
		<br>
		<a style="font-size:32px;text-decoration:none;" href="https://writiger.cn">作者</a>`)

	//设置服务器相关的配置
	err := em.Send("smtp.163.com:25", smtp.PlainAuth("", c.EmailAccount, c.EmailAuthorizationCode, "smtp.163.com"))
	if err != nil {
		return err
	}
	return nil
}
