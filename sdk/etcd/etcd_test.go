package test

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"go-demo/utils/env"
	"runtime"
	"testing"
	"time"
)

var (
	ctx = context.TODO()
)

func TestEtcd(t *testing.T) {
	if env.IsCI() {
		return
	}
	cli, err := clientv3.New(clientv3.Config{
		// 集群列表
		Endpoints:   []string{"ip:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()

	// 监听值
	go func() {
		watch := cli.Watch(ctx, "name")
		res := <-watch
		t.Log("name发生改变", res)
	}()

	// 存值
	if resp, err := cli.Put(ctx, "name", "Hello", clientv3.WithPrevKV()); err != nil {
		t.Error(err)
	} else {
		t.Log("旧值: ", string(resp.PrevKv.Value))
	}
	// 取值
	if resp, err := cli.Get(ctx, "name", clientv3.WithPrefix()); err != nil {
		t.Error(err)
	} else {
		t.Log("count: ", resp.Count)
		t.Log("value: ", resp.Kvs)
	}

	// 改值
	if resp, err := cli.Put(ctx, "name", "pibigstar", clientv3.WithPrevKV()); err != nil {
		t.Error(err)
	} else {
		t.Log("旧值: ", string(resp.PrevKv.Value))
	}
	// 删值
	if resp, err := cli.Delete(ctx, "name"); err != nil {
		t.Error(err)
	} else {
		t.Log(resp.PrevKvs)
	}

	// 带租期的key
	lease := clientv3.NewLease(cli)
	// 申请一个5秒的租约(5s后key会被删除)
	if response, err := lease.Grant(ctx, 5); err != nil {
		t.Error(err)
	} else {

		// 自动续约
		if responses, err := lease.KeepAlive(ctx, response.ID); err == nil {
			go func() {
				for {
					select {
					case keepResp := <-responses:
						if keepResp == nil {
							t.Log("租约已失效或context已取消")
							runtime.Goexit()
						} else {
							t.Log("自动续约...")
						}
					}
				}
			}()
		}

		if _, err := cli.Put(ctx, "age", "18", clientv3.WithLease(response.ID)); err != nil {
			t.Error(err)
		}
	}
}
