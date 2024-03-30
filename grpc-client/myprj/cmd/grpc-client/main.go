package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	p "myprj/pkg/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:9901", grpc.WithInsecure())

	if err != nil {
		fmt.Println("error connecting grpc server. error=", err.Error())
		return
	}

	fmt.Println("Connection to grpc is successfull", conn)

	client := p.NewBookClient(conn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEyNTMyNDUsIlVzZXJuYW1lIjoibmtzIiwicm9sZSI6ImFkbWluIn0.83Tskgj2IOb2sxzS-c_1PoggIwK-w-u5uDDc6cItUhA")
	response, err := client.GetBookInfo(ctx, &p.GetBookInfoReq{Name: "nks1"})
	if err != nil {
		fmt.Println("error calling grpc api. error=", err.Error())
		return
	}
	fmt.Println(response.GetName())
	fmt.Println(response.GetAuther())
	fmt.Println(response.GetPublisher())
}
