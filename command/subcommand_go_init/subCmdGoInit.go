package subcommand_go_init

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/convention-change/zymosis/command"
	"github.com/convention-change/zymosis/command/git_command"
	"github.com/convention-change/zymosis/constant"
	"github.com/convention-change/zymosis/internal/urfave_cli"
	"github.com/convention-change/zymosis/internal/urfave_cli/cli_exit_urfave"
	"github.com/convention-change/zymosis/source"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"strings"
)

const commandName = "go"

var commandEntry *GolangCommand

type GolangCommand struct {
	isDebug      bool
	generateMode bool

	lengthHashShort uint

	isInit              bool
	isCoverageExitsFile bool
	gitCommand          git_command.GitCommand
	GoCodePath          string
}

func (n *GolangCommand) Exec() error {
	slog.Debugf("-> Exec subCommand [ %s ]", commandName)

	if !n.gitCommand.IsPathUnderGitManagement() {
		return cli_exit_urfave.Err(fmt.Errorf("not in git root path"))
	}

	if n.generateMode {

		generateTargetPath := filepath.Join(n.GoCodePath, source.PathTargetGo())
		dataRaw := []byte(n.gitCommand.HeadHashShort(n.lengthHashShort))
		errWriteGenerate := filepath_plus.WriteFileByByte(generateTargetPath, dataRaw, os.FileMode(0766), true)
		if errWriteGenerate != nil {
			return cli_exit_urfave.Err(errWriteGenerate)
		}

		slog.Infof("generate go source file at path: %s", generateTargetPath)
		return nil
	}

	rootPath := n.gitCommand.GitRootPath()

	if n.isInit {
		errCheckResource := source.CheckAllResource(rootPath)
		if errCheckResource != nil {
			return cli_exit_urfave.Err(errCheckResource)
		}
		getGoResource, errGetGoResource := source.GetGoResource()
		if errGetGoResource != nil {
			return cli_exit_urfave.Err(errGetGoResource)
		}
		for _, resource := range getGoResource {
			if resource.IsDir() {
				continue
			}
			dataRaw, errRaw := resource.Raw()
			if errCheckResource != nil {
				return cli_exit_urfave.Err(errRaw)
			}
			relativePath := resource.RelativePath()
			relativePath = strings.Replace(relativePath, source.DirGoZymosisSource, n.GoCodePath, 1)
			relativePath = strings.Replace(relativePath, source.DotFilePathMarkFrom, source.DotFilePathMarkTo, -1)
			pathResource := filepath.Join(rootPath, relativePath)

			if !filepath_plus.PathExistsFast(pathResource) {
				errWriteResource := filepath_plus.WriteFileByByte(pathResource, dataRaw, os.FileMode(0766), false)
				if errWriteResource != nil {
					return cli_exit_urfave.Err(errWriteResource)
				}
				slog.Infof("init go source file at path: %s", pathResource)
			} else {
				if n.isCoverageExitsFile {
					errWriteResource := filepath_plus.WriteFileByByte(pathResource, dataRaw, os.FileMode(0766), true)
					if errWriteResource != nil {
						return cli_exit_urfave.Err(errWriteResource)
					}
					slog.Infof("coverage go source file at path: %s", pathResource)
				} else {
					slog.Debugf("skip init go source file at path: %s", pathResource)
				}
			}
		}

		slog.Infof("-> init go source file done")
		return nil
	}

	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "go-package-path",
			Usage: "Set the kit go package path, defaults will use project root path",
		},
		&cli.BoolFlag{
			Name:  constant.NameInit,
			Usage: "init subcommand for project",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  constant.NameCoverageExistFile,
			Usage: "coverage exist file",
			Value: false,
		},
		&cli.StringFlag{
			Name:     "go-code-path",
			Usage:    "Set the kit go code path, defaults will use project root path",
			FilePath: "",
		},
		&cli.UintFlag{
			Name:  "length-hash-short",
			Usage: "Set the length of the short hash, defaults is 7",
			Value: 7,
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
		isDebug:    globalEntry.Verbose,
		gitCommand: gitCommand,

		generateMode:    globalEntry.RootCfg.GenerateMode,
		lengthHashShort: c.Uint("length-hash-short"),

		isInit:              c.Bool(constant.NameInit),
		isCoverageExitsFile: c.Bool(constant.NameCoverageExistFile),
		GoCodePath:          c.String("go-code-path"),
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
