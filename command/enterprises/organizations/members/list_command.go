package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/enterprises/organizations/members/list"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/enterprises"
	"github.com/jmatsu/dpg/command/enterprises/organizations"
	"gopkg.in/urfave/cli.v2"
)

func ListCommand() *cli.Command {
	return &cli.Command{
		Name:   "list",
		Usage:  "Show users who have joined to the specified enterprise's organization",
		Action: command.AuthorizedCommandAction(NewListCommand),
		Flags:  listFlags(),
	}
}

type listCommand struct {
	endpoint      *api.EnterpriseOrganizationsMembersEndpoint
	requestParams *list.Request
}

func NewListCommand(c *cli.Context) (command.Command, error) {
	cmd := listCommand{
		endpoint: &api.EnterpriseOrganizationsMembersEndpoint{
			BaseURL:          api.EndpointURL,
			EnterpriseName:   enterprises.GetEnterpriseName(c),
			OrganizationName: organizations.GetOrganizationName(c),
		},
		requestParams: &list.Request{},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

/*
Endpoint:
	enterprise name is required
	organization name is required
Parameters:
	none
*/
func (cmd listCommand) VerifyInput() error {
	if err := enterprises.RequireEnterpriseName(cmd.endpoint.EnterpriseName); err != nil {
		return err
	}

	if err := organizations.RequireOrganizationName(cmd.endpoint.OrganizationName); err != nil {
		return err
	}

	return nil
}

func (cmd listCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.GetListRequest(*authorization, *cmd.requestParams); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
