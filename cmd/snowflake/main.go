package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
	"os/signal"

	"github.com/gwaypg/genid/module/etc"
)

func main() {

	// 任务定时器
	go func() {
		addr := etc.Etc.String("cmd/snowflake", "rpc_listen")
		log.Infof("Rpc listen: %s", addr)
		conn, err := net.Listen("tcp", addr)
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		rpc.Accept(conn)
	}()

	// exit event
	fmt.Println("[ctrl+c to exit]")
	end := make(chan os.Signal, 2)
	signal.Notify(end, os.Interrupt, os.Kill)
	<-end
}
