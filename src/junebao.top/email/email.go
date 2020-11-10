package email

import (
	"JuneGoBlog/src/util"
	ge "gopkg.in/gomail.v2"
)

var EmailDialer *ge.Dialer

type dialerInfo struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type SimpleEmail struct {
	Recipient string // 收件人
	Sender    string // 发件人
	Subject   string // 主题
	Text      string // 正文
}

func init() {
	di := new(dialerInfo)
	util.Load("../../../secret.ini", "email", di)
	EmailDialer = ge.NewDialer(di.Host, di.Port, di.Username, di.Password)
}

func Send(subject, body string, toList []string) (err error) {
	m := ge.NewMessage()
	m.SetHeader("From", "15364968962@163.com")
	m.SetHeader("To", toList...)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	// m.Attach("/home/Alex/lolcat.jpg")
	err = EmailDialer.DialAndSend(m)
	return err
}

// 带附件发送
func SendWithFile(path, subject, body string, toList []string) (err error) {
	m := ge.NewMessage()
	m.SetHeader("From", "15364968962@163.com")
	m.SetHeader("To", toList...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	m.Attach(path)
	err = EmailDialer.DialAndSend(m)
	return err
}
