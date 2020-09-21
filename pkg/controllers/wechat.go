package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/goushikun6021003/kube-local/pkg/model"
)

type WText struct {
	Content string `json:"content"`
}
type WechatMessage struct {
	Msgtype string `json:"msgtype"`
	Text    WText  `json:"text"`
}

func PostToWechat(ruler *model.Ruleser) string {
	// 转化发送内容为json格式
	rulerJson, err := json.Marshal(ruler)
	if err != nil {
		log.Println("[wechat]", err.Error())
	}
	// 读取Model中的结构体
	wechat := model.Config.Wechat
	// 判断接口是否打开
	open := wechat.Open
	if open == 0 {
		log.Println("[wechat]", "企业微信接口未配置未开启状态,请先配置open-wechat为1")
		msg := fmt.Sprintf("企业微信接口未配置未开启状态,请先配置openWechat为1")
		return msg
	}
	// 初始化Message
	u := WechatMessage{
		"text",
		WText{Content: string(rulerJson)},
	}
	// 发送Message至Webhook
	result := PostToWebhook(wechat.Url, u)

	model.AlertToCounter.WithLabelValues("Wechat", string(rulerJson), "").Add(1)
	log.Println("Wechat " + result)

	msg := fmt.Sprintf("Wechat: %s", result)
	return msg
}
