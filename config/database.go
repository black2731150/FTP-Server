package config

//数据库配置
type Database struct {
	Driver   string `yaml:"driver"`
	Database string `taml:"database"`
}
