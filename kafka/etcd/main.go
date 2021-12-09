package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	//put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	key := "collect_log_conf"
	value := `[{"path":"/tmp/logagent/mylog","topic":"mylog"},{"path":"/tmp/logagent/mylog1","topic":"mylog1"}]`

	_, err = client.Put(ctx, key, value)
	if err != nil {
		log.Fatal("put err: \n", err)
		return
	}
	cancel()

	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	gr, err := client.Get(ctx, key)
	if err != nil {
		log.Fatal("put err: \n", err)
		return
	}

	fmt.Printf("gr: %#v", gr)

	for _, ev := range gr.Kvs {
		fmt.Printf("key:%s\tvalue:%s", ev.Key, ev.Value)
	}
	cancel()

	// watch
	//watchCh := client.Watch(context.Background(), "foo")
	//for wresp := range watchCh {
	//for _, evt := range wresp.Events {
	//fmt.Printf("type:%s\tvalue:%s\tvalue:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
	//}
	//}
}
