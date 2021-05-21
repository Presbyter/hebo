package hebo

type Config struct {
	Port     uint   `yaml:"port"`
	IpBind   string `yaml:"ipbind"`
	UpStream []struct {
		Address string `yaml:"address"`
		Weight  uint32 `yaml:"weight"`
	} `yaml:"upstream"`
}
