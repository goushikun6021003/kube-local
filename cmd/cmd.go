package cmd
import (
	"flag"

	"github.com/goushikun6021003/kube-local/pkg/provider/process"
	"github.com/goushikun6021003/kube-local/pkg/rpc"
)

var confPath = flag.String("conf", "./configs/config.toml", "The path of config.")

func Run() {
	// init config
	process.Init(*confPath)
	// init client
	rpc.Init()

}
