package main

import (
	"log"
	"net"

	"github.com/shandysiswandi/grpc-service/pb"
	"github.com/shandysiswandi/grpc-service/repository"
	"github.com/shandysiswandi/grpc-service/repository/inmem"
	"github.com/shandysiswandi/grpc-service/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	var repo repository.DBReaderWriter
	{
		repo = inmem.New()
	}

	var srv pb.TodoServiceServer
	{
		srv = service.New(repo)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTodoServiceServer(grpcServer, srv)

	log.Println("serve on port 50000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s\n", err)
	}
}
