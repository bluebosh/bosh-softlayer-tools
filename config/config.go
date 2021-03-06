package config

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigInfo struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	TargetUrl string `json:"target_url"`
}

type Config interface {
	Path() string
	LoadConfig() (ConfigInfo, error)
	SaveConfig(configInfo ConfigInfo) error
}

const (
	CONFIG_FILE_NAME = ".bmp_config"
	CONFIG_PATH      = "~/" + CONFIG_FILE_NAME
)

type config struct {
	configInfo ConfigInfo
	path       string
}

func NewConfig(path string) *config {
	if path == "" {
		path = CONFIG_PATH
	}

	return &config{
		configInfo: ConfigInfo{},
		path:       path,
	}
}

func (c *config) Path() string {
	return c.path
}

func (c *config) LoadConfig() (ConfigInfo, error) {
	configFileContents, err := ioutil.ReadFile(c.path)
	if err != nil {
		return ConfigInfo{}, err
	}

	configInfo := ConfigInfo{}
	err = json.Unmarshal(configFileContents, &configInfo)
	if err != nil {
		return ConfigInfo{}, err
	}

	return configInfo, nil
}

func (c *config) SaveConfig(configInfo ConfigInfo) error {
	contents, err := json.Marshal(configInfo)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(c.path, contents, 0666)
	if err != nil {
		return err
	}

	return nil
}
