package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	localPort  int
	remotePort int
	host       string
)

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "服务器ip地址")
	flag.IntVar(&localPort, "lp", 8080, "穿透本地端口号")
	flag.IntVar(&remotePort, "rp", 20012, "远程Server端口")
}

//与browser相关的conn
type browser struct {
	conn net.Conn
	// 标识位通道
	err        chan bool
	closeWrite chan bool
	// 数据传输通道
	recv chan []byte
	send chan []byte
}

//读取browser过来的数据
func (b browser) read() {
	for {
		recv := make([]byte, 10240)
		n, err := b.conn.Read(recv)
		if err != nil {
			b.closeWrite <- true
			b.err <- true
			fmt.Println("读取browser失败", err)
			break
		}
		b.recv <- recv[:n]

	}
}

//把数据发送给browser
func (b browser) write() {

	for {
		var send []byte = make([]byte, 10240)
		select {
		case send = <-b.send:
			b.conn.Write(send)
		case <-b.closeWrite:
			fmt.Println("写入browser进程关闭")
			break

		}

	}

}

//与server相关的conn
type server struct {
	conn net.Conn
	er   chan bool
	writ chan bool
	recv chan []byte
	send chan []byte
}

//读取server过来的数据
func (s *server) read() {
	//isheart与timeout共同判断是不是自己设定的SetReadDeadline
	var isHeart bool = false
	//20秒发一次心跳包
	_ = s.conn.SetReadDeadline(time.Now().Add(time.Second * 20))
	for {
		var recv []byte = make([]byte, 10240)
		n, err := s.conn.Read(recv)
		if err != nil {
			if strings.Contains(err.Error(), "timeout") && !isHeart {
				//fmt.Println("发送心跳包")
				_, _ = s.conn.Write([]byte("hh"))
				//4秒时间收心跳包
				_ = s.conn.SetReadDeadline(time.Now().Add(time.Second * 4))
				isHeart = true
				continue
			}
			//浏览器有可能连接上不发消息就断开，此时就发一个0，为了与服务器一直有一条tcp通路
			s.recv <- []byte("0")
			s.er <- true
			s.writ <- true
			fmt.Println("没收到心跳包或者server关闭，关闭此条tcp", err)
			break
		}
		//收到心跳包
		if recv[0] == 'h' && recv[1] == 'h' {
			//fmt.Println("收到心跳包")
			_ = s.conn.SetReadDeadline(time.Now().Add(time.Second * 20))
			isHeart = false
			continue
		}
		s.recv <- recv[:n]
	}
}

//把数据发送给server
func (s server) write() {

	for {
		send := make([]byte, 10240)

		select {
		case send = <-s.send:
			_, _ = s.conn.Write(send)
		case <-s.writ:
			fmt.Println("写入server进程关闭")
			break
		}

	}

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	if !(localPort >= 0 && localPort < 65536) {
		fmt.Println("端口设置错误")
		os.Exit(1)
	}
	if !(remotePort >= 0 && remotePort < 65536) {
		fmt.Println("端口设置错误")
		os.Exit(1)
	}

	target := net.JoinHostPort(host, fmt.Sprintf("%d", remotePort))
	for {
		//链接端口
		serverConn := dial(target)
		recv := make(chan []byte)
		send := make(chan []byte)
		//1个位置是为了防止两个读取线程一个退出后另一个永远卡住
		er := make(chan bool, 1)
		writ := make(chan bool)
		next := make(chan bool)
		server := &server{serverConn, er, writ, recv, send}
		go server.read()
		go server.write()
		go handle(server, next)
		<-next
	}

}

//链接端口
func dial(hostPort string) net.Conn {
	conn, err := net.Dial("tcp", hostPort)
	if err != nil {
		fmt.Println(hostPort)
		panic(err)
	}
	return conn
}

//两个socket衔接相关处理
func handle(server *server, next chan bool) {
	var serverRecv = make([]byte, 10240)
	//阻塞这里等待server传来数据再链接browser
	fmt.Println("等待server发来消息")
	serverRecv = <-server.recv
	//连接上，下一个tcp连上服务器
	next <- true
	//fmt.Println("开始新的tcp链接，发来的消息是：", string(serverRecv))
	var browse *browser
	//server发来数据，链接本地80端口
	serverConn := dial(fmt.Sprintf("127.0.0.1:%d", localPort))
	recv := make(chan []byte)
	send := make(chan []byte)
	er := make(chan bool, 1)
	writ := make(chan bool)
	browse = &browser{serverConn, er, writ, recv, send}
	go browse.read()
	go browse.write()
	browse.send <- serverRecv

	for {
		var serverRecv = make([]byte, 10240)
		var browserRecv = make([]byte, 10240)
		select {
		case serverRecv = <-server.recv:
			if serverRecv[0] != '0' {

				browse.send <- serverRecv
			}

		case browserRecv = <-browse.recv:
			server.send <- browserRecv
		case <-server.er:
			//fmt.Println("server关闭了，关闭server与browse")
			close(server.conn)
			close(browse.conn)
			runtime.Goexit()
		case <-browse.err:
			//fmt.Println("browse关闭了，关闭server与browse")
			close(server.conn)
			close(browse.conn)
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
