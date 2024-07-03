package service

import (
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/logger"

	"google.golang.org/grpc/codes"
)

type LidReaderService struct {
	repo   repository.Repository
	logger logger.Logger
}

func NewLidReaderService(repo repository.Repository, logger logger.Logger) *LidReaderService {
	return &LidReaderService{
		repo:   repo,
		logger: logger,
	}
}

func (s *LidReaderService) GetList(pagination *model.Pagination) ([]model.Lid, error) {
	lids, err := s.repo.LidRepo.LidReader.GetList(pagination)
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	return lids, nil
}

func (s *LidReaderService) GetById(id int64) (model.Lid, error) {
	lid, err := s.repo.LidRepo.LidReader.GetById(id)
	if err != nil {
		return model.Lid{}, serviceError(err, codes.Internal)
	}

	return lid, nil
}
