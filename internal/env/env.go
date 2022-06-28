package env

import (
	"fmt"
	"main/internal/cfg"
	"main/internal/repositories"
	"os"
)

type Environment struct {
	C  cfg.Config
	LR repositories.Letters
	TR repositories.Tasks
}

var E *Environment

var (
	ConfigPath = "CONFIG_PATH"
)

func init() {
	// Get config path from environment variable
	path := os.Getenv(ConfigPath)
	if path == "" {
		path = "config.yaml"
	}

	var err error
	E, err = NewEnvironment(path)
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}
}

func NewEnvironment(yamlFile string) (*Environment, error) {
	conf, err := cfg.NewConfig(yamlFile)
	if err != nil {
		return nil, err
	}

	return &Environment{C: *conf}, nil
}
