package main

import (
	"context"
	"flag"
	"log"
	"time"

	"app/golearning/rpcx/tserver"
	"github.com/smallnest/rpcx/client"
)

var (
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/etcd_test", "prefix path")
)

func main() {
	flag.Parse()

	d := client.NewEtcdV3Discovery(*basePath, "Arith", []string{*etcdAddr}, nil)
	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()

	args := &tserver.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &tserver.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Printf("failed to call: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)

		time.Sleep(5 * time.Second)
	}

}