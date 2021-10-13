package repository

import (
	"github.com/shandysiswandi/grpc-service/pb"
)

type DBReaderWriter interface {
	GetAll() ([]*pb.Todo, error)
	GetByID(id string) (*pb.Todo, error)
	Create(id, title string) error
}
