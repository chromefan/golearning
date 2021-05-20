package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/protocol"
	"log"
	"time"
	"github.com/smallnest/rpcx/client"
	pb "gitlab.wd.com/webgroup/user-center/user-center-protobuf/passport"
)

var (
	addr = flag.String("addr", "127.0.0.1:9102", "server address")
)

func main() {
	flag.Parse()

	option := client.DefaultOption
	option.SerializeType = protocol.ProtoBuffer
	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr}, {Key: *addr}})
	xclient := client.NewXClient( "PassportServer",client.Failover, client.RoundRobin, d, option)
	defer xclient.Close()


	for {
		//a := time.Now().Unix()
		args := &pb.LoginRequest{
			Type: int64(1),
			Account: "234",
			Password: "234234",
			ThirdPartyType: 1,
		}
		reply := &pb.LoginResponse{}
		err := xclient.Call(context.Background(), "Login", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("req:%d  resp:%v", args, reply)
		time.Sleep(1e9)
	}

}