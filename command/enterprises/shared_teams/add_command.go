package shared_teams

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/enterprises/shared_teams/add"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/enterprises"
	"gopkg.in/urfave/cli.v2"
)

func AddCommand() *cli.Command {
	return &cli.Command{
		Name:   "add",
		Usage:  "Added a shared team to the specified enterprise",
		Action: command.AuthorizedCommandAction(NewAddCommand),
		Flags:  addFlags(),
	}
}

type addCommand struct {
	endpoint    *api.EnterpriseSharedTeamsEndpoint
	requestBody *add.Request
}

func NewAddCommand(c *cli.Context) (command.Command, error) {
	cmd := addCommand{
		endpoint: &api.EnterpriseSharedTeamsEndpoint{
			BaseURL:        api.EndpointURL,
			EnterpriseName: enterprises.GetEnterpriseName(c),
		},
		requestBody: &add.Request{
			SharedTeamName: GetSharedTeamName(c),
		},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

/*
Endpoint:
	enterprise name is required
Parameters:
	shared team name is required
*/
func (cmd addCommand) VerifyInput() error {
	if err := enterprises.RequireEnterpriseName(cmd.endpoint.EnterpriseName); err != nil {
		return err
	}

	if err := RequireSharedTeamName(cmd.requestBody.SharedTeamName); err != nil {
		return err
	}

	return nil
}

func (cmd addCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.MultiPartFormRequest(*authorization, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
