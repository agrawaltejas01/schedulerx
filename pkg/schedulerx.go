package schedulerx

import (
	"context"
	"fmt"
	"time"

	db "github.com/agrawaltejas01/schedulerx/internal/database"
	commandInterface "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/interface"
	commandModel "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/model"
	commandService "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/service"
	executorInterface "github.com/agrawaltejas01/schedulerx/internal/schedulerx/executor/interface"
	executorService "github.com/agrawaltejas01/schedulerx/internal/schedulerx/executor/service"
	jobInterface "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/interface"
	job_model "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/model"
	jobService "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/service"
	schedulerInterface "github.com/agrawaltejas01/schedulerx/internal/schedulerx/scheduler/interface"
	schedulerService "github.com/agrawaltejas01/schedulerx/internal/schedulerx/scheduler/service"
)

type SchedulerX struct {
	scheduler schedulerInterface.SchedulerService
	executor  executorInterface.ExecutorService
	command   commandInterface.Service
	job       jobInterface.JobService

	SchedulerFreq time.Duration
	ExecutorFreq  time.Duration
}

type DBConfig struct {
	Config db.DBConfig
}

func startScheduler(ctx context.Context, freq time.Duration,
	schedulerService schedulerInterface.SchedulerService) {

	ticker := time.NewTicker(freq)
	defer ticker.Stop()

	// First run to ensure the scheduler starts immediately
	err := schedulerService.Schedule(ctx)
	if err != nil {
		fmt.Println("Error in starting the scheduler service: " + err.Error())
	}

	for range ticker.C {
		err := schedulerService.Schedule(ctx)
		if err != nil {
			fmt.Println("Error in starting the scheduler service: " + err.Error())
		}
	}

}

func startExecutor(ctx context.Context, freq time.Duration,
	executorService executorInterface.ExecutorService) {

	ticker := time.NewTicker(freq)
	defer ticker.Stop()

	// First run to ensure the scheduler starts immediately
	err := executorService.Execute(ctx)
	if err != nil {
		fmt.Println("Error in starting the executor service: " + err.Error())
	}

	for range ticker.C {
		err := executorService.Execute(ctx)
		if err != nil {
			fmt.Println("Error in starting the executor service: " + err.Error())
		}
	}

}

// NewSchedulerX initializes a new SchedulerX instance with the provided context, scheduler frequency,
// executor frequency, and database configuration. It connects to the database, migrates the schema,
// and starts the scheduler and executor services in separate goroutines.
func NewSchedulerX(ctx context.Context, schedulerFreq, executorFreq time.Duration, dbConfig *DBConfig) *SchedulerX {
	db.Connect(&dbConfig.Config)
	db.Migrate()
	instance := &SchedulerX{
		scheduler:     schedulerService.NewService(),
		executor:      executorService.NewService(),
		command:       commandService.NewService(),
		job:           jobService.NewService(),
		SchedulerFreq: schedulerFreq,
		ExecutorFreq:  executorFreq,
	}

	go startScheduler(ctx, instance.SchedulerFreq, instance.scheduler)
	go startExecutor(ctx, instance.ExecutorFreq, instance.executor)

	return instance
}

func (s *SchedulerX) RegisterCommand(ctx context.Context, command string, cron string, params []string) error {
	s.command.CreateCommand(ctx, commandModel.Command{
		Command: command,
		Cron:    cron,
		Params:  params,
	})

	return nil
}

func (s *SchedulerX) GetExecutionOfCommands(ctx context.Context, cmd string) ([]job_model.Job, error) {
	return s.job.GetJobsByCommand(ctx, cmd)
}
