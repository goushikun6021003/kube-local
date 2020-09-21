package rpc

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/goushikun6021003/kube-local/pkg/controllers"
	"github.com/goushikun6021003/kube-local/pkg/model"
)

type Sender int

func (s *Sender) SendMessage(recvData *model.RecvData, reply *string) error {
	// 通过desobj字段判断发送种类
	switch recvData.DstObj {
	// 发送给slack
	case "slack":
		*reply = controllers.PostToSlack(&recvData.Rules)
	// 发送mail
	case "mail":
		*reply = controllers.SendEmail(&recvData.Rules)
	// 发送给Dingding
	case "dingding":
		*reply = controllers.PostToDingDing(&recvData.Rules)
	// 发送给lark
	case "lark":
		*reply = controllers.PostToLark(&recvData.Rules)
	// 发送给wechat
	case "wechat":
		*reply = controllers.PostToWechat(&recvData.Rules)
	// 发送错误
	default:
		return errors.New("error type")
	}
	return nil
}

func Init() {
	sender := new(Sender)
	err := rpc.Register(sender)
	if err != nil {
		log.Fatal("Register error:", err)
	}

	rpc.HandleHTTP()
	l, e := net.Listen("tcp", model.Config.Port)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	//go http.Serve(l, nil)
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}

}
