package shared

import (
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)

type TreeBranch interface {
	Home() string
}

type TreeBranchPlugin struct {
	Impl TreeBranch
}

func (t *TreeBranchPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &TreeBranchRpcServer{Impl: t.Impl}, nil
}

func (t *TreeBranchPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &TreeBranchRpc{client: c}, nil
}

type TreeBranchRpc struct {
	client *rpc.Client
}

func (tb *TreeBranchRpc) Home() string {
	var resp string
	err := tb.client.Call("Plugin.Home", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}
	return resp
}

type TreeBranchRpcServer struct {
	Impl TreeBranch
}

func (tb *TreeBranchRpcServer) Home(args interface{}, resp *string) error {
	*resp = tb.Impl.Home()
	return nil
}
