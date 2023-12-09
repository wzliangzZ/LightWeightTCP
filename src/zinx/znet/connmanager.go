package znet

import (
	"fmt"
	"sync"
	"zx/src/zinx/ziface"
	"errors"
)

type ConnManager struct {
	// 管理的连接
	connections map[uint32]ziface.IConnection
	// 读写锁
	connLock sync.RWMutex
}


func NewConnManager() *ConnManager{
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}


func (cm *ConnManager) Add (conn ziface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	cm.connections[conn.GetConnID()] = conn
	fmt.Println("conn add succ", conn.GetConnID())
}
func (cm *ConnManager) Remove(conn ziface.IConnection){
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	if _, ok := cm.connections[conn.GetConnID()]; ok {
		delete(cm.connections, conn.GetConnID())
		fmt.Println("succ remove:", conn.GetConnID(), "connmanager len=:", cm.Len())
	}

}
func (cm *ConnManager) Get(id uint32) (ziface.IConnection, error) {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()
	if f, ok := cm.connections[id]; ok {
		fmt.Println("succ get:", id)
		return f, nil
	}
	return nil, errors.New("no such conn id")
}
func (cm *ConnManager) Len() int {
	return len(cm.connections)
}
func (cm *ConnManager) ClearConn() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	//删除并停止 conn
	for id, conn := range cm.connections{
		// 停止
		conn.Stop()
		// 删除
		delete(cm.connections, id)
	}
	fmt.Println("succ clearConn")
}