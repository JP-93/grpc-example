package main

import (
	pb "github.com/JP-93/grpc-example/protobuf/hello/v1"
	"github.com/playground.com/grpcserver/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"time"
)

var ()

func main() {
	system := pb.HelloService_ServiceDesc.ServiceName
	listener, err := net.Listen("tcp", ":9003")
	if err != nil {
		log.Fatalf("error creating new tcp listener: %v", err)
	}
	log.Println("running")

	s := &service.Service{}

	grpcServer := grpc.NewServer()
	healthcheck := health.NewServer()

	pb.RegisterHelloServiceServer(grpcServer, s)
	healthpb.RegisterHealthServer(grpcServer, healthcheck)

	go func() {
		for {
			status := healthpb.HealthCheckResponse_SERVING
			if time.Now().Second()%2 == 0 {
				status = healthpb.HealthCheckResponse_NOT_SERVING
			}
			healthcheck.SetServingStatus(system, status)
			healthcheck.SetServingStatus("", status)

			time.Sleep(1 * time.Second)
		}
	}()

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
