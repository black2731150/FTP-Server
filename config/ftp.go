package config

//文件服务配置信息
type Ftp struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Name    string `yaml:"name"`
	Storage string `taml:"storage"`
}
