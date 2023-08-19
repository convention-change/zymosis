//go:build !test

package main

import (
	"fmt"
	"github.com/convention-change/zymosis"
	"github.com/convention-change/zymosis/cmd/cli"
	"github.com/convention-change/zymosis/internal/pkgJson"
	"github.com/gookit/color"
	os "os"
)

const (
	exitCodeCmdArgs = 2
)

func main() {
	pkgJson.InitPkgJsonContent(zymosis.PackageJson)

	app := cli.NewCliApp()

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("%s %s --help\n", color.Yellow.Render("please see help as:"), app.Name)
		os.Exit(exitCodeCmdArgs)
	}
	if err := app.Run(args); nil != err {
		color.Redf("cli err at %v\n", err)
	}
}
