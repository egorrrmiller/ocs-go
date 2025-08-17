package service

import "ocs-go/internal/repository"

type LineService struct {
	repository *repository.Line
}

func (s *LineService) AddLine() {
	//TODO implement me
	panic("implement me")
}

func NewLineService(repository *repository.Line) *LineService {
	return &LineService{
		repository: repository,
	}
}
