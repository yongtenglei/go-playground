package settings

import (
	"fmt"

	"my_bubble/configs"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init(configPath string) (err error) {
	viper.SetConfigFile(configPath)

	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath(".")

	// 读取配置
	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	// 解析成结构体
	if err = viper.Unmarshal(configs.Conf); err != nil {
		fmt.Println("Read Config failed\n", err)
	}

	// 热加载配置
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config was Changed: \n", e.Name)
	})

	if err := viper.Unmarshal(configs.Conf); err != nil {
		fmt.Println("Read Config failed\n", err)
	}
	return

}
