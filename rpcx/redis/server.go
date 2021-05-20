package main

import (
	"flag"
	"log"
	"time"

	"app/golearning/rpcx/tserver"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr      = flag.String("addr", "localhost:8972", "server address")
	redisAddr = flag.String("redisAddr", "localhost:6379", "redis address")
	basePath  = flag.String("base", "passport", "user")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	s.RegisterName("Arith", new(tserver.Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}

func addRegistryPlugin(s *server.Server) {

	r := &serverplugin.RedisRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		RedisServers:   []string{*redisAddr},
		BasePath:       *basePath,
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}