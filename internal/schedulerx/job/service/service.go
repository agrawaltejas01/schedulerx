package job_service

import (
	"context"

	jobInterface "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/interface"
	jobModels "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/model"
	jobRepo "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/repo"
)

type Service struct {
	repo jobInterface.JobRepo
}

func NewService() *Service {
	return &Service{
		repo: jobRepo.NewJobRepo(),
	}
}

func (s *Service) CreateJobs(ctx context.Context, job []jobModels.Job) ([]jobModels.Job, error) {
	jobs, err := s.repo.CreateJobs(ctx, job)
	if err != nil {
		return []jobModels.Job{}, err
	}
	return jobs, nil
}
