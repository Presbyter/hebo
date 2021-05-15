package main

import (
	"flag"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "c", "./config.yaml", "配置文件路径")
}

func main() {
	flag.Parse()

	if "" == configPath {

	}
}
