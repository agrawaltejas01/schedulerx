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

func (j *JobRepo) GetScheduledJobs(ctx context.Context, scheduleStart, scheduleEnd int64) ([]jobModels.Job, error) {
	txn := dbUtils.GetTxnOrDb(ctx)
	var jobs []jobModels.Job
	err := txn.Where("status = ? AND schedule BETWEEN ? AND ?",
		jobModels.STATUS_SCHEDULED, scheduleStart, scheduleEnd).
		Find(&jobs).
		Order("id ASC").
		Limit(3).
		Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "SKIP LOCKED",
		}).
		Error
	if err != nil {
		fmt.Println("Error in Repo Layer for Get Scheduled Jobs")
		return []jobModels.Job{}, err
	}
	return jobs, nil
}

func (j *JobRepo) UpdateStatus(ctx context.Context, jobIds []string, status string) error {
	txn := dbUtils.GetTxnOrDb(ctx)
	err := txn.Model(&jobModels.Job{}).
		Where("id IN ?", jobIds).
		Update("status", status).Error
	if err != nil {
		fmt.Println("Error in Repo Layer for Update Jobs status")
		return err
	}
	return nil
}

func (j *JobRepo) UpdateAfterExecution(ctx context.Context, job jobModels.Job) error {
	txn := dbUtils.GetTxnOrDb(ctx)
	err := txn.Model(&jobModels.Job{}).
		Where("id = ?", job.ID).
		Updates(map[string]any{
			"status":     job.Status,
			"started_at": job.StartedAt,
			"ended_at":   job.EndedAt,
		}).Error
	if err != nil {
		fmt.Println("Error in Repo Layer for Update After Execution")
		return err
	}
	return nil
}
