package ziface


// 路由接口，数据都是IRequest
type IRouter interface {

	// 业务前方法
	PreHandle(req IRequest)
	// 业务主方法
	Handle(req IRequest)
	// 业务后方法
	PostHandle(req IRequest)
}