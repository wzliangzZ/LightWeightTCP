package ziface

import "net"

// 定义连接模块的抽象层		一个连接绑定一个，绑定什么业务就做什么事情
type IConnection interface {
	// 启动连接
	Start()
	// 停止连接
	Stop()
	// 获取scoket coon
	GetTCPConnection() *net.TCPConn
	// 获取连接ID
	GetConnID() uint32
	// 获取客户端TCP IP，port
	RemoteAddr() net.Addr
	// 发送数据
	SendMsg(uint32, []byte) error

	// 设置连接属性
	SetProperty(string, any)
	// 获取连接属性
	GetProperty(string) (any, error)
	// 移除连接属性
	RemoveProperty(string)
}

