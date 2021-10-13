package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/shandysiswandi/grpc-service/pb"
	"github.com/shandysiswandi/grpc-service/repository"
)

type todoService struct {
	repo repository.DBReaderWriter
}

func New(repo repository.DBReaderWriter) pb.TodoServiceServer {
	return &todoService{
		repo: repo,
	}
}

func (ts *todoService) GetAll(context.Context, *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	data, err := ts.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return &pb.GetAllResponse{
		Todos: data,
	}, nil
}

func (ts *todoService) GetByID(ctx context.Context, req *pb.GetByIDRequest) (*pb.GetByIDResponse, error) {
	data, err := ts.repo.GetByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetByIDResponse{
		Todo: data,
	}, nil
}

func (ts *todoService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	id := uuid.New().String()
	if err := ts.repo.Create(id, req.Title); err != nil {
		return nil, err
	}
	return &pb.CreateResponse{}, nil
}
