package config

//整个服务的配置
type Configuration struct {
	Ftp      Ftp `yaml:"ftp"`
	Database `yaml:"database"`
}
