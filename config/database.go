package config

//数据库配置
type Database struct {
	Driver       string `yaml:"driver"`
	DatabaseName string `yaml:"databasename"`
	DatabaseFile string `yaml:"databasefile"`
}
