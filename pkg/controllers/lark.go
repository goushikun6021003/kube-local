package controllers

import (
	"encoding/json"
	"log"

	"github.com/goushikun6021003/kube-local/pkg/model"
)

type Content struct {
	Text string `json:"text"`
}
type LarkMessage struct {
	MsgType string  `json:"msg_type"`
	Content Content `json:"content"`
}

func PostToLark(ruler *model.Ruleser) string {
	// 转化发送内容为json格式
	rulerJson, err := json.Marshal(ruler)
	if err != nil {
		log.Println("[lark]", err.Error())
	}
	// 读取Model中的结构体
	lark := model.Config.Lark
	// 判断接口是否打开
	open := lark.Open
	if open == 0 {
		log.Println("[lark]", "lark接口未配置未开启状态,请先配置open-lark为1")
		return "lark接口未配置未开启状态,请先配置open-lark为1"
	}
	// 初始化Message
	u := LarkMessage{
		"text",
		Content{string(rulerJson)},
	}
	// 发送Message至Webhook
	result := PostToWebhook(lark.Url, u)

	model.AlertToCounter.WithLabelValues("Lark", string(rulerJson), "").Add(1)
	log.Println("Lark" + result)

	return result
}
