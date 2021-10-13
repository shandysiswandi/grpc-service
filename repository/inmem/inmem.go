package inmem

import (
	"errors"
	"sync"

	"github.com/shandysiswandi/grpc-service/pb"
	"github.com/shandysiswandi/grpc-service/repository"
)

type inmemory struct {
	mu   sync.RWMutex
	data []*pb.Todo
}

func New() repository.DBReaderWriter {
	return &inmemory{
		data: make([]*pb.Todo, 0),
	}
}

func (in *inmemory) GetAll() ([]*pb.Todo, error) {
	in.mu.RLock()
	defer in.mu.RUnlock()

	return in.data, nil
}

func (in *inmemory) GetByID(id string) (*pb.Todo, error) {
	in.mu.RLock()
	defer in.mu.RUnlock()

	data := pb.Todo{}
	exist := false
	for _, row := range in.data {
		if row.Id == id {
			data.Id = row.Id
			data.Title = row.Title
			exist = true
			break
		}
	}

	if !exist {
		return nil, errors.New("data not found")
	}

	return &data, nil
}

func (in *inmemory) Create(id, title string) error {
	data := pb.Todo{Id: id, Title: title}

	in.mu.RLock()
	in.data = append(in.data, &data)
	in.mu.RUnlock()

	return nil
}
