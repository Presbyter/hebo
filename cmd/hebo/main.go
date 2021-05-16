package main

import (
	"flag"
	"github.com/Presbyter/hebo/pkg/log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "./config.yaml", "配置文件路径")
	flag.Parse()
}

func main() {
	l := log.New()
	l.SetOutput("stdout.log")
	l.Debug(configPath)
}
