package subcommand_go_init

import (
	"github.com/bar-counter/slog"
	"github.com/convention-change/zymosis/command"
	"github.com/convention-change/zymosis/command/git_command"
	"github.com/convention-change/zymosis/internal/urfave_cli"
	"github.com/urfave/cli/v2"
)

const commandName = "go"

var commandEntry *GolangCommand

type GolangCommand struct {
	isDebug bool

	isInit     bool
	gitCommand git_command.GitCommand
}

func (n *GolangCommand) Exec() error {
	slog.Debugf("-> Exec subCommand [ %s ]", commandName)

	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "go-package-path",
			Usage: "Set the kit go package path, defaults will use project root path",
		},
	}
}

func withEntry(c *cli.Context) (*GolangCommand, error) {
	slog.Debugf("-> withEntry subCommand [ %s ]", commandName)

	if c.Bool("lib") {
		slog.Info("new lib mode")
	}

	gitCommand, err := git_command.BindGitFlag(c)
	if err != nil {
		return nil, err
	}

	globalEntry := command.CmdGlobalEntry()

	return &GolangCommand{
		isDebug: globalEntry.Verbose,

		gitCommand: gitCommand,

		isInit: globalEntry.RootCfg.IsInit,
	}, nil
}

func action(c *cli.Context) error {
	slog.Debugf("-> Sub Command action [ %s ] start", commandName)
	entry, err := withEntry(c)
	if err != nil {
		return err
	}
	commandEntry = entry
	return commandEntry.Exec()
}

func Command() []*cli.Command {
	urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())
	return []*cli.Command{
		{
			Name:   commandName,
			Usage:  "",
			Action: action,
			Flags:  urfave_cli.UrfaveCliAppendCliFlag(flag(), git_command.GitFlag()),
		},
	}
}
