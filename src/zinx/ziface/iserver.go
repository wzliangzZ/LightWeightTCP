package ziface

//定义一个服务接口
type IServer interface {
	// 启动服务器
	Start()
	// 停止服务器
	Stop()
	// 运行服务器
	Serve()
	// 添加路由，供客户端连接使用
	AddRouter(uint32, IRouter)
	// 获取当前conManager
	GetConnManager() IConnmanager
	// 注册方法
	SetOnConnStart(func(IConnection))
	SetOnConnStop(func(IConnection))
	// 调用方法
	CallOnConnStart(IConnection)
	CallOnConnStop(IConnection)
}
