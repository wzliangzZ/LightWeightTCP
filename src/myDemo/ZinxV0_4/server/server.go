package main

import (
	"fmt"
	"zx/src/zinx/ziface"
	"zx/src/zinx/znet"
)

type PingRouter struct{
	znet.BaseRouter

}
// 业务前方法
func (pr *PingRouter) PreHandle(req ziface.IRequest) {
	fmt.Println("PreHandle")
	_, err := req.GetConnetcion().GetTCPConnection().Write([]byte("before ping..."))
	if err != nil{
		fmt.Println("call before ping err")
	}
}
// 业务主方法
func (br *PingRouter)  Handle(req ziface.IRequest) {
	fmt.Println("Handle")
	_, err := req.GetConnetcion().GetTCPConnection().Write([]byte("ping..."))
	if err != nil{
		fmt.Println("call ping err")
	}
}
// 业务后方法
func (br *PingRouter)  PostHandle(req ziface.IRequest) {
	fmt.Println("PostHandle")
	_, err := req.GetConnetcion().GetTCPConnection().Write([]byte("PostHandle ping..."))
	if err != nil{
		fmt.Println("call PostHandle ping err")
	}
}


func main() {
	// 创建
	s := znet.NewServer("v0_4")
	// 添加自定义路由
	s.AddRouter(&PingRouter{})
	// 运行
	s.Serve()
}