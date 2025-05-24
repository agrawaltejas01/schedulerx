package job_interface

import (
	"context"

	jobModels "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/model"
)

type JobRepo interface {
	CreateJobs(ctx context.Context, job []jobModels.Job) ([]jobModels.Job, error)
}

type JobService interface {
	CreateJobs(ctx context.Context, job []jobModels.Job) ([]jobModels.Job, error)
}
