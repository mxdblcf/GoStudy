package etcdStudy

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"

	"time"
)

func Use_etcd()  {


	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
		if err!=nil{
			fmt.Println("cant connect")
			return
		}

		fmt.Println("connect success")
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	get, err := client.Get(ctx, "k")
	count := get.Count
	fmt.Println(count)
cancel()
	for _, ev := range get.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}
