package controllers

import (
	"log"
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"

	"github.com/goushikun6021003/kube-local/pkg/model"
)

type WText struct {
	Content string `json:"content"`
}
type WechatMessage struct {
	Msgtype string `json:"msgtype"`
	Text WText `json:"text"`
}

func PostToWechat( ruler *model.Ruleser )(string)  {
	// 转化发送内容为json格式
	rulerJson, err := json.Marshal(ruler)
	if err != nil {
		log.Println("[wechat]",err.Error())
	}
	// 读取Model中的结构体
	wechat := model.Config.Wechat
	// 判断接口是否打开
	open := wechat.Open
	if open == 0 {
		log.Println("[wechat]","企业微信接口未配置未开启状态,请先配置open-wechat为1")
		return "企业微信接口未配置未开启状态,请先配置open-wechat为1"
	}
	// 初始化Message
	u := WechatMessage{
		"text",
		WText{Content:string(rulerJson)},
	}
	// 序列化Message
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	log.Println("[wechat]",b)
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
	res,err  := client.Post(wechat.Url, "application/json", b)
	if err != nil {
		log.Println("[wechat]",err.Error())
	}
	// 关闭连接通道
	defer res.Body.Close()
	// 读取返回信息
	result,err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("[wechat]",err.Error())
	}
	model.AlertToCounter.WithLabelValues("wechat",string(rulerJson),"").Add(1)
	log.Println("[wechat]",string(result))

	return string(result)
}