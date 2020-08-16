package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
	"thrift/gen-go/echo"
)

func main() {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket("localhost:8070")
	if err != nil {
		fmt.Errorf("NewTSocket failed. err: [%v]\n", err)
		return
	}

	transport, err = thrift.NewTBufferedTransportFactory(8192).GetTransport(transport)
	if err != nil {
		fmt.Errorf("NewTransport failed. err: [%v]\n", err)
		return
	}
	defer transport.Close()

	if err := transport.Open(); err != nil {
		fmt.Errorf("Transport.Open failed. err: [%v]\n", err)
		return
	}

	protocolFactory := thrift.NewTCompactProtocolFactory()
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	client := echo.NewEchoClient(thrift.NewTStandardClient(iprot, oprot))

	var res *echo.EchoRes
	var params  []string
	params = append(params,"helel")
	res, err = client.Echo(context.Background(), &echo.EchoReq{
		Msg: strings.Join(params, " "),
	})
	if err != nil {
		fmt.Errorf("client echo failed. err: [%v]", err)
		return
	}

	fmt.Printf("message from server: %v", res.GetMsg())
}
