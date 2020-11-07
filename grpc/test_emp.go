package main

import (
	pb "gitlab.etsus.net/emq/emq-protobuf/emq/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = ":9000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewEmqClient(conn)
	requestId := int64(1232134)
	deviceIds := []string{"0251131537537611","0251131537537611"}
	message := []byte("hello")
	sendTime := time.Now().UnixNano()
	r, err := c.Push(context.Background(), &pb.PushRequest{
		RequestId: requestId,
		DeviceId: deviceIds,
		Message: message,
		SendTime: sendTime,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(r.ErrorMsg)
}