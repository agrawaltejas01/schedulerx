package command_service

import (
	"context"

	commandModel "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/model"
	commandRepo "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/repo"
	databaseUtils "github.com/agrawaltejas01/schedulerx/internal/utils/database"
)

func CreateParams(ctx context.Context, cmdId string, params []string) error {
	paramModels := make([]commandModel.Params, 0)

	for _, param := range params {
		paramModel := commandModel.Params{
			Command: cmdId,
			Param:   param,
			ID:      databaseUtils.CreateID(),
		}
		paramModels = append(paramModels, paramModel)
	}

	_, err := commandRepo.CreateParams(ctx, paramModels)
	if err != nil {
		return err
	}
	return nil

}

func CreateCommand(ctx context.Context, command commandModel.Command) (cmdModel commandModel.Command, err error) {
	command.Active = true
	command.ID = databaseUtils.CreateID()

	ctx = databaseUtils.StartTransaction(ctx)
	defer databaseUtils.EndTransaction(ctx, err)

	cmdModel, err = commandRepo.CreateCommand(ctx, command)
	if err != nil {
		return commandModel.Command{}, err
	}

	if len(command.Params) > 0 {
		err := CreateParams(ctx, command.Command, command.Params)
		if err != nil {
			return commandModel.Command{}, err
		}
	}

	return command, nil
}

func GetCommand(ctx context.Context, cmd string) (commandModel.Command, error) {
	command, err := commandRepo.GetCommand(ctx, cmd)

	if err != nil {
		return commandModel.Command{}, err
	}

	params, err := commandRepo.GetParams(ctx, cmd)
	if err != nil {
		return commandModel.Command{}, err
	}

	command.Params = make([]string, 0)
	for _, param := range params {
		command.Params = append(command.Params, param.Param)
	}

	return command, nil
}
