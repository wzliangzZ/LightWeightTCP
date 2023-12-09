package znet

import (
	"fmt"
	"net"
	"zx/src/zinx/utils"
	"zx/src/zinx/ziface"
)

// iServer接口的实现，定义一个Server服务器模块
type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的IP
	IP string
	// 服务器监听的端口
	Port int
	//server注册的连接对应的处理业务
	MsgHandler ziface.IMsgHandle
	// 当前server管理的conn
	ConnManager ziface.IConnmanager

	// 连接启动前后调用方法
	OnConnStart func(ziface.IConnection)
	OnConnStop  func(ziface.IConnection)
}

// 默认参数
const (
	defalutIPVersion = "tcp4"
)

// 初始化Server
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:        utils.Global_obj.Name,
		IPVersion:   defalutIPVersion,
		IP:          utils.Global_obj.Host,
		Port:        utils.Global_obj.TcpPort,
		MsgHandler:  NewMsgHandle(),
		ConnManager: NewConnManager(),
	}
	return s
}

func (s *Server) AddRouter(id uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(id, router)
}

func (s *Server) Start() {
	fmt.Printf("[zinx] server name %s listenner at ip %s, port %d\n", s.Name, s.IP, s.Port)

	s.MsgHandler.StartWorkerPool()
	fmt.Println("workpoll start")

	go func() {

		// 1.获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error:", err)
			return
		}

		// 2.监听服务器地址
		lis, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen error:", s.IPVersion, err)
			return
		}
		fmt.Println("start Zinx server succ,", s.Name, "succ Listennging...")
		var cid uint32
		// 3.阻塞，等待客户端连接，处理客户连接断连接业务（读写）
		for {
			// 有客户端连接过来，阻塞会返回
			conn, err := lis.AcceptTCP()

			if err != nil {
				fmt.Println("accept error:", err)
				continue
			}
			// 连接是否达到上限
			if s.ConnManager.Len() >= utils.Global_obj.MaxConn {
				fmt.Println("超过最大服务器连接！！！")
				conn.Close()
				continue
			}
			// 连接与业务绑定
			mycoon := NewConnection(s, conn, cid, s.MsgHandler)
			cid++
			fmt.Println("当前CONN管理数量：", s.GetConnManager().Len())
			go mycoon.Start()

		}
	}()

}

func (s *Server) Stop() {
	fmt.Println("[stop] zinx server name", s.Name)
	s.ConnManager.ClearConn()

}

func (s *Server) Serve() {
	// 启动服务
	s.Start()

	// 阻塞状态
	select {}
}

func (s *Server) GetConnManager() ziface.IConnmanager {
	return s.ConnManager
}

func (s *Server) SetOnConnStart(f func(ziface.IConnection)) {
	s.OnConnStart = f
}
func (s *Server) SetOnConnStop(f func(ziface.IConnection)) {
	s.OnConnStop = f
}

// 调用方法
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("succ onconnstart...")
		s.OnConnStart(conn)
	}

}
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("succ onconnstop...")
		s.OnConnStop(conn)
	}

}
