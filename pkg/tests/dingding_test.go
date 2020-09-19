package test

import (

	"log"
	"testing"
	"fmt"

	"github.com/BurntSushi/toml"
	"kube-local/pkg/controllers"
	"kube-local/pkg/model"
	"kube-local/pkg/provider/process"

)

var ruler model.Ruleser

func Init(confPath string) {


	ruler = model.Ruleser{
		Id:          1,
		Expr:        "2",
		Op:          "3",
		Value:       "4",
		For:         "5",
		Summary:     "6",
		Description: "7",
		Prom:        &model.Proms{
			Id:   8,
			Name: "9",
			Url:  "10",
		},
		Plan:        &model.Plans{
			Id:          11,
			RuleLabels:  "12",
			Description: "13",
		},
	}
	// init runtime
	if _, err := toml.DecodeFile(confPath, &model.Config); err != nil {
		log.Println(err)
		return
	}

	process.InitLog()
}

func TestDingDing(t *testing.T) {
	Init("config.toml")

	fmt.Print(controllers.PostToDingDing(&ruler))

}

func TestLark(t *testing.T) {
	//title := "哆啦A梦告警"
	Init("config.toml")

	fmt.Print(controllers.PostToLark(&ruler))
}

func TestSlack(t *testing.T) {
	//title := "哆啦A梦告警"
	Init("config.toml")

	fmt.Print(controllers.PostToSlack(&ruler))
}

func TestWeChat(t *testing.T)  {
	//title := "哆啦A梦告警"
	Init("config.toml")

	fmt.Print(controllers.PostToWechat(&ruler))
}

func TestEmail(t *testing.T)  {
	//title := "哆啦A梦告警"
	Init("config.toml")


	fmt.Print(controllers.SendEmail(&ruler))
}

