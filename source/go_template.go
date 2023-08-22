package source

import (
	"embed"
	"github.com/convention-change/zymosis/constant"
	"github.com/convention-change/zymosis/internal/embed_source"
	"path"
	"path/filepath"
)

const (
	DirGoZymosisSource = "go_source"

	DirGoZymosisRecordName = "go_zymosis_record"

	GoCodeMarkName = "mark_git.go"
)

var (
	//go:embed go_source
	embedGoTemplateFs embed.FS
)

func PathTargetGo() string {
	return filepath.Join(DirZymosisName, DirGoZymosisRecordName, KeyZymosisMarkFile)
}

func initGoEmbedSource() error {
	//err := embed_source.InitResourceByDir(DirGoZymosisSource, embedGoTemplateFs, []string{
	//	path.Join(DirGoZymosisSource, DirZymosisName),
	//})
	//if err != nil {
	//	return err
	//}
	err := embed_source.InitResourceByDir(DirGoZymosisSource, embedGoTemplateFs, []string{
		path.Join(DirGoZymosisSource, DirZymosisName, DirGoZymosisRecordName),
	})
	if err != nil {
		return err
	}

	err = embed_source.InitResourceGroupByLanguage(DirGoZymosisSource, embedGoTemplateFs,
		path.Join(DirZymosisName, GoCodeMarkName),
		constant.SupportLanguage())
	if err != nil {
		return err
	}

	return nil
}

func GetGoResource() ([]embed_source.EmbedResource, error) {
	goResource, err := embed_source.GetResourceGroupByDir(DirGoZymosisSource, path.Join(DirZymosisName, DirGoZymosisRecordName))
	if err != nil {
		return nil, err
	}
	goCodeResource, err := embed_source.GetResourceByLanguageDefault(DirGoZymosisSource, path.Join(DirZymosisName, GoCodeMarkName))
	if err != nil {
		return nil, err
	}
	goResource = append(goResource, goCodeResource)
	return goResource, nil
}
