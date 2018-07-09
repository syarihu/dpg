package members

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/apps/members/add"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/urfave/cli"
)

func AddCommand() cli.Command {
	return cli.Command{
		Name:   "add",
		Usage:  "Invite users to the specified application",
		Action: command.AuthorizedCommandAction(newAddCommand),
		Flags:  addFlags(),
	}
}

type addCommand struct {
	endpoint    *api.AppMembersEndpoint
	requestBody *add.Request
}

func newAddCommand(c *cli.Context) (command.Command, error) {
	platform, err := apps.GetAppPlatform(c)

	if err != nil {
		return nil, err
	}

	cmd := addCommand{
		endpoint: &api.AppMembersEndpoint{
			BaseURL:      api.EndpointURL,
			AppOwnerName: apps.GetAppOwnerName(c),
			AppId:        apps.GetAppId(c),
			AppPlatform:  platform,
		},
		requestBody: &add.Request{
			UserNamesOrEmails: getInvitees(c),
			DeveloperRole:     isDeveloperRole(c),
		},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd addCommand) VerifyInput() error {
	if err := apps.RequireAppOwnerName(cmd.endpoint.AppOwnerName); err != nil {
		return err
	}

	if err := apps.RequireAppId(cmd.endpoint.AppId); err != nil {
		return err
	}

	if len(cmd.requestBody.UserNamesOrEmails) == 0 {
		return errors.New("the number of invitees must be greater than 0")
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
