package config

import (
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// Config contains general app settings
type Config struct {
	Mode string `yaml:"mode"`
}

// RunOptions contains volatile settings from environment vars
type RunOptions struct {
	Port       string
	ConfigPath string
}

// LoadYamlConfig loads config from yaml file to struct
func LoadYamlConfig(path string) (Config, error) {
	var config Config
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	return config, err
}

// LoadRunOptions gets env vars with the prefix to RunOptions struct
func LoadRunOptions(prefix string) (RunOptions, error) {
	var ro RunOptions
	err := envconfig.Process(prefix, &ro)
	return ro, err
}
