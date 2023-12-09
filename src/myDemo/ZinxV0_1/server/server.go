package main

import (
	"zx/src/zinx/znet"
)

func main() {
	// 创建
	s := znet.NewServer("v0_1")
	// 运行
	s.Serve()
}