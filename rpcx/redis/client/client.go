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
	redisAddr = flag.String("redisAddr", "localhost:6379", "redis address")
	basePath  = flag.String("base", "passport", "user")
)

func main() {
	flag.Parse()

	d := client.NewRedisDiscovery(*basePath, "Arith", []string{*redisAddr}, nil)
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