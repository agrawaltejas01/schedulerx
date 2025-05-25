package executor_service

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	jobInterface "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/interface"
	jobModel "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/model"
	jobService "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/service"
)

type Service struct {
	jobService jobInterface.JobService
}

func NewService() *Service {
	return &Service{
		jobService: jobService.NewService(),
	}
}

const (
	RANGE = 5 * time.Second
)

func randomiser() bool {
	max := 10
	min := 0
	randInt := rand.Intn(max-min) + min
	return randInt%2 == 0
}

func execute(job jobModel.Job) error {

	fmt.Printf("Executing job: %s with params: %s\n", job.Command, job.Params)

	if randomiser() {
		return fmt.Errorf("simulated error executing job %s", job.ID)
	}
	return nil

}

func (s *Service) executeJob(ctx context.Context, job jobModel.Job, wg *sync.WaitGroup) {
	defer wg.Done()

	job.StartedAt = time.Now().Unix()
	err := execute(job)
	job.EndedAt = time.Now().Unix()

	if err != nil {
		job.Status = jobModel.STATUS_FAILED
		fmt.Printf("Error executing job %s: %s\n", job.ID, err.Error())
	} else {
		job.Status = jobModel.STATUS_COMPLETED
	}

	err = s.jobService.UpdateAfterExecution(ctx, job)
	if err != nil {
		fmt.Printf("Error updating job %s after execution: %s\n", job.ID, err.Error())
	}
}

func (s *Service) Execute(ctx context.Context) error {

	scheduleEndAt := time.Now()
	// subtract 5 seconds

	scheduleStartAt := scheduleEndAt.Add(-RANGE)

	jobs, err := s.jobService.GetScheduledJobsAndMarkPicked(ctx, scheduleStartAt.Unix(), scheduleEndAt.Unix())
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, job := range jobs {
		wg.Add(1)
		go s.executeJob(ctx, job, &wg)
	}

	wg.Wait()

	fmt.Println("Jobs to execute:", jobs)
	return nil
}
