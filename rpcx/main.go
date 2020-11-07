package main

import (
	"app/golearning/rpcx/tserver"
	"flag"
	"github.com/smallnest/rpcx/log"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.RegisterName("Arith", new(example.Arith), "")
	err := s.Register(new(tserver.Arith), "")
	log.Info("addr 127.0.0.1:8972")
	err = s.Serve("tcp", *addr)
	if err != nil {
		log.Fatal("err",err)
	}

}