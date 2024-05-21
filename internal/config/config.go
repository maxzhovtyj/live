package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type Config struct {
	StaticDirPath string `yaml:"staticDirPath"`
	Hostname      string `yaml:"hostname"`
	DBConnection  string `yaml:"DBConnection"`
}

var (
	config = flag.String("config", "config.yml", "Path to config file")

	cfg  *Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		rawCfg, err := os.ReadFile(*config)
		if err != nil {
			panic(err)
		}

		if err = yaml.Unmarshal(rawCfg, &cfg); err != nil {
			panic(err)
		}
	})

	return cfg
}
