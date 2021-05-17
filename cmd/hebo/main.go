package main

import (
	"flag"
	"github.com/Presbyter/hebo"
	"github.com/Presbyter/hebo/pkg/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

var (
	configPath string
	config     hebo.Config
)

func init() {
	flag.StringVar(&configPath, "config", "./config.yaml", "配置文件路径")
	flag.Parse()
}

func main() {
	l := log.New()
	l.SetOutput("stdout.log")
	l.Debug(configPath)

	cfgBytes, err := readConfig()
	if err != nil {
		l.Errorf("read config file fail. error: %s", err)
	}

	if err := yaml.Unmarshal(cfgBytes, &config); err != nil {
		l.Errorf("unmarshal config file fail. error: %s", err)
		return
	}
}

func readConfig() ([]byte, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	cfgBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return cfgBytes, nil
}
