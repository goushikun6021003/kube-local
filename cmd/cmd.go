package cmd
import (
	"flag"
	"kube-local/pkg/provider/process"
	"kube-local/pkg/rpc"
)

var confPath = flag.String("conf", "./configs/config.toml", "The path of config.")

func Run() {
	// init config
	process.Init(*confPath)
	// init client
	rpc.Init()

}
