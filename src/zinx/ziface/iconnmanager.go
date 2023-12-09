package ziface


// 连接管理模块
type IConnmanager interface {
	Add(conn IConnection)
	Remove(conn IConnection)
	Get(uint32) (IConnection, error)
	Len() int
	ClearConn()
}