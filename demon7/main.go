package main

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var (
	config clientv3.Config
	client *clientv3.Client
	err error
)

func main() {
	config = clientv3.Config{
		Endpoints:[]string{"10.16.34.85:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); nil != err {
		fmt.Println(err)
		return
	}

	client = client
}
