package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"

	"github.com/Presbyter/hebo"
	"github.com/Presbyter/hebo/pkg/log"
	"gopkg.in/yaml.v3"
)

var (
	configPath string
	config     *hebo.Config
)

func init() {
	flag.StringVar(&configPath, "config", "./config.yaml", "配置文件路径")
	flag.Parse()
}

func main() {
	l := log.New()
	// l.SetOutput("stdout.log")
	l.Debug(configPath)

	var err error
	config, err = parseConfig()
	if err != nil {
		panic(err)
	}

	server := hebo.New(config)
	server.SetLogger(l)
	server.Run(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	fmt.Printf("server exit now. %v", <-c)
}

//parseConfig 解析配置文件
func parseConfig() (*hebo.Config, error) {
	if configPath == "" {
		return nil, errors.New("config can not be empty")
	}

	if _, err := os.Stat(configPath); os.IsExist(err) {
		return nil, errors.New("config file not exist")
	}

	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cfgBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	cfg := hebo.Config{}
	if err := yaml.Unmarshal(cfgBytes, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
