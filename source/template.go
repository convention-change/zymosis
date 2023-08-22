package source

import "github.com/convention-change/zymosis/internal/embed_source"

const (
	DotFilePathMarkFrom = "/dot_"
	DotFilePathMarkTo   = "/."

	KeyZymosisMarkFile = "git_rev_parse"

	DirZymosisName = "zymosis"
)

func CheckAllResource(root string) error {
	embed_source.SettingResourceRootPath(root)

	err := initGoEmbedSource()
	if err != nil {
		return err
	}
	return nil
}
