package main

import (
	"fmt"
	"go-blog/conf"
	"go-blog/dao/mysql"
	"go-blog/models"
	"go-blog/pkg/setting"
	"go-blog/router"
	"log"
	"net/http"
)

func main() {
	if err := setting.Init("./conf/config.yaml"); err != nil {
		log.Println("Init vaper failed")
	}

	if err := mysql.Init(conf.Conf.MysqlConf); err != nil {
		log.Println("Init mysql failed")
	}

	db := mysql.Db()
	db.AutoMigrate(models.Tag{}, models.Article{}, models.User{})

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
