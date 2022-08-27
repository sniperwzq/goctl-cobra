package gogen

import (
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/vars"
)

const (
	configFile     = "config"
	configTemplate = `package config

import (
	"github.com/gogf/gf/v2/container/gmap"
	{{.authImport}}
)

const (
	defaultName = "gz.config"
)

var instances = gmap.NewStrAnyMap(true)

type Config struct {
	rest.RestConf
	{{.auth}}
	{{.jwtTrans}}
}

func LoadCfg(path string) *Config {
	return instances.GetOrSetFuncLock(defaultName, func() interface{} {
		var c Config
		conf.MustLoad(path, &c)
		return &c
	}).(*Config)
}

func Cfg() *Config {
	return instances.Get(defaultName).(*Config)
}
`

	jwtTemplate = ` struct {
		AccessSecret string
		AccessExpire int64
	}
`
	jwtTransTemplate = ` struct {
		Secret     string
		PrevSecret string
	}
`
)

func genConfig(dir string, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, configFile)
	if err != nil {
		return err
	}

	service := api.Service

	authNames := getAuths(api)
	var auths []string
	for _, item := range authNames {
		auths = append(auths, fmt.Sprintf("%s %s", item, jwtTemplate))
	}

	jwtTransNames := getJwtTrans(api)
	var jwtTransList []string
	for _, item := range jwtTransNames {
		jwtTransList = append(jwtTransList, fmt.Sprintf("%s %s", item, jwtTransTemplate))
	}

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          configDir,
		filename:        filename + ".go",
		templateName:    "configTemplate",
		category:        category,
		templateFile:    configTemplateFile,
		builtinTemplate: configTemplate,
		data: map[string]string{
			"authImport":  genConfigImports(rootPkg),
			"serviceName": service.Name,
			"auth":        strings.Join(auths, "\n"),
			"jwtTrans":    strings.Join(jwtTransList, "\n"),
		},
	})
}

func genConfigImports(parentPkg string) string {
	var imports []string
	imports = append(imports, fmt.Sprintf("\"%s/core/conf\"", vars.ProjectOpenSourceURL))
	imports = append(imports, fmt.Sprintf("\"%s/rest\"", vars.ProjectOpenSourceURL))
	return strings.Join(imports, "\n\t")
}
