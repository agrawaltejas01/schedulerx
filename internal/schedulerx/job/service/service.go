package job_service

import (
	"context"

	jobInterface "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/interface"
	jobModels "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/model"
	jobRepo "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/repo"
	dbUtils "github.com/agrawaltejas01/schedulerx/internal/utils/database"
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

func (s *Service) GetScheduledJobsAndMarkPicked(ctx context.Context,
	scheduleStart, scheduleEnd int64) ([]jobModels.Job, error) {

	ctx = dbUtils.StartTransaction(ctx)
	defer dbUtils.EndTransaction(ctx, nil)

	jobs, err := s.repo.GetScheduledJobs(ctx, scheduleStart, scheduleEnd)
	if err != nil {
		return []jobModels.Job{}, err
	}
	if len(jobs) == 0 {
		return []jobModels.Job{}, nil
	}

	jobIds := make([]string, len(jobs))
	for _, job := range jobs {
		jobIds = append(jobIds, job.ID)
	}
	err = s.repo.UpdateStatus(ctx, jobIds, jobModels.STATUS_PICKED)
	if err != nil {
		return []jobModels.Job{}, err
	}

	return jobs, nil
}

func (s *Service) UpdateAfterExecution(ctx context.Context, job jobModels.Job) error {
	ctx = dbUtils.StartTransaction(ctx)
	defer dbUtils.EndTransaction(ctx, nil)

	err := s.repo.UpdateStatus(ctx, []string{job.ID}, job.Status)
	if err != nil {
		return err
	}

	err = s.repo.UpdateAfterExecution(ctx, job)
	if err != nil {
		return err
	}

	return nil
}
