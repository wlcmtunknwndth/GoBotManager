package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

const (
	op            = "internal.config."
	configPathEnv = "config_path"
)

type Config struct {
	Env        string `yaml:"env" env-default:"prod"`
	TgToken    string `yaml:"tg_token" env-required:"true"`
	StaticPath string `yaml:"static_path" env-required:"true"`
}

func MustLoad() *Config {
	const scope = op + "MustLoad"

	path, ok := os.LookupEnv(configPathEnv)
	if !ok || path == "" {
		panic("config path is empty")
	}

	var err error
	if _, err = os.Stat(path); err != nil {
		panic(fmt.Sprintf("Config path %s doesn't exist: %s", path, err.Error()))
	}

	var cfg Config

	if err = cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("Couldn't read config file: " + err.Error())
	}

	return &cfg
}
