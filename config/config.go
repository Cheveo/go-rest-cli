package config

import (
	"embed"
	"io/fs"

	"github.com/Cheveo/go-rest-cli/types"
	"gopkg.in/yaml.v2"
)

//go:embed *.yaml
var data embed.FS

func GetConfigs() embed.FS {
	return data
}
func ReadConfig(config *types.Configuration, configName string) error {
	configs := GetConfigs()
	configFile, err := fs.ReadFile(configs, configName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		return err
	}

	return nil
}
