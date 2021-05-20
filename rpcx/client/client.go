package main

import (
	"app/golearning/rpcx/tserver"
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()




		a := rand.Intn(100)
		b := rand.Intn(100)
		args := &tserver.Args{
			A: a,
			B: b,
		}
		reply := &tserver.Reply{}
		ctxt, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer func() {
			cancel()
			fmt.Println("timeout",ctxt)
		}()

		err := xclient.Call(ctxt, "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(1e9)


}