package main

import (
	"flag"
	"fmt"
	pb "github.com/JP-93/grpc-example/protobuf/hello/v1"
	"github.com/playground.com/grpcserver/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"time"
)

var (
	port  = flag.Int("port", 9003, "the port to serve on")
	sleep = flag.Duration("sleep", time.Second*5, "duration between changes in health")

	system = pb.HelloService_ServiceDesc.ServiceName // empty string represents the health of the system
)

func main() {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
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
		// asynchronously inspect dependencies and toggle serving status as needed
		next := healthpb.HealthCheckResponse_SERVING

		for {
			healthcheck.SetServingStatus(system, next)

			if next != healthpb.HealthCheckResponse_SERVING {
				next = healthpb.HealthCheckResponse_NOT_SERVING
			} else {
				next = healthpb.HealthCheckResponse_SERVING
			}

			time.Sleep(*sleep)
		}
	}()

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
