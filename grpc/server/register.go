package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"log"
	"time"
)

var (
	stopSignal = make(chan bool, 1)
)

const (
	etcdAddr    = "http://localhost:2379"
	serviceName = "hello_world"
)

// registerToEtcd
func registerToEtcd(interval time.Duration, ttl int, port int) (string, error) {

	type Service struct {
		Addr     string `json:"Addr"`
		Metadate string `json:"Metadate"`
	}
	var (
		service = Service{
			Addr:     fmt.Sprintf("127.0.0.1:%d", port),
			Metadate: "...",
		}
		ctx          = context.Background()
		serviceKey   = fmt.Sprintf("%s/%s", serviceName, service.Addr)
		serviceValue string
	)

	bts, err := json.Marshal(service)
	if err != nil {
		return "", err
	}
	serviceValue = string(bts)

	log.Printf("register to etcd, key: %s, val: %s", serviceKey, serviceValue)

	// get endpoints for register dial address
	cli, err := clientv3.NewFromURL(etcdAddr)
	if err != nil {
		return "", fmt.Errorf("grpclb: create etcd3 client failed: %v", err)
	}

	go func() {
		// invoke self-register with ticker
		ticker := time.NewTicker(interval)
		for {
			// minimum lease TTL is ttl-second
			resp, err := cli.Grant(ctx, int64(ttl))
			if err != nil {
				log.Printf("grant fail, err: %v", err)
				continue
			}
			// should get first, if not exist, set it
			_, err = cli.Get(ctx, serviceKey)
			if err != nil {
				if err == rpctypes.ErrKeyNotFound {
					if _, err := cli.Put(ctx, serviceKey, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
						log.Printf("grpclb: set service with ttl to etcd3 failed: %s", err.Error())
					}
				} else {
					log.Printf("grpclb: service connect to etcd3 failed: %s", err.Error())
				}
			} else {
				// refresh set to true for not notifying the watcher
				if _, err := cli.Put(ctx, serviceKey, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
					log.Printf("grpclb: refresh service with ttl to etcd3 failed: %s", err.Error())

				}
			}
			select {
			case <-stopSignal:
				err = cli.Close()
				if err != nil {
					panic(err)
				}
				return
			case <-ticker.C:
			}
		}
	}()

	return serviceKey, nil
}

// UnRegisterToEtcd delete registered service from etcd
func UnRegisterToEtcd(serviceKey string) error {
	cli, cerr := clientv3.NewFromURL(etcdAddr)
	if cerr != nil {
		return fmt.Errorf("grpclb: create etcd3 client failed: %v", cerr)
	}
	stopSignal <- true
	stopSignal = make(chan bool, 1) // just a hack to avoid multi UnRegisterToEtcd deadlock
	var err error
	if _, err := cli.Delete(context.Background(), serviceKey); err != nil {
		log.Printf("grpclb: deregister '%s' failed: %s", serviceKey, err.Error())
	} else {
		log.Printf("grpclb: deregister '%s' ok.", serviceKey)
	}
	err = cli.Close()
	if err != nil {
		panic(err)
	}

	return err
}
