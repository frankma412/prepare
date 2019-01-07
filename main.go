package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
		kv     clientv3.KV
	)

	config = clientv3.Config{
		Endpoints:   []string{"10.16.34.85:2379"},
		DialTimeout: 5 * time.Second,
	}

	//创建一个客户端
	if client, err = clientv3.New(config); nil != err {
		fmt.Println(err)
		return
	}

	//用于读写etcd的键值对
	kv = clientv3.NewKV(client)
	if putResp, err := kv.Put(context.TODO(), "/cron/jobs/job1", "hello"); nil != err {
		fmt.Println(err)
	} else {
		fmt.Println("revision:", putResp.Header.Revision)
	}
}
