package job_repo

import (
	"context"
	"fmt"

	jobModels "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/model"
	dbUtils "github.com/agrawaltejas01/schedulerx/internal/utils/database"
	"gorm.io/gorm/clause"
)

type JobRepo struct{}

func NewJobRepo() *JobRepo {
	return &JobRepo{}
}

func (j *JobRepo) CreateJobs(ctx context.Context, job []jobModels.Job) ([]jobModels.Job, error) {
	txn := dbUtils.GetTxnOrDb(ctx)
	err := txn.Clauses(clause.OnConflict{DoNothing: true}).Create(&job).Error
	// err := txn.Create(&job).Error
	if err != nil {
		fmt.Println("Error in Repo Layer for Create Job")
		return []jobModels.Job{}, err
	}
	return job, nil
}
