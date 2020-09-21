package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/goushikun6021003/kube-local/pkg/model"
)

type DdText struct {
	Content string `json:"content"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type DDMessage struct {
	MsgType string `json:"msgtype"`
	Text    DdText `json:"text"`
	At      At     `json:"at"`
}

func PostToDingDing(ruler *model.Ruleser) string {
	// 转化发送内容为json格式
	rulerJson, err := json.Marshal(ruler)
	if err != nil {
		log.Println("[dingding]", err.Error())
	}
	// 读取Model中的结构体
	Dingding := model.Config.Dingding
	// 判断接口是否打开
	open := Dingding.Open
	if open == 0 {
		log.Println("[dingding]", "钉钉接口未配置未开启状态,请先配置openDingding为1")
		msg := fmt.Sprintf("钉钉接口未配置未开启状态,请先配置openDingding为1")
		return msg
	}
	// 判断是否@全部人
	isAtAll := Dingding.All
	atAll := true
	if isAtAll == 0 {
		atAll = false
	}
	// 初始化Message
	u := DDMessage{
		"text",
		DdText{string(rulerJson)},
		At{[]string{""}, atAll},
	}

	// 发送Message至Webhook
	result := PostToWebhook(Dingding.Url, u)

	model.AlertToCounter.WithLabelValues("Dingding", string(rulerJson), "").Add(1)
	log.Println("Dingding " + result)

	msg := fmt.Sprintf("Dingding: %s", result)
	return msg
}
