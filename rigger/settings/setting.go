package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 读取配置
	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("Read config failed", err)
		return
	}

	// 热加载配置
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config was Changed")
	})
	return

}
