package service

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/JP-93/grpc-example/protobuf/hello/v1"
	"github.com/playground.com/grpcserver/health"
	ghc "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"time"
)

type Service struct {
	pb.UnimplementedHelloServiceServer
	healthpb.UnimplementedHealthServer
	Hc health.Health
}

var (
	sleep = flag.Duration("sleep", time.Second*5, "duration between changes in health")
)

func (s *Service) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	ht := ghc.NewServer()
	var res *healthpb.HealthCheckResponse

	go func() {
		res.Status = healthpb.HealthCheckResponse_SERVING
		for {
			ht.SetServingStatus("", res.Status)

			if res.Status == healthpb.HealthCheckResponse_SERVING {
				res.Status = healthpb.HealthCheckResponse_NOT_SERVING
			} else {
				res.Status = healthpb.HealthCheckResponse_SERVING
			}

			time.Sleep(*sleep)
		}
	}()

	_, err := s.Hc.Check(ctx, res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) Watch(request *healthpb.HealthCheckRequest, server healthpb.Health_WatchServer) error {
	return server.Send(&healthpb.HealthCheckResponse{
		Status: healthpb.HealthCheckResponse_SERVING,
	})
}

func (s *Service) CreateHello(ctx context.Context, hello *pb.Hello) (*pb.HelloResponse, error) {
	conc := fmt.Sprintf("%s %s", hello.Nome, hello.Msg)
	log.Println("Na função")
	log.Println(conc)
	return &pb.HelloResponse{Resposta: conc}, nil
}
