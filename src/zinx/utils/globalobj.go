package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"zx/src/zinx/ziface"
)

// 存储zinx全局参数，可以通过zinx.json由用户配置
type GlobalObj struct {
	// Server
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string

	// Zinx
	Vsersion       string
	MaxConn        int
	MaxPackageSize uint32 //当前框架数据包的最大值

	// 当前工作池最多数量
	WorkerPollSize uint32
	// 每个worker的最大任务书
	MaxWorkerPollLen uint32
}

// 定义一个全局对象对外访问
var Global_obj *GlobalObj

// 加载用户自定义配置文件
func (g *GlobalObj) Reload() {
	data, err := os.ReadFile(`F:/GoPro/zinx/src/myDemo/ZinxV0_4/conf/zinx.json`)
	if err != nil {
		fmt.Println("os.ReadFile err ", err)
		panic(err)
	}
	// 将json文件解析到struct中
	err = json.Unmarshal(data, &Global_obj)
	if err != nil {
		panic(err)
	}

}

func init() {
	// 默认配置
	Global_obj = &GlobalObj{
		Name:             "ZinxServerApp",
		Vsersion:         "V0_9",
		TcpPort:          8999,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		MaxWorkerPollLen: 1024,
		WorkerPollSize:   10,
	}

	//加载用户配置，更换默认配置
	Global_obj.Reload()
}
