package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

type Config struct {
	HTTPHost string `yaml:"host"`
	HTTPPort int    `yaml:"port"`
}

func NewConfig() *Config {
	file, err := os.Open("./config.yaml")
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		panic(err)
	}
	ret := &Config{}
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func (cfg *Config) ListenAddr() string {
	return fmt.Sprintf("%s:%d", cfg.HTTPHost, cfg.HTTPPort)
}
