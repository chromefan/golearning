package tserver

import (
	"context"
	"fmt"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

type Context struct{
	ctx *context.Context
}
func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = 0
	fmt.Printf("call: %d * %d = %d\n", args.A, args.B, reply.C)
	return nil
}

func (t *Arith) Add(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = 0
	fmt.Printf("call: %d + %d = %d\n", args.A, args.B, reply.C)
	return nil
}

func (t *Arith) Say(ctx context.Context, args *string, reply *string) error {
	*reply = "hello " + *args
	return nil
}

type Greeter struct{}

func (s *Greeter) Say(ctx context.Context, name *string, reply *string) error {
	*reply = fmt.Sprintf("hello %s!", *name)
	return nil
}