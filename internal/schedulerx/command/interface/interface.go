package command_interface

import (
	"context"

	commandModel "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/model"
)

type Repo interface {
	CreateCommand(ctx context.Context, command commandModel.Command) (commandModel.Command, error)
	CreateParams(ctx context.Context, params []commandModel.Params) ([]commandModel.Params, error)
	GetCommand(ctx context.Context, cmd string) (commandModel.Command, error)
	GetParams(ctx context.Context, cmd string) ([]commandModel.Params, error)
	GetAllActiveCommands(ctx context.Context) ([]commandModel.Command, error)
	GetParamsForMultipleCommands(ctx context.Context, cmds []string) ([]commandModel.Params, error)
}

type Service interface {
	CreateParams(ctx context.Context, cmdId string, params []string) error
	CreateCommand(ctx context.Context,
		command commandModel.Command) (cmdModel commandModel.Command, err error)
	GetCommand(ctx context.Context, cmd string) (commandModel.Command, error)
	GetAllActiveCommands(ctx context.Context) ([]commandModel.Command, error)
}
