package command_repo

import (
	"context"
	"fmt"

	commandModel "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/model"
	dbUtils "github.com/agrawaltejas01/schedulerx/internal/utils/database"
)

type CmdRepo struct{}

func NewCmdRepo() *CmdRepo {
	return &CmdRepo{}
}

func (c *CmdRepo) CreateCommand(ctx context.Context, command commandModel.Command) (commandModel.Command, error) {
	txn := dbUtils.GetTxnOrDb(ctx)
	err := txn.Create(&command).Error
	if err != nil {
		fmt.Println("Error in Repo Layer for Create Command")
		return commandModel.Command{}, err
	}
	return command, nil
}

func (c *CmdRepo) CreateParams(ctx context.Context, params []commandModel.Params) ([]commandModel.Params, error) {
	txn := dbUtils.GetTxnOrDb(ctx)
	err := txn.Create(&params).Error
	if err != nil {
		fmt.Println("Error in Repo Layer for Create Params")
		return nil, err
	}
	return params, nil
}

func (c *CmdRepo) GetCommand(ctx context.Context, cmd string) (commandModel.Command, error) {
	txn := dbUtils.GetTxnOrDb(ctx)
	var command commandModel.Command
	err := txn.Where("command = ?", cmd).First(&command).Error
	if err != nil {
		fmt.Println("Error in Repo Layer for Get Command")
		return commandModel.Command{}, err
	}
	return command, nil
}

func (c *CmdRepo) GetParams(ctx context.Context, cmd string) ([]commandModel.Params, error) {
	txn := dbUtils.GetTxnOrDb(ctx)
	var params []commandModel.Params
	err := txn.Model(commandModel.Params{}).Where("command = ?", cmd).Scan(&params).Error
	if err != nil {
		fmt.Println("Error in Repo Layer for Get Command")
		return []commandModel.Params{}, err
	}
	return params, nil
}

func (c *CmdRepo) GetAllActiveCommands(ctx context.Context) ([]commandModel.Command, error) {
	txn := dbUtils.GetTxnOrDb(ctx)
	var commands []commandModel.Command
	err := txn.Where("active = ?", true).Find(&commands).Error
	if err != nil {
		fmt.Println("Error in Repo Layer for Get All Active Commands")
		return nil, err
	}
	return commands, nil
}

func (c *CmdRepo) GetParamsForMultipleCommands(ctx context.Context, cmds []string) ([]commandModel.Params, error) {
	txn := dbUtils.GetTxnOrDb(ctx)
	var params []commandModel.Params
	err := txn.Model(commandModel.Params{}).Where("command in ?", cmds).Scan(&params).Error
	if err != nil {
		fmt.Println("Error in Repo Layer for Get Params For Multiple Commands")
		return nil, err
	}

	return params, nil
}
