package controllers

import (
	"log"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"kube-local/pkg/model"
)


type SlackMessage struct {
	Text string `json:"text"`
}

func PostToSlack( ruler *model.Ruleser )(string)  {
	// 转化发送内容为json格式
	rulerJson, err := json.Marshal(ruler)
	if err != nil {
		log.Println("[slack]",err.Error())
	}
	// 读取Model中的结构体
	slack := model.Config.Slack
	// 判断接口是否打开
	open := slack.Open
	if open == 0 {
		log.Println("[slack]","slack接口未配置未开启状态,请先配置open-slack为1" )
		return "slack接口未配置未开启状态,请先配置open-slack为1"
	}
	// 初始化Message
	u := SlackMessage{
		string(rulerJson),
	}
	// 序列化Message
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	log.Println("[slack]",b)
	// 发起连接请求
	var tr *http.Transport
	if proxyUrl := model.Config.Proxy;proxyUrl != ""{
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		}
		tr = &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			Proxy: proxy,
		}
	}else{
		tr = &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		}
	}
	client := &http.Client{Transport: tr}
	res,err  := client.Post(slack.Url, "application/json", b)
	if err != nil {
		log.Println("[slack]",err.Error())
	}
	// 关闭连接通道
	defer res.Body.Close()
	// 读取返回信息
	result,err:=ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("[slack]",err.Error())
	}
	model.AlertToCounter.WithLabelValues("slack",string(rulerJson),"").Add(1)
	log.Println("[slack]",string(result))

	return string(result)
}
