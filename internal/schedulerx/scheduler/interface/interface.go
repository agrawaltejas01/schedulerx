package scheduler_interface

import (
	"context"

	commandModel "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/model"
)

type SchedulerService interface {
	Command(ctx context.Context, activeCmds []commandModel.Command) error
	Schedule(ctx context.Context) error
}
