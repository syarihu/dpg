package members

import (
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/enterprises"
	"github.com/urfave/cli"
)

type packageOption int

const (
	userName packageOption = iota
)

func (o packageOption) name() string {
	switch o {
	case userName:
		return "username"
	}

	panic("Option name mapping is not found")
}

func (o packageOption) flag() cli.Flag {
	switch o {
	case userName:
		return cli.StringSliceFlag{
			Name:  o.name(),
			Usage: "A name of a user",
		}
	}

	panic("Option name mapping is not found")
}

func getUserName(c *cli.Context) string {
	return c.String(userName.name())
}

func addFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		enterprises.EnterpriseName.Flag(),
		userName.flag(),
	}
}

func listFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		enterprises.EnterpriseName.Flag(),
	}
}

func removeFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		enterprises.EnterpriseName.Flag(),
		userName.flag(),
	}
}