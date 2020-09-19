package controllers

import (
	"log"
	"strings"
	"encoding/json"

	"github.com/go-gomail/gomail"

	"kube-local/pkg/model"
)

// SendEmail
func SendEmail( ruler *model.Ruleser ) string {
	// 转化发送内容为json格式
	rulerJson, err := json.Marshal(ruler)
	if err != nil {
		log.Println("[email]",err.Error())
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
	// 收件人,...代表打散列表填充不定参数
	m.SetHeader("To", SendToEmails...)
	// 发件人
	m.SetAddressHeader("From", mail.User, mail.Title)
	// 主题
	m.SetHeader("Subject", mail.Title)
	// 正文
	m.SetBody("text/html", string(rulerJson))
	d := gomail.NewDialer(mail.Host, mail.Port, mail.User, mail.Password)
	// 发送
	err = d.DialAndSend(m)
	model.AlertToCounter.WithLabelValues("email", string(rulerJson), mail.Emails).Add(1)
	if err != nil {
		log.Println( "[email]", err.Error())
	}
	log.Println( "[email]", "email send ok to " + mail.Emails)
	return "email send ok to " + mail.Emails
}