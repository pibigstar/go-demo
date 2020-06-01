package pool

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// 手写Mysql链接池

type Conn struct {
	context.Context
	maxConn        int                     // 最大连接数
	maxIdle        int                     // 最大可用连接数
	maxWait        int                     // 最大等待连接数
	maxWaitTimeout int                     // 最大等待连接数超时时间(毫秒)
	freeConn       int                     // 线程池空闲连接数
	connPool       []int                   // 连接池
	openCount      int                     // 已经打开的连接数
	waitConn       map[int]chan Permission // 排队等待的连接队列
	waitCount      int                     // 等待个数
	lock           sync.Mutex              // 锁
	nextConnIndex  NextConnIndex           // 下一个连接的ID标识（用于区分每个ID）
	freeConns      map[int]Permission      // 连接池的里的空闲连接
}

// 用其表示拿到了连接
type Permission struct {
	NextConnIndex               // 对应Conn中的NextConnIndex
	Content       string        // 通行证的具体内容，比如"PASSED"表示成功获取
	CreatedAt     time.Time     // 创建时间，即连接的创建时间
	MaxLifeTime   time.Duration // 连接的存活时间，本次没有用到这个属性，保留
}

type NextConnIndex struct {
	Index int
}

type Config struct {
	MaxConn        int
	MaxIdle        int
	MaxWait        int
	MaxWaitTimeout int
}

// 初始化连接参数
func Prepare(ctx context.Context, config *Config) (conn *Conn) {
	return &Conn{
		Context:        ctx,
		maxConn:        config.MaxConn,
		maxIdle:        config.MaxIdle,
		maxWait:        config.MaxWait,
		maxWaitTimeout: config.MaxWaitTimeout,
		openCount:      0,
		connPool:       []int{},
		waitConn:       make(map[int]chan Permission),
		waitCount:      0,
		freeConns:      make(map[int]Permission),
	}
}

// 创建连接
func (conn *Conn) New(ctx context.Context) (permission Permission, err error) {
	/**
	1、如果当前连接池非空，即len(freeConns)>0, 则直接从连接池获取一个连接
	2、如果连接池为空，则判定openConn是否大于maxConn，如果小于，则考虑创建新连接
	如果大于，判断当前等待队列是否大于最大等待数，如果大于则丢弃，如果小于则放到等待队列
	*/
	conn.lock.Lock()
	select {
	default:
	case <-ctx.Done():
		conn.lock.Unlock()
		return
	}

	// 连接池不为空，从连接池获取连接
	if len(conn.freeConns) > 0 {
		var (
			popPermission Permission
			popReqKey     int
		)

		// 获取其中一个连接
		for popReqKey, popPermission = range conn.freeConns {
			break
		}
		// 从连接池删除
		delete(conn.freeConns, popReqKey)

		fmt.Println("使用连接池中链接", "openCount: ", conn.openCount)
		conn.lock.Unlock()

		return popPermission, nil
	}

	// 当前连接数大于上限，则加入等待队列
	if conn.openCount >= conn.maxConn {
		// 如果超过最大等待数，则丢弃此次连接
		if conn.waitCount >= conn.maxWait {
			return Permission{}, errors.New("创建链接失败，等待队列已满")
		}

		nextConnIndex := getNextConnIndex(conn)
		req := make(chan Permission, 1)
		conn.waitConn[nextConnIndex] = req
		conn.waitCount++
		conn.lock.Unlock()

		select {
		// 如果在等待指定超时时间后，仍然无法获取释放连接，则放弃获取连接，这里如果不在超时时间后退出会一直阻塞
		case <-time.After(time.Millisecond * time.Duration(conn.maxWaitTimeout)):
			fmt.Println("等待超时，通知主线程退出....")
			return Permission{}, errors.New("wait conn timeout")
		// 有放回的连接, 直接拿来用
		case ret, ok := <-req:
			if !ok {
				return Permission{}, errors.New("创建新的连接失败，没有可用的链接释放")
			}
			fmt.Println("收到重新被释放的链接", "openCount: ", conn.openCount)
			return ret, nil
		}
	}

	// 新建连接
	conn.openCount++
	conn.lock.Unlock()
	permission = Permission{
		NextConnIndex: NextConnIndex{getNextConnIndex(conn)},
		Content:       "PASSED",
		CreatedAt:     time.Now(),
		MaxLifeTime:   time.Second * 5,
	}
	fmt.Println("创建新的连接", "openCount: ", conn.openCount)
	return permission, nil
}

func getNextConnIndex(conn *Conn) int {
	currentIndex := conn.nextConnIndex.Index
	conn.nextConnIndex.Index = currentIndex + 1
	return conn.nextConnIndex.Index
}

// 释放连接
func (conn *Conn) Release(ctx context.Context) (result bool, err error) {
	conn.lock.Lock()
	// 如果等待队列有等待任务，则通知正在阻塞等待获取连接的进程（即New方法中"<-req"逻辑）
	// 这里没有做指定连接的释放，只是保证释放的连接会被利用起来
	if len(conn.waitConn) > 0 {
		var req chan Permission
		var reqKey int
		for reqKey, req = range conn.waitConn {
			break
		}
		// 假定释放的连接就是下面新建的连接
		permission := Permission{
			NextConnIndex: NextConnIndex{reqKey},
			Content:       "PASSED", CreatedAt: time.Now(),
			MaxLifeTime: time.Second * 5,
		}
		req <- permission
		conn.waitCount--
		delete(conn.waitConn, reqKey)
		conn.lock.Unlock()
	} else {
		if conn.openCount > 0 {
			conn.openCount--

			if len(conn.freeConns) < conn.maxIdle { // 确保连接池大小不会超过maxIdle
				nextConnIndex := getNextConnIndex(conn)
				permission := Permission{
					NextConnIndex: NextConnIndex{nextConnIndex},
					Content:       "PASSED",
					CreatedAt:     time.Now(),
					MaxLifeTime:   time.Second * 5,
				}
				conn.freeConns[nextConnIndex] = permission
			}
		}
		conn.lock.Unlock()
	}
	return
}
