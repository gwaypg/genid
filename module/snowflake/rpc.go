package snowflake

import (
	"github.com/gwaylib/errors"
	"github.com/gwaylib/rpctry"
	"github.com/gwaypg/genid/module/etc"
)

// Rpc exported method
type RpcService interface {
	Gen(arg *GenArg, ret *GenRet) error
}

var (
	RpcName   = etc.Etc.String("cmd/snowflake", "rpc_name")
	rpcClient = rpctry.NewClient(etc.Etc.String("cmd/snowflake", "rpc_client"))
)

type GenArg struct {
	// 需要取的ID数，需要大于0
	BufferLen int
}
type GenRet struct {
	// 返回对应的ID数量，数量与需要传入的Buffer值大小相同
	ID []int64
}

// 获取服务器生成的序列ID
func Gen(arg *GenArg) (*GenRet, error) {
	if arg.BufferLen <= 0 {
		panic("need buffer len > 0")
	}
	ret := &GenRet{}
	if err := rpcClient.TryCall(RpcName+".Gen", arg, ret); err != nil {
		return ret, errors.As(err)
	}
	return ret, nil
}
