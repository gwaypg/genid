package main

import (
	"net/rpc"

	"github.com/gwaypg/genid/module/snowflake"
)

func init() {
	rpc.RegisterName(snowflake.RpcName, NewService())
}

type serviceImpl struct {
}

func NewService() snowflake.RpcService {
	return &serviceImpl{}
}

func (impl *serviceImpl) Gen(arg *snowflake.GenArg, ret *snowflake.GenRet) error {
	ret.ID = genIds(arg.BufferLen)
	return nil
}
