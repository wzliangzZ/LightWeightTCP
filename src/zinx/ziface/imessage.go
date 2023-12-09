package ziface

// 将请求的消息封装，定义一个抽象模块

type IMessage interface {
	GetMsgId() uint32
	GetMsgLen() uint32
	GetData() []byte

	SetMsgId(uint32)
	SetData([]byte)
	SetDataLen(uint32)
}
