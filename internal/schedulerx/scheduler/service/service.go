package scheduler_service

import (
	"context"
	"fmt"
	"strings"
	"time"

	cmdInterface "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/interface"
	cmdService "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/service"
	jobInterface "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/interface"
	jobModel "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/model"
	jobService "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/service"
	dbUtils "github.com/agrawaltejas01/schedulerx/internal/utils/database"
	"github.com/gorhill/cronexpr"
)

type Service struct {
	cmdService cmdInterface.Service
	jobService jobInterface.JobService
}

func NewService() *Service {
	return &Service{
		cmdService: cmdService.NewService(),
		jobService: jobService.NewService(),
	}
}

func (s *Service) Schedule(ctx context.Context) error {
	// Read the schedule from the database
	activeCmds, err := s.cmdService.GetAllActiveCommands(ctx)
	if err != nil {
		panic("Error in getting active commands: " + err.Error())
	}

	jobs := make([]jobModel.Job, 0)

	for _, cmd := range activeCmds {
		nextSchedules := cronexpr.MustParse(cmd.Cron).NextN(time.Now(), 2)
		for _, nextSchedule := range nextSchedules {

			job := jobModel.Job{
				Command:   cmd.Command,
				ID:        dbUtils.CreateID(),
				Status:    jobModel.STATUS_SCHEDULED,
				Schedule:  nextSchedule.Unix(),
				StartedAt: 0,
				EndedAt:   0,
				Params:    strings.Join(cmd.Params, " "),
			}
			jobs = append(jobs, job)
		}
	}

	_, err = s.jobService.CreateJobs(ctx, jobs)
	if err != nil {
		fmt.Printf("Error in creating jobs: %s\n", err.Error())
	}

	return nil

}
