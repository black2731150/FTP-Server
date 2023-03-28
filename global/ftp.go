package global

import (
	"fpt/config"

	"github.com/spf13/viper"
)

type FtpServer struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
}

var Ftpserver = new(FtpServer)
