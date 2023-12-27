package service

import (
	"context"
	"fmt"
	pb "github.com/JP-93/grpc-example/protobuf/hello/v1"
	"log"
)

type Service struct {
	pb.UnimplementedHelloServiceServer
}

func (s *Service) CreateHello(ctx context.Context, hello *pb.Hello) (*pb.HelloResponse, error) {
	conc := fmt.Sprintf("%s %s", hello.Nome, hello.Msg)
	log.Println("Na função")
	log.Println(conc)
	return &pb.HelloResponse{Resposta: conc}, nil
}
