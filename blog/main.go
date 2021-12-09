package main

import (
	"blog/conf"
	"blog/dao/mysql"
	"blog/router"
	"blog/setting"
	"fmt"
	"net/http"
)

func main() {
	// 初始化viper， 获取配置文件
	if err := setting.Init("./conf/config.yaml"); err != nil {
		panic(err)
	}

	// 初始化mysql
	if err := mysql.Init(conf.Conf.MysqlConf); err != nil {
		panic(err)
	}

	// 初始化router
	r := router.RouterSetup()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Conf.ServerConf.HttpPort),
		Handler: r,
		//ReadTimeout:    time.Duration(conf.Conf.ServerConf.ReadTimeout),
		//WriteTimeout:   time.Duration(conf.Conf.ServerConf.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
