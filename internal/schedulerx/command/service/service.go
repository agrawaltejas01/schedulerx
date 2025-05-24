package command_service

import (
	"context"

	cmdInterface "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/interface"
	commandModel "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/model"
	cmdRepo "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/repo"
	databaseUtils "github.com/agrawaltejas01/schedulerx/internal/utils/database"
)

type Service struct {
	Repo cmdInterface.Repo
}

func NewService() *Service {
	repo := cmdRepo.NewCmdRepo()
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateParams(ctx context.Context, cmdId string, params []string) error {
	paramModels := make([]commandModel.Params, 0)

	for _, param := range params {
		paramModel := commandModel.Params{
			Command: cmdId,
			Param:   param,
			ID:      databaseUtils.CreateID(),
		}
		paramModels = append(paramModels, paramModel)
	}

	_, err := s.Repo.CreateParams(ctx, paramModels)
	if err != nil {
		return err
	}
	return nil

}

func (s *Service) CreateCommand(ctx context.Context,
	command commandModel.Command) (cmdModel commandModel.Command, err error) {
	command.Active = true
	command.ID = databaseUtils.CreateID()

	ctx = databaseUtils.StartTransaction(ctx)
	defer databaseUtils.EndTransaction(ctx, err)

	cmdModel, err = s.Repo.CreateCommand(ctx, command)
	if err != nil {
		return commandModel.Command{}, err
	}

	if len(command.Params) > 0 {
		err := s.CreateParams(ctx, command.Command, command.Params)
		if err != nil {
			return commandModel.Command{}, err
		}
	}

	return command, nil
}

func (s *Service) GetCommand(ctx context.Context, cmd string) (commandModel.Command, error) {
	command, err := s.Repo.GetCommand(ctx, cmd)

	if err != nil {
		return commandModel.Command{}, err
	}

	params, err := s.Repo.GetParams(ctx, cmd)
	if err != nil {
		return commandModel.Command{}, err
	}

	command.Params = make([]string, 0)
	for _, param := range params {
		command.Params = append(command.Params, param.Param)
	}

	return command, nil
}

func (s *Service) GetAllActiveCommands(ctx context.Context) ([]commandModel.Command, error) {
	commands, err := s.Repo.GetAllActiveCommands(ctx)
	if err != nil {
		return nil, err
	}

	allCmds := make([]string, 0)
	cmdToIndex := make(map[string]int)
	for _, command := range commands {
		allCmds = append(allCmds, command.Command)
		command.Params = make([]string, 0)
		cmdToIndex[command.Command] = len(allCmds) - 1
	}

	params, err := s.Repo.GetParamsForMultipleCommands(ctx, allCmds)
	if err != nil {
		return nil, err
	}

	for _, param := range params {
		if index, exists := cmdToIndex[param.Command]; exists {
			commands[index].Params = append(commands[index].Params, param.Param)
		} else {
			return nil, err
		}
	}

	return commands, nil
}
