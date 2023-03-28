package global

import (
	"fpt/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type FtpServer struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	DB          *gorm.DB
}

var Ftpserver = new(FtpServer)
