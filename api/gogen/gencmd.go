package gogen

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/vars"
)

const (
	cmdDir  = "cmd"
	cmdFile = "root"
)

//go:embed cmd.tpl
var cmdTemplate string

func genCmd(dir string, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, cmdFile)
	if err != nil {
		return err
	}

	service := api.Service

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          cmdDir,
		filename:        filename + ".go",
		templateName:    "cmdTemplate",
		category:        category,
		templateFile:    cmdTemplateFile,
		builtinTemplate: cmdTemplate,
		data: map[string]string{
			"importPackages": genCmdImports(rootPkg),
			"serviceName":    service.Name,
		},
	})
}

func genCmdImports(parentPkg string) string {
	var imports []string
	imports = append(imports, fmt.Sprintf("\"%s/rest\"\n", vars.ProjectOpenSourceURL))
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, configDir)))
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, handlerDir)))
	imports = append(imports, fmt.Sprintf("\"%s\"\n", pathx.JoinPackages(parentPkg, contextDir)))
	return strings.Join(imports, "\n\t")
}
