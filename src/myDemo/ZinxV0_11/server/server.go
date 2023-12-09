package main

import (
	"fmt"
	"zx/src/zinx/ziface"
	"zx/src/zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

// 业务主方法
func (br *PingRouter) Handle(req ziface.IRequest) {
	fmt.Println("Handle...")
	fmt.Println("msgid:", req.GetMsgId(), "DATA=", string(req.GetData()))

	err := req.GetConnetcion().SendMsg(0, []byte("ping ping ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type HRouter struct {
	znet.BaseRouter
}

// 业务主方法
func (h *HRouter) Handle(req ziface.IRequest) {
	fmt.Println("Handle...")
	fmt.Println("msgid:", req.GetMsgId(), "DATA=", string(req.GetData()))

	err := req.GetConnetcion().SendMsg(1, []byte("hhhhh hhhhh hhhhhh"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// 创建一个服务器实例
	s := znet.NewServer("v0_9")

	s.AddRouter(0, &PingRouter{})  // 添加PingRouter路由，路由ID为0
	s.AddRouter(1, &HRouter{})    // 添加HRouter路由，路由ID为1

	// 注册连接开始的hook函数
	s.SetOnConnStart(func(i ziface.IConnection) {
		fmt.Println("-->doconn begin is called...")

		// 向连接发送消息
		if err := i.SendMsg(10, []byte("doconn begin")); err != nil {
			fmt.Println(err)
		}

		fmt.Println("set conn property")
		// 设置连接属性
		i.SetProperty("name", "jack")
		i.SetProperty("age", 18)
	})

	// 注册连接断开的hook函数
	s.SetOnConnStop(func(i ziface.IConnection) {
		fmt.Println("-->doconn stop is called...")
		fmt.Println("conn id:", i.GetConnID(), " is lost...")

		// 获取连接属性
		if name, err := i.GetProperty("name"); err == nil {
			fmt.Println("name=", name)
		}
		if age, err := i.GetProperty("age"); err == nil {
			fmt.Println("age=", age)
		}
	})

	// 启动服务器
	s.Serve()
}
