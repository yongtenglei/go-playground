package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type CollectEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

var (
	Client *clientv3.Client
)

func Init(address []string) (err error) {
	Client, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	return
}

func GetConf(key string) (collectEntryList *[]CollectEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := Client.Get(ctx, key)
	if err != nil {
		logrus.Errorf("get conf from etcd by key %s failed, err: %b\n", key, err)
		return
	}

	if len(resp.Kvs) == 0 {
		log.Warningf("get nothing from etcd by key %s", key)
		return
	}

	ret := resp.Kvs[0]
	if err := json.Unmarshal(ret.Value, &collectEntryList); err != nil {
		log.Errorf("json Unmarshal failed, err: %v", err)
	}
	return
}

func WatchConf(key string) {
	watchCh := Client.Watch(context.Background(), "foo")
	for wresp := range watchCh {
		for _, evt := range wresp.Events {
			fmt.Printf("type:%s\tvalue:%s\tvalue:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
		}
	}

}
