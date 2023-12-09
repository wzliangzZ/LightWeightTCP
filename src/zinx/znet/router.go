package znet


import(
	"zx/src/zinx/ziface"
)


// 初衷：先嵌入改BaseRouter基类，再根据需求对方法重写
// 不做具体实现，若实现接口，则必须实现3个方法，继承该结构体只需按需求重写方法
type BaseRouter struct {}

// 业务前方法
func (br *BaseRouter) PreHandle(req ziface.IRequest) {}
// 业务主方法
func (br *BaseRouter)  Handle(req ziface.IRequest) {}
// 业务后方法
func (br *BaseRouter)  PostHandle(req ziface.IRequest) {}
