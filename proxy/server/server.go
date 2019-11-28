//server端，运行在有外网ip的服务器上
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
	"time"
)

var (
	localPort  int
	remotePort int
)

func init() {
	flag.IntVar(&localPort, "lp", 3002, "user访问地址端口")
	flag.IntVar(&remotePort, "rp", 20012, "与client通讯端口")
}

//与client相关的conn
type client struct {
	conn net.Conn
	// 错误通道
	err chan bool
	//未收到心跳包通道
	heart chan bool
	//暂未使用！！！原功能tcp连接已经接通，不在需要心跳包
	disheart bool

	//关闭写通道标识
	closeWrite chan bool
	// 数据传输通道
	recv chan []byte
	send chan []byte
}

//读取从client过来的数据
func (client *client) read() {
	for {
		//40秒没有数据传输则断开
		client.conn.SetReadDeadline(time.Now().Add(time.Second * 40))
		recv := make([]byte, 10240)
		n, err := client.conn.Read(recv)

		if err != nil {
			fmt.Printf("read from client, err: %s \n", err.Error())
			client.heart <- true
			client.err <- true
			client.closeWrite <- true

			fmt.Println("长时间未传输信息，或者client已关闭。断开并继续accept新的tcp，", err)
		}
		//收到心跳包hh，原样返回回复
		if recv[0] == 'h' && recv[1] == 'h' {
			client.conn.Write([]byte("hh"))
			continue
		}
		client.recv <- recv[:n]

	}
}

//把数据发送给client
func (client client) write() {
	for {
		send := make([]byte, 10240)
		select {
		case send = <-client.send:
			client.conn.Write(send)
		case <-client.closeWrite:
			fmt.Println("写入client进程关闭")
			break
		}
	}
}

//处理心跳包
func (client client) cHeart() {
	for {
		client.conn.SetReadDeadline(time.Now().Add(time.Second * 30))

		recv := make([]byte, 2)
		chanRecv := make(chan []byte)
		_, err := client.conn.Read(recv)
		chanRecv <- recv
		if err != nil {
			client.heart <- true
			fmt.Println("心跳包超时", err)
			break
		}
		// 如果收到心跳包，则回应一个心跳包
		if recv[0] == 'h' && recv[1] == 'h' {
			client.conn.Write([]byte("hh"))
		}

	}
}

//与user相关的conn
type user struct {
	conn net.Conn
	// 标识位通道
	err        chan bool
	closeWrite chan bool

	// 数据传输通道
	recv chan []byte
	send chan []byte
}

//读取从user过来的数据
func (u user) read() {
	u.conn.SetReadDeadline(time.Now().Add(time.Millisecond * 800))
	for {
		recv := make([]byte, 10240)
		n, err := u.conn.Read(recv)
		u.conn.SetReadDeadline(time.Time{})
		if err != nil {
			u.err <- true
			u.closeWrite <- true
			break
		}
		u.recv <- recv[:n]
	}
}

//把数据发送给user
func (u user) write() {
	for {
		send := make([]byte, 10240)
		select {
		case send = <-u.send:
			u.conn.Write(send)
		case <-u.closeWrite:
			fmt.Println("写入user进程关闭")
			break
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if !(localPort >= 0 && localPort < 65536) {
		fmt.Println("端口设置错误")
		os.Exit(1)
	}
	if !(remotePort >= 0 && remotePort < 65536) {
		fmt.Println("端口设置错误")
		os.Exit(1)
	}

	//监听端口
	remoteLister, err := net.Listen("tcp", fmt.Sprintf(":%d", remotePort))
	if err != nil {
		panic(err)
	}

	localLister, err := net.Listen("tcp", fmt.Sprintf(":%d", localPort))
	if err != nil {
		panic(err)
	}

	//第一条tcp关闭或者与浏览器建立tcp都要返回重新监听
TOP:
	//监听本地user链接
	userConn := make(chan net.Conn)
	go localAccept(localLister, userConn)

	// 等待Client端的连接
	clientConn, err := remoteLister.Accept()
	if err != nil {
		panic(err)
	}

	fmt.Println("client已连接", clientConn.LocalAddr().String())
	recv := make(chan []byte)
	send := make(chan []byte)
	heart := make(chan bool, 1)
	//大小为1是为了防止两个读取线程一个退出后另一个永远卡住
	er := make(chan bool, 1)
	writ := make(chan bool)
	client := &client{clientConn, er, heart, false, writ, recv, send}
	go client.read()
	go client.write()

	//这里可能需要处理心跳
	for {
		select {
		case <-client.heart:
			goto TOP
		case userConn := <-userConn:
			//暂未使用
			client.disheart = true
			recv = make(chan []byte)
			send = make(chan []byte)
			//大小为1是为了防止两个读取线程一个退出后另一个永远卡住
			er = make(chan bool, 1)
			writ = make(chan bool)
			user := &user{userConn, er, writ, recv, send}
			go user.read()
			go user.write()
			//当两个socket都创立后进入handle处理
			go handle(client, user)
			goto TOP
		}

	}

}

//监听端口函数
func remoteAccept(con net.Listener) net.Conn {
	client, err := con.Accept()
	if err != nil {
		panic(err)
	}
	return client
}

//在另一个进程监听端口函数
func localAccept(con net.Listener, userConn chan net.Conn) {
	client, err := con.Accept()
	if err != nil {
		panic(err)
	}
	userConn <- client
}

//两个socket衔接相关处理
func handle(client *client, user *user) {
	for {
		var clientRecv = make([]byte, 10240)
		var userRecv = make([]byte, 10240)
		select {

		case clientRecv = <-client.recv:
			user.send <- clientRecv
		case userRecv = <-user.recv:
			//fmt.Println("浏览器发来的消息", string(userRecv))
			client.send <- userRecv
			//user出现错误，关闭两端socket
		case <-user.err:
			//fmt.Println("user关闭了，关闭client与user")
			close(client.conn)
			close(user.conn)
			runtime.Goexit()
			//client出现错误，关闭两端socket
		case <-client.err:
			//fmt.Println("client关闭了，关闭client与user")
			close(client.conn)
			close(user.conn)
			runtime.Goexit()
		}
	}
}

func close(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		panic(err)
	}
}
