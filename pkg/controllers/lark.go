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
type Content struct {
	Text string `json:"text"`
}
type LarkMessage struct {
	Msg_tpye string `json:"msg_type"`
	Content Content `json:"content"`
}

func PostToLark( ruler *model.Ruleser )(string)  {
	// 转化发送内容为json格式
	rulerJson, err := json.Marshal(ruler)
	if err != nil {
		log.Println("[lark]",err.Error())
	}
	// 读取Model中的结构体
	lark := model.Config.Lark
	// 判断接口是否打开
	open := lark.Open
	if open == 0 {
		log.Println("[lark]","lark接口未配置未开启状态,请先配置open-lark为1")
		return "lark接口未配置未开启状态,请先配置open-lark为1"
	}
	// 初始化Message
	u := LarkMessage{
		"text",
		Content{string(rulerJson)},
	}
	// 序列化Message
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	log.Println("[lark]",b)

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
	res,err  := client.Post(lark.Url, "application/json", b)
	if err != nil {
		log.Println("[lark]",err.Error())
	}
	// 关闭连接通道
	defer res.Body.Close()
	// 读取返回信息
	result,err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("[lark]",err.Error())
	}
	model.AlertToCounter.WithLabelValues("lark",string(rulerJson),"").Add(1)
	log.Println("[lark]",string(result))

	return string(result)
}
