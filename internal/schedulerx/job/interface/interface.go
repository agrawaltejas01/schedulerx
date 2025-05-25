package job_interface

import (
	"context"

	jobModels "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/model"
)

type JobRepo interface {
	CreateJobs(ctx context.Context, job []jobModels.Job) ([]jobModels.Job, error)
	GetScheduledJobs(ctx context.Context, scheduleStart, scheduleEnd int64) ([]jobModels.Job, error)
	UpdateStatus(ctx context.Context, jobIds []string, status string) error
	UpdateAfterExecution(ctx context.Context, job jobModels.Job) error
}

type JobService interface {
	CreateJobs(ctx context.Context, job []jobModels.Job) ([]jobModels.Job, error)
	GetScheduledJobsAndMarkPicked(ctx context.Context, scheduleStart, scheduleEnd int64) ([]jobModels.Job, error)
	UpdateAfterExecution(ctx context.Context, job jobModels.Job) error
}
