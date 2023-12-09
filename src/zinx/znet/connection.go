package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"zx/src/zinx/utils"
	"zx/src/zinx/ziface"
)

// 连接模块
type Connection struct {
	// 哪个server创建的conn
	TcpServer ziface.IServer
	// 当前连接TPC套接字
	Conn *net.TCPConn

	// 套接字ID
	ConnID uint32

	// 当前连接状态
	isClosed bool

	// 退出通知chan
	ExitChan chan bool

	// 传递IO消息的管道
	msgChan chan []byte

	// 管理msgid 对于的 api
	MsgHandler ziface.IMsgHandle

	// 连接属性集合
	property map[string]any

	// 保护属性的锁
	propertyLock sync.Mutex
}

// 初始化
func NewConnection(is ziface.IServer, conn *net.TCPConn, id uint32, mh ziface.IMsgHandle) *Connection {
	c := &Connection{
		TcpServer:  is,
		Conn:       conn,
		ConnID:     id,
		MsgHandler: mh,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte),
		property:   make(map[string]any),
	}
	// 将conn加入到conmanager中
	c.TcpServer.GetConnManager().Add(c)

	return c
}

// 读业务方法
func (c *Connection) StartReader() {
	fmt.Println("[Reader go is running...]")
	defer fmt.Println(c.ConnID, "[reader] is exit")
	defer c.Stop()

	for {
		// 创建拆包对象
		dp := NewDataPkg()
		headData := make([]byte, dp.GetHeadLen())
		// 得到 msg head 8字节
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("readfull err", err)

			break
		}

		// 拆包，得到msgID 和 msgDataLen,封装好msg

		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("msg unpack err", err)
		}
		// 根据len, 再次读取data
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("2 readFull err", err)
				break
			}
		}
		msg.SetData(data)
		req := Request{
			Conn: c,
			msg:  msg,
		}
		// 已开启工作池
		if utils.Global_obj.WorkerPollSize > 0 {
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			// 根据绑好的msgid执行api业务
			go c.MsgHandler.DoMsgHandler(&req)
		}

	}
}

// 写消息，专门发给客户端消息的模块
func (c *Connection) StartWriter() {
	fmt.Println("[wirter go is running...]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn writer exit]")

	// 等待chan消息
	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("start writer conn err", err)
			}
		case <-c.ExitChan:
			return
		}

	}
}

// sendMsg() 将数据封包，再发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("conn closeed")
	}
	// len|id|data
	bin, err := NewDataPkg().Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("pack err", err)
		return errors.New("pack err")
	}

	c.msgChan <- bin

	return nil
}

// 启动连接
func (c *Connection) Start() {
	fmt.Println("conn start...", c.ConnID)

	// 启动当前连接的读数据的业务
	go c.StartReader()

	// 启动当前连接的写数据的业务
	go c.StartWriter()
	// 调用hook方法
	c.TcpServer.CallOnConnStart(c)
}

// 停止连接
func (c *Connection) Stop() {
	fmt.Println("conn stop id:", c.ConnID)
	if c.isClosed {
		return
	}
	// 调用hook方法
	c.TcpServer.CallOnConnStop(c)
	c.isClosed = true
	c.Conn.Close()

	// 告知writer关闭
	c.ExitChan <- true
	// 从管理器中删除当前conn
	c.TcpServer.GetConnManager().Remove(c)

	close(c.ExitChan)
	close(c.msgChan)

}

// 获取scoket coon
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取客户端TCP IP，port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据
func (c *Connection) Send([]byte) error {
	return nil
}

// 设置连接属性
func (c *Connection) SetProperty(key string, val any) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	c.property[key] = val
}

// 获取连接属性
func (c *Connection) GetProperty(key string) (any, error) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	if val, ok := c.property[key]; ok {
		return val, nil
	}
	return nil, errors.New("no such property")
}

// 移除连接属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}
