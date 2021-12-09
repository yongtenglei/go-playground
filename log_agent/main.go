package main

import (
	"log_agent/config"
	"log_agent/etcd"
	"log_agent/kafka"
	"log_agent/tailfile"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

func run() {
	select {}

}

func main() {
	// Load config file
	c := new(config.Config)
	if err := ini.MapTo(c, "config/config.ini"); err != nil {
		log.Fatalf("Loading config err: %v\n", err)
	}
	log.Infoln("Loading config successfully")

	// Init kafka
	if err := kafka.Init(strings.Split(c.KafkaConfig.Address, ","), c.KafkaConfig.ChanSize); err != nil {
		log.Fatalf("Init kafka err: %v\n", err)
	}
	log.Infoln("Init kafka successfully")
	defer kafka.Client.Close()

	// Init etcd
	if err := etcd.Init(strings.Split(c.EtcdConfig.Address, ",")); err != nil {
		log.Fatalf("Init etcd err: %v\n", err)
	}
	log.Infoln("Init etcd successfully")
	defer etcd.Client.Close()

	confList, err := etcd.GetConf(c.EtcdConfig.CollectKey)
	if err != nil {
		log.Fatalf("Get etcd confList err: %v\n", err)
	}
	log.Infoln("Get etcd confList successfully")
	log.Infoln("confList: ", confList)

	// Watch config
	go etcd.WatchConf(c.EtcdConfig.CollectKey)

	// Init tail
	if err := tailfile.Init(confList); err != nil {
		log.Fatalf("Init tail err: %v\n", err)
	}
	log.Infoln("Init tail successfully")
	defer kafka.Client.Close()

	// work....
	run()
}
