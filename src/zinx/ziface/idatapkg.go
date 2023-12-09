package ziface

// 封包、拆包
// 面向tcp中额数据流，处理TCP粘包问题

type IDataPkg interface {
	GetHeadLen() uint32
	Pack(msg IMessage) ([]byte, error)
	Unpack([]byte) (IMessage, error)
}
