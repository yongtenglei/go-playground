package settings

import (
	"app/configs"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init(ConfigPath string) (err error) {
	viper.SetConfigFile(ConfigPath)
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath(".")

	// 读取配置

	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("Read config failed", err)
		return
	}

	if err = viper.Unmarshal(configs.Conf); err != nil {
		fmt.Println("Read Config failed\n", err)
	}

	// 热加载配置
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config was Changed", in.Name)
		if err := viper.Unmarshal(configs.Conf); err != nil {
			fmt.Println("Read Config failed\n", err)
		}
	})
	return

}
