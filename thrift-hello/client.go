package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"thrift/gen-go/rpc"
)


func main() {
	addr := "localhost:8080"

	var transport thrift.TTransport
	var err error

	// 网络传输方式:需要与服务端一致
	transport, err = thrift.NewTSocket(addr)
	defer transport.Close()
	if err != nil {
		panic(err)
	}
	err = transport.Open()
	if err != nil {
		panic(err)
	}
	defer transport.Close()

	// 传输协议:需要与服务端一致
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	iProtocol := protocolFactory.GetProtocol(transport)
	oProtocol := protocolFactory.GetProtocol(transport)
	tClient := thrift.NewTStandardClient(iProtocol, oProtocol)

	// 实际业务
	ctx := context.Background()
	cli := rpc.NewHelloClient(tClient)
	rs, err := cli.HelloString(ctx, 248243)
	fmt.Println(rs,err)
}
