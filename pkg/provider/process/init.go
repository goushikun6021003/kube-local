package process

import (
	"log"

	"github.com/BurntSushi/toml"

	"github.com/goushikun6021003/kube-local/pkg/model"
)

func Init(confPath string) {
	// init runtime
	if _, err := toml.DecodeFile(confPath, &model.Config); err != nil {
		log.Println(err)
		return
	}
	// init log
	InitLog()
}
