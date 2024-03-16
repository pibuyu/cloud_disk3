package helper

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"time"
)

func SendVerifyCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "<3531095171@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "测试Golang邮件发送"
	e.HTML = []byte("<b>" + "您的验证码是：" + code + "</b>")
	//返回EOF的时候，关闭SSL重试
	return e.SendWithTLS(
		"smtp.qq.com:465",
		smtp.PlainAuth("",
			"3531095171@qq.com",
			"eyhbritymwqgcjca",
			"smtp.qq.com"),
		&tls.Config{
			ServerName:         "smtp.qq.com",
			InsecureSkipVerify: true,
		},
	)
}

func GenerateCode() string {
	////生成6位随机验证码
	//numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//r := len(numeric)
	//rand.Seed(time.Now().UnixNano())
	//
	//var sb strings.Builder
	//for i := 0; i < 6; i++ {
	//	fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	//}
	//return sb.String()
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
