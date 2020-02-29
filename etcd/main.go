package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"playground"
	"time"
)

func main() {
	getSetRange()
	lease()
	watch()
}

// get, set, range
func getSetRange() {
	cli, err := clientv3.NewFromURL("http://localhost:2379")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := cli.Close()
		if err != nil {
			panic(err)
		}
	}()
	ctx := context.Background()

	putResp, err := cli.Put(ctx, "xuta", "xu")
	if err != nil {
		panic(err)
	}
	fmt.Println("just put (xuta, xu), put resp:", playground.ToJson(putResp))

	getResp, err := cli.Get(ctx, "xutao")
	if err != nil {
		panic(err)
	}
	fmt.Println("get xutao resp:", playground.ToJson(getResp))

	getResp, err = cli.Get(ctx, "a", clientv3.WithRange("z"))
	if err != nil {
		panic(err)
	}
	for i := range getResp.Kvs {
		fmt.Println("get range[a-z) resp:", getResp.Kvs[i])
	}
}

// lease
func lease() {
	cli, err := clientv3.NewFromURL("http://localhost:2379")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := cli.Close()
		if err != nil {
			panic(err)
		}
	}()
	ctx := context.Background()

	leaseResp, err := cli.Grant(ctx, 5)
	if err != nil {
		panic(err)
	}
	fmt.Println("lease ttl=5s resp:", playground.ToJson(leaseResp))

	putResp, err := cli.Put(ctx, "lease", "test", clientv3.WithLease(leaseResp.ID))
	if err != nil {
		panic(err)
	}
	fmt.Println("put (lease, test) with lease resp:", playground.ToJson(putResp))

	getResp, err := cli.Get(ctx, "lease")
	if err != nil {
		panic(err)
	}
	fmt.Println("get lease resp:", playground.ToJson(getResp))

	time.Sleep(time.Second * 10)

	getResp, err = cli.Get(ctx, "lease")
	if err != nil {
		panic(err)
	}
	fmt.Println("get lease after 10s resp:", playground.ToJson(getResp))
}

// watch
func watch() {
	cli, err := clientv3.NewFromURL("http://localhost:2379")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := cli.Close()
		if err != nil {
			panic(err)
		}
	}()
	ctx := context.Background()

	watch := cli.Watch(ctx, "xxt", clientv3.WithRev(1))

	go func() {
		for {
			r, ok := <-watch
			if !ok {
				return
			}
			fmt.Println("watch xxt(xxt is a new key) start with 1 resp:", playground.ToJson(r.Events))
		}
	}()

	putResp, err := cli.Put(ctx, "xxt", "test")
	if err != nil {
		panic(err)
	}
	fmt.Println("put resp:", playground.ToJson(putResp))

	time.Sleep(time.Second*5)
}
