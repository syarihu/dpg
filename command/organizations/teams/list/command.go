package list

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/urfave/cli"
	"strings"
	"github.com/jmatsu/dpg/api/response"
	"encoding/json"
	"github.com/jmatsu/dpg/api/request/app/users"
)

func Command() cli.Command {
	return cli.Command{
		Name:    "list-teams",
		Aliases: []string{"i"},
		Usage:   "Show teams which belong to the specified organization",
		Action:  action,
		Flags:   allFlags(),
	}
}

func action(c *cli.Context) error {
	authority := api.Authority{
		Token: apiToken.Value(c).(string),
	}

	endpoint := api.AppMemberEndpoint{
		BaseURL:      "https://deploygate.com",
		AppOwnerName: appOwnerName.Value(c).(string),
		AppId:        appId.Value(c).(string),
		AppPlatform:  appPlatform.Value(c).(string),
	}

	_, err := listUsers(
		endpoint,
		authority,
		users.Request{
		},
		c.GlobalBoolT("verbose"),
	)

	if err != nil {
		return err
	}

	return nil
}

func listUsers(e api.AppMemberEndpoint, authority api.Authority, requestParam users.Request, verbose bool) (response.AppUsersResponse, error) {
	var r response.AppUsersResponse

	if err := verifyInput(e, authority, requestParam); err != nil {
		return r, err
	}

	if bytes, err := e.GetQueryRequest(authority, requestParam, verbose); err != nil {
		return r, err
	} else if err := json.Unmarshal(bytes, &r); err != nil {
		return r, err
	} else {
		return r, nil
	}
}

func verifyInput(e api.AppMemberEndpoint, authority api.Authority, _ users.Request) error {
	if authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if e.AppOwnerName == "" {
		return errors.New("application owner must be specified")
	}

	if e.AppId == "" {
		return errors.New("application id must be specified")
	}

	if !strings.EqualFold(e.AppPlatform, "android") && !strings.EqualFold(e.AppPlatform, "ios") {
		return errors.New("A platform must be either of `android` or `ios`")
	}

	return nil
}
