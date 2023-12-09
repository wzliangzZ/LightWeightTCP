package znet

import (
	"zx/src/zinx/ziface"
)

type Request struct {
	Conn ziface.IConnection
	msg  ziface.IMessage
}

// 获得当前连接
func (r *Request) GetConnetcion() ziface.IConnection {
	return r.Conn
}

// 获得请求的消息数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}
