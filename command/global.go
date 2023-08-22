package command

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/convention-change/zymosis/constant"
	"github.com/convention-change/zymosis/internal/log"
	"github.com/convention-change/zymosis/internal/pkgJson"
	"github.com/convention-change/zymosis/internal/urfave_cli/cli_exit_urfave"
	"github.com/urfave/cli/v2"
)

type GlobalConfig struct {
	LogLevel      string
	TimeoutSecond uint

	GenerateMode bool
}

type (
	// GlobalCommand
	//	command root
	GlobalCommand struct {
		Name    string
		Version string
		Verbose bool
		RootCfg GlobalConfig
	}
)

var (
	cmdGlobalEntry *GlobalCommand
)

// CmdGlobalEntry
//
//	return global command entry
func CmdGlobalEntry() *GlobalCommand {
	return cmdGlobalEntry
}

// globalExec
//
//	do global command exec
func (c *GlobalCommand) globalExec() error {

	slog.Debug("-> start GlobalAction")

	return nil
}

// withGlobalFlag
//
// bind global flag to globalExec
func withGlobalFlag(c *cli.Context, cliVersion, cliName string) (*GlobalCommand, error) {
	slog.Debug("-> withGlobalFlag")

	isVerbose := c.Bool(constant.NameKeyCliVerbose)
	config := GlobalConfig{
		LogLevel:      c.String(constant.NameLogLevel),
		TimeoutSecond: c.Uint(constant.NamePluginTimeOut),

		GenerateMode: c.Bool(constant.NameGenerate),
	}

	p := GlobalCommand{
		Name:    cliName,
		Version: cliVersion,
		Verbose: isVerbose,
		RootCfg: config,
	}
	return &p, nil
}

// GlobalBeforeAction
// do command Action before flag global.
func GlobalBeforeAction(c *cli.Context) error {
	isVerbose := c.Bool(constant.NameKeyCliVerbose)
	err := log.InitLog(isVerbose, !isVerbose)
	if err != nil {
		panic(err)
	}
	cliVersion := pkgJson.GetPackageJsonVersionGoStyle(false)
	if isVerbose {
		slog.Warnf("-> open verbose, and now command version is: %s", cliVersion)
	}
	appName := pkgJson.GetPackageJsonName()
	cmdGlobalEntry, err = withGlobalFlag(c, cliVersion, appName)
	if err != nil {
		return cli_exit_urfave.Err(err)
	}

	return nil
}

// GlobalAction
// do command Action flag.
func GlobalAction(c *cli.Context) error {
	if cmdGlobalEntry == nil {
		panic(fmt.Errorf("not init GlobalBeforeAction success to new cmdGlobalEntry"))
	}

	err := cmdGlobalEntry.globalExec()
	if err != nil {
		return cli_exit_urfave.Format("run GlobalAction err: %v", err)
	}
	return nil
}

// GlobalAfterAction
//
//	do command Action after flag global.
//
//nolint:golint,unused
func GlobalAfterAction(c *cli.Context) error {
	isVerbose := c.Bool(constant.NameKeyCliVerbose)
	if isVerbose {
		slog.Infof("-> finish run command: %s, version %s", cmdGlobalEntry.Name, cmdGlobalEntry.Version)
	}
	return nil
}
