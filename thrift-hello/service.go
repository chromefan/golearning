package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"thrift/gen-go/rpc"
)
type RpcHello struct{}

// thrift文件定义中只有1个返回值，但这里要2个返回值
func (r RpcHello) HelloString(ctx context.Context, id int64) (string, error) {
	fmt.Printf("message from client: %v\n", id)
	id = id%1024
	str := fmt.Sprintf("hello world_id=%d", id)
	return str, nil
}

func main() {
	addr := ":8080"
	// 网络传输方式:客户端与服务端需一致
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		panic(err)
	}
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	// 服务处理
	handler := &RpcHello{}
	processor := rpc.NewHelloProcessor(handler)

	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)

	fmt.Println("start server")
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
