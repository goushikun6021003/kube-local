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
	// 1发送给slack
	case "slack":
		*reply = controllers.PostToSlack(&recvData.Rules)
	// 2发送mail
	case "mail":
		*reply = controllers.SendEmail(&recvData.Rules)
	// 3发送给Dingding
	case "dingding":
		*reply = controllers.PostToDingDing(&recvData.Rules)
	// 4发送给lark
	case "lark":
		*reply = controllers.PostToLark(&recvData.Rules)
	// 5发送给wechat
	case "wechat":
		*reply = controllers.PostToWechat(&recvData.Rules)
	// 发送错误
	default:
		return errors.New("Error type!")
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
