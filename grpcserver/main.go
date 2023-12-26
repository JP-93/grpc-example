package main

import (
	pb "github.com/JP-93/grpc-example/protobuf/hello/v1"
	"github.com/playground.com/grpcserver/health"
	"github.com/playground.com/grpcserver/service"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":9003")
	if err != nil {
		log.Fatalf("error creating new tcp listener: %v", err)
	}
	log.Println("running")

	newCheck := health.NewHealth()
	s := &service.Service{
		Hc: newCheck,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterHelloServiceServer(grpcServer, s)
	healthpb.RegisterHealthServer(grpcServer, s)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("error serving new tcp listener, %v", err)
	}

}
