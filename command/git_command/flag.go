package git_command

import (
	"fmt"
	"github.com/sinlov-go/go-git-tools/git_info"
	"github.com/urfave/cli/v2"
	"os"
)

type GitCommand interface {
	GitRootPath() string

	GitRemote() string

	IsPathUnderGitManagement() bool

	HeadBranchShort() string

	HeadHashShort(length uint) string

	headHash() string
}

type gitCommandEntry struct {
	gitRootPath string
	gitRemote   string
}

func (g *gitCommandEntry) HeadHashShort(length uint) string {
	if len(g.headHash()) == 0 {
		return ""
	}
	return g.headHash()[:length]
}

func (g *gitCommandEntry) headHash() string {
	headReference, err := git_info.RepositoryHeadByPath(g.gitRootPath)
	if err != nil {
		return ""
	}
	if headReference == nil {
		return ""
	}
	return headReference.Hash().String()
}

func (g *gitCommandEntry) HeadBranchShort() string {
	headReference, err := git_info.RepositoryHeadByPath(g.gitRootPath)
	if err != nil {
		return ""
	}
	if headReference == nil {
		return ""
	}
	return headReference.Name().Short()
}

func (g *gitCommandEntry) IsPathUnderGitManagement() bool {
	_, err := git_info.IsPathGitManagementRoot(g.gitRootPath)
	if err != nil {
		return false
	}
	return true
}

func (g *gitCommandEntry) GitRootPath() string {
	return g.gitRootPath
}

func (g *gitCommandEntry) GitRemote() string {
	return g.gitRemote
}

func BindGitFlag(c *cli.Context) (GitCommand, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("can not get pwdfoler err: %v", err)
	}
	gitRootFolder := dir

	return &gitCommandEntry{
		gitRootPath: gitRootFolder,
		gitRemote:   c.String("git-remote"),
	}, nil
}

func GitFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "git-remote",
			Usage: "change git remote name. default origin",
			Value: "origin",
		},
	}
}
