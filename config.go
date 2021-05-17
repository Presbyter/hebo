package hebo

type Config struct {
	UpStream []struct {
		Address string `yaml:"address"`
		Weight  uint8  `yaml:"weight"`
	} `yaml:"upstream"`
}
