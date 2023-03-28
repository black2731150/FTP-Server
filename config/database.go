package config

type Database struct {
	Driver   string `yaml:"driver"`
	Database string `taml:"database"`
}
