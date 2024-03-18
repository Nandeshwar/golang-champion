package main

import (
	"fmt"
	"myprj/pkg/router"
)

func main() {
	fmt.Println("Welcome to GRPC world")
	s := router.NewServer()
	go s.ServeGrpc()
	s.ServeRestApi()
}
