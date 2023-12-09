package znet

import (
	"fmt"
	"zx/src/zinx/utils"
	"zx/src/zinx/ziface"
)

type MsgHandle struct {
	// msgid 对应的处理方法
	Apis map[uint32]ziface.IRouter

	// 消息队列
	TaskQueue []chan ziface.IRequest

	// 工作池数量
	WorkerPoolSize uint32

}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.Global_obj.WorkerPollSize,
		TaskQueue: make([]chan ziface.IRequest, utils.Global_obj.WorkerPollSize),
	}
}

// 调度对应的router消息处理
func (mh *MsgHandle) DoMsgHandler(req ziface.IRequest) {
	f, ok := mh.Apis[req.GetMsgId()]
	if !ok {
		fmt.Println("no such id:",req.GetMsgId())
	}
	f.PreHandle(req)
	f.Handle(req)
	f.PostHandle(req)
	fmt.Println("succ doMsgHandler")
}


// 添加处理逻辑
func (mh *MsgHandle) AddRouter(id uint32,router ziface.IRouter) {
	mh.Apis[id] = router
	fmt.Println("succ add router")
}

// 启动worker工作池（只执行一次）
func (mh *MsgHandle) StartWorkerPool(){
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		
		// 给当前chan消息队列开辟空间
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.Global_obj.MaxWorkerPollLen)
		
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}


// 启动worker
func (mh *MsgHandle) StartOneWorker(id int, q chan ziface.IRequest) {
	fmt.Println("worker id = ", id, "is start")

	for req := range q{
		mh.DoMsgHandler(req)
	}

	// 推荐使用上面方法

	// for {
	// 	select{
	// 	case req := <-q:
	// 		mh.DoMsgHandler(req)
	// 	}
	// }
	
}
// 将消息发给 消息队列
func (mh *MsgHandle) SendMsgToTaskQueue(req ziface.IRequest) {
	wid := req.GetConnetcion().GetConnID() % mh.WorkerPoolSize
	fmt.Println("msd id:",req.GetMsgId(), " to wid:", wid)
	mh.TaskQueue[wid] <- req
}