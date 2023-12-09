package main

import (
	"fmt"
	"zx/src/zinx/ziface"
	"zx/src/zinx/znet"
)

type PingRouter struct{
	znet.BaseRouter

}

// 业务主方法
func (br *PingRouter)  Handle(req ziface.IRequest) {
	fmt.Println("Handle...")
	fmt.Println("msgid:", req.GetMsgId(), "DATA=", string(req.GetData()))

	err := req.GetConnetcion().SendMsg(1, []byte("ping ping ping"))
	if err != nil {
		fmt.Println(err)
	}
}



func main() {
	// 创建
	s := znet.NewServer("v0_5")
	// 添加自定义路由
	s.AddRouter(&PingRouter{})
	// 运行
	s.Serve()
}