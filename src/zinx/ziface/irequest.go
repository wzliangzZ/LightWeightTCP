package ziface

// 把客户端连接的 请求 和 数据 封装到request中
type IRequest interface {
	// 获得当前连接
	GetConnetcion() IConnection
	// 获得请求的消息数据
	GetData() []byte
	// 获得请求的消息的ID
	GetMsgId() uint32
}
