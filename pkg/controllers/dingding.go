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

type DdText struct {
	Content string `json:"content"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll bool `json:"isAtAll"`
}

type DDMessage struct {
	Msgtype string `json:"msgtype"`
	Text DdText `json:"text"`
	At At `json:"at"`
}

func PostToDingDing( ruler *model.Ruleser )(string)  {
	// 转化发送内容为json格式
	rulerJson, err := json.Marshal(ruler)
	if err != nil {
		log.Println("[dingding]",err.Error())
	}
	// 读取Model中的结构体
	Dingding := model.Config.Dingding
	// 判断接口是否打开
	open:= Dingding.Open
	if open == 0 {
		log.Println("[dingding]","钉钉接口未配置未开启状态,请先配置open-dingding为1")
		return "钉钉接口未配置未开启状态,请先配置open-dingding为1"
	}
	// 判断是否@全部人
	Isatall := Dingding.All
	Atall := true
	if Isatall == 0 {
		Atall = false
	}
	// 初始化Message
	u := DDMessage{
		"text",
		DdText{ string(rulerJson)},
		At{[]string{""} , Atall},

	}
	// 序列化Message
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	log.Println("[dingding]",b)
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
	res,err  := client.Post(Dingding.Url, "application/json", b)
	if err != nil {
		log.Println("[dingding]",err.Error())
	}
	// 关闭连接通道
	defer res.Body.Close()
	// 读取返回信息
	result,err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("[dingding]",err.Error())
	}
	model.AlertToCounter.WithLabelValues("dingding",string(rulerJson),"").Add(1)
	log.Println("[dingding]",string(result))

	return string(result)
}