package ziface

// 消息管理抽象层

type IMsgHandle interface {
	// 调度对应的router消息处理
	DoMsgHandler(IRequest)
	// 添加处理逻辑
	AddRouter(uint32, IRouter)
	// 启动worker工作池（只执行一次）
	StartWorkerPool()
	// 将消息发给 消息队列
	SendMsgToTaskQueue(IRequest)
}
