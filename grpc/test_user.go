package main

import (
	pb "gitlab.wd.com/webgroup/user-center/user-center-protobuf/passport"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	Uaddress = ":9102"
)

func main() {
	conn, err := grpc.Dial(Uaddress, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewPassportClient(conn)
	requestId := int64(1232134)

	r, err := c.Login(context.Background(), &pb.LoginRequest{
		Type: requestId,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r.Code,r.Msg)
}