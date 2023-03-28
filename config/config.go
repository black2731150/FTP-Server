package config

type Configuration struct {
	Ftp      Ftp `yaml:"ftp"`
	Database `yaml:"database"`
}
