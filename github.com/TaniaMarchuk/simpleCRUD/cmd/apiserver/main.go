package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"goproject/simpleCRUD/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "github.com/TaniaMarchuk/simpleCRUD/configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatalf("error decoding config file: %v", err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		logrus.Fatalf("error starting apiserver: %v", err)
	}
}
