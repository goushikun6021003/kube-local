package controllers

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/goushikun6021003/kube-local/pkg/model"
	"gopkg.in/gomail.v2"
)

func MailSetter(m *gomail.Message, SendToEmails []string, user string, title string, text string) *gomail.Message {
	// 收件人,...代表打散列表填充不定参数
	m.SetHeader("To", SendToEmails...)
	// 发件人
	m.SetAddressHeader("From", user, title)
	// 主题
	m.SetHeader("Subject", title)
	// 正文
	m.SetBody("text/html", text)

	return m
}

// SendEmail
func SendEmail(ruler *model.Ruleser) string {
	// 转化发送内容为json格式
	rulerJson, err := json.Marshal(ruler)
	if err != nil {
		log.Println("[email]", err.Error())
	}
	// 读取Model中的结构体
	mail := model.Config.Mail
	// 判断接口是否打开
	open := mail.Open
	if open == 0 {
		log.Println("[email]", "email未配置未开启状态,请先配置open-email为1")
		return "eamil未配置未开启状态,请先配置open-email为1"
	}

	//Emails= xxx1@qq.com,xxx2@qq.com,xxx3@qq.com
	SendToEmails := []string{}
	m := gomail.NewMessage()
	if len(mail.Emails) == 0 {
		return "收件人不能为空"
	}
	for _, Email := range strings.Split(mail.Emails, ",") {
		SendToEmails = append(SendToEmails, strings.TrimSpace(Email))
	}

	m = MailSetter(m, SendToEmails, mail.User, mail.Title, string(rulerJson))

	d := gomail.NewDialer(mail.Host, mail.Port, mail.User, mail.Password)
	// 发送
	err = d.DialAndSend(m)
	model.AlertToCounter.WithLabelValues("email", string(rulerJson), mail.Emails).Add(1)
	if err != nil {
		log.Println("[email]", err.Error())
	}
	log.Println("[email]", "email send ok to "+mail.Emails)
	return "email send ok to " + mail.Emails
}
