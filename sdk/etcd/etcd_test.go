package test

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	recipe "github.com/coreos/etcd/contrib/recipes"
	"go-demo/utils/env"
	"runtime"
	"testing"
	"time"
)

var (
	ctx = context.TODO()
	cli *clientv3.Client
	err error
)

func TestMain(m *testing.M) {
	if env.IsCI() {
		return
	}
	cli, err = clientv3.New(clientv3.Config{
		// 集群列表
		Endpoints:   []string{"ip:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return
	}
	m.Run()
}

// etcd 增删改查
func TestEtcd(t *testing.T) {
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

// etcd 实现分布式锁
func TestEtcdLock(t *testing.T) {
	session, err := concurrency.NewSession(cli)
	if err != nil {
		t.Fatal(err)
	}
	// 使用 Lock
	locker := concurrency.NewLocker(session, "test")

	t.Log("请求锁...")
	locker.Lock()
	t.Log("获取到锁...")

	locker.Unlock()
	t.Log("释放锁")

	// Mutex, 请求锁的时候可以设置超时时间
	mux := concurrency.NewMutex(session, "test")
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	// 请求锁
	if err := mux.Lock(ctx); err != nil {
		t.Fatal(err)
	}

	// 释放锁
	if err := mux.Unlock(ctx); err != nil {
		t.Fatal(err)
	}

	// RWMutex 读写锁
	rwMux := recipe.NewRWMutex(session, "test")
	// 请求读锁
	if err := rwMux.RLock(); err != nil {
		t.Fatal(err)
	}
	// 释放读锁
	if err := rwMux.Unlock(); err != nil {
		t.Fatal(err)
	}
}

// etcd 节点选举
func TestEtcdLeader(t *testing.T) {
	session, err := concurrency.NewSession(cli)
	if err != nil {
		t.Fatal(err)
	}
	// 生成一个选举对象。下面主要使用它进行选举和查询等操作
	e1 := concurrency.NewElection(session, "test")
	// 把一个节点选举为Leader，并且会设置一个值
	if err := e1.Campaign(ctx, "pi"); err != nil {
		t.Fatal(err)
	}
	// 重新设置 Leader 的值，但是不会重新选主，
	if err := e1.Proclaim(ctx, "big"); err != nil {
		t.Fatal(err)
	}
	// 开启一次新的选举
	if err := e1.Resign(ctx); err != nil {
		t.Fatal(err)
	}

	// 获取当前主节点
	if leader, err := e1.Leader(ctx); err == nil {
		t.Log(leader)
	}

	// 每次主节点的变动都会生成一个新的版本号
	v := e1.Rev()
	t.Log("版本号:", v)

	// 监听主节点的变动
	go func() {
		ch := e1.Observe(ctx)
		for {
			select {
			case resp := <-ch:
				t.Log("主节点改变:", string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
			}
		}
	}()
}

// etcd 队列
func TestEtcdQueue(t *testing.T) {
	queue := recipe.NewQueue(cli, "test")
	// 入队
	for i := 0; i < 10; i++ {
		if err = queue.Enqueue(fmt.Sprintf("%d", i)); err != nil {
			t.Fatal(err)
		}
	}
	// 出队
	for i := 0; i < 10; i++ {
		if value, err := queue.Dequeue(); err == nil {
			t.Log(value)
		}
	}

	// 带优先级的队列，优先级高的元素会优先出队
	priQueue := recipe.NewPriorityQueue(cli, "test")
	// 入队
	for i := 0; i < 10; i++ {
		// 额外接受一个 pr 值（优先级）
		if err = priQueue.Enqueue(fmt.Sprintf("%d", i), uint16(i)); err != nil {
			t.Fatal(err)
		}
	}
	// 出队
	for i := 0; i < 10; i++ {
		if value, err := priQueue.Dequeue(); err == nil {
			t.Log(value)
		}
	}
}

// etcd栅栏
func TestEtcdBarrier(t *testing.T) {
	// 如果持有 Barrier 的节点释放了它，
	// 所有等待这个 Barrier 的节点就不会被阻塞，而是会继续执行。
	barrier := recipe.NewBarrier(cli, "test")

	// 创建一个栅栏
	if err := barrier.Hold(); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		go func(i int) {
			// 等待栅栏开启
			// 如果这个栅栏不存在，调用者不会被阻塞，而是会继续执行。
			if err := barrier.Wait(); err != nil {
				t.Fatal(err)
			}
			t.Log("go:", i)
		}(i)
	}

	// 开启栅栏
	if err := barrier.Release(); err != nil {
		t.Fatal(err)
	}
}

// etcd 计数型栅栏
func TestEtcdDoubleBarrier(t *testing.T) {
	session, err := concurrency.NewSession(cli)
	if err != nil {
		t.Fatal(err)
	}
	// 传递一个count
	doubleBar := recipe.NewDoubleBarrier(session, "test", 3)

	for i := 0; i < 3; i++ {
		go func(i int) {
			// 当调用 count 次，栅栏将会被开启
			// 编排一组节点，让这些节点在同一个时刻开始执行任务
			if err := doubleBar.Enter(); err != nil {
				t.Fatal(err)
			}
			t.Log("go:", time.Now().Nanosecond())
		}(i)
	}

	for i := 0; i < 3; i++ {
		go func(i int) {
			// 当调用 count 次，栅栏将会被开启
			// 编排一组节点，让这些节点在同一个时刻完成任务
			if err := doubleBar.Leave(); err != nil {
				t.Fatal(err)
			}
			t.Log("go:", time.Now().Nanosecond())
		}(i)
	}
}

// etcd 事务操作
func TestEtcdSTM(t *testing.T) {
	key := "test"
	// 定义一组事务操作
	exchange := func(stm concurrency.STM) error {
		stm.Put(key, "111")

		stm.Get(key)

		stm.Del(key)

		return nil
	}
	// 执行事务，可以保证这组操作要么全部成功，要么全部失败
	if _, err := concurrency.NewSTM(cli, exchange); err != nil {
		t.Fatal(err)
	}
}
