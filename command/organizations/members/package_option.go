package members

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/constant"
	"github.com/jmatsu/dpg/command/organizations"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/urfave/cli.v2"
)

type packageOption int

const (
	userName packageOption = iota
	userEmail
)

func (o packageOption) name() string {
	switch o {
	case userName:
		return constant.UserName
	case userEmail:
		return constant.UserEmail
	}

	panic("Option name mapping is not found")
}

func (o packageOption) flag() cli.Flag {
	switch o {
	case userName:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "The name of the registered user",
		}
	case userEmail:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "The email address of the user. If the user is not registered when using add command, the invitation e-mail will be sent",
		}
	}

	panic("Option name mapping is not found")
}

func addFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		organizations.OrganizationName.Flag(),
		userName.flag(),
		userEmail.flag(),
	}
}

func listFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		organizations.OrganizationName.Flag(),
	}
}

func removeFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		organizations.OrganizationName.Flag(),
		userName.flag(),
		userEmail.flag(),
	}
}

func getUserName(c *cli.Context) null.String {
	return null.NewString(c.String(userName.name()), c.IsSet(userName.name()))
}

func getUserEmail(c *cli.Context) null.String {
	return null.NewString(c.String(userEmail.name()), c.IsSet(userEmail.name()))
}

func requireUserNameOrUserEmail(name, email null.String) error {
	if name.String != "" && email.String != "" {
		return errors.New(fmt.Sprintf("only one of --%s or --%s is allowed", userName.name(), userEmail.name()))
	} else if name.String == "" && email.String == "" {
		return errors.New(fmt.Sprintf("either of --%s or --%s must be specified", userName.name(), userEmail.name()))
	}

	return nil
}
