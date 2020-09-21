package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/goushikun6021003/kube-local/pkg/model"
)

type SlackMessage struct {
	Text string `json:"text"`
}

func PostToSlack(ruler *model.Ruleser) string {
	// 转化发送内容为json格式
	rulerJson, err := json.Marshal(ruler)
	if err != nil {
		log.Println("[slack]", err.Error())
	}
	// 读取Model中的结构体
	slack := model.Config.Slack
	// 判断接口是否打开
	open := slack.Open
	if open == 0 {
		log.Println("[slack]", "slack接口未配置未开启状态,请先配置open-slack为1")
		msg := fmt.Sprintf("slack接口未配置未开启状态,请先配置openSlack为1")
		return msg
	}
	// 初始化Message
	u := SlackMessage{
		string(rulerJson),
	}
	// 发送Message至Webhook
	result := PostToWebhook(slack.Url, u)

	model.AlertToCounter.WithLabelValues("Slack", string(rulerJson), "").Add(1)
	log.Println("Slack " + result)

	msg := fmt.Sprintf("Slack: %s", result)
	return msg
}
