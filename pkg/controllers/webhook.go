package controllers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/goushikun6021003/kube-local/pkg/model"
)

func PostToWebhook(webhook string, text interface{}) (results string) {
	// 序列化Message
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(text)
	log.Println(b)
	// 发起连接请求
	var tr *http.Transport
	if proxyUrl := model.Config.Proxy; proxyUrl != "" {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		}
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           proxy,
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	client := &http.Client{Transport: tr}
	res, err := client.Post(webhook, "application/json", b)
	if err != nil {
		log.Println(err.Error())
	}
	// 关闭连接通道
	defer res.Body.Close()
	// 读取返回信息
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
	}

	return string(result)
}
