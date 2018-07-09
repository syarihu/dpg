package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/teams/members/remove"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/organizations"
	"github.com/urfave/cli"
)

func RemoveCommand() cli.Command {
	return cli.Command{
		Name:   "remove",
		Usage:  "Remove users from the specified team",
		Action: command.AuthorizedCommandAction(newRemoveCommand),
		Flags:  removeFlags(),
	}
}

type removeCommand struct {
	endpoint    *api.OrganizationTeamsMembersEndpoint
	requestBody *remove.Request
}

func newRemoveCommand(c *cli.Context) (command.Command, error) {
	cmd := removeCommand{
		endpoint: &api.OrganizationTeamsMembersEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: organizations.GetOrganizationName(c),
			TeamName:         getTeamName(c),
			UserName:         getUserName(c),
		},
		requestBody: &remove.Request{},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd removeCommand) VerifyInput() error {
	if err := organizations.RequireOrganizationName(cmd.endpoint.OrganizationName); err != nil {
		return err
	}

	if err := requireTeamName(cmd.endpoint.TeamName); err != nil {
		return err
	}

	if err := requireUserName(cmd.endpoint.UserName); err != nil {
		return err
	}

	return nil
}

func (cmd removeCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.DeleteRequest(*authorization, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
