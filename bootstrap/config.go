package bootstrap

import (
	"fmt"
	"fpt/global"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitializeConfig() *viper.Viper {
	//设置配置文件路径
	config := "config.yaml"
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	//初始化viper
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s", err))
	}

	//监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed: ", in.Name)
		//重载配置
		if err := v.Unmarshal(&global.Ftpserver.Config); err != nil {
			fmt.Println(err)
		}
	})

	//配置文件赋值给全局变量
	if err := v.Unmarshal(&global.Ftpserver.Config); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Config initialize success!")

	return v
}
