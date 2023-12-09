package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zx/src/zinx/utils"
	"zx/src/zinx/znet"

)

func main() {
	fmt.Println("client start...")
	// 连接服务器，得到一个conn
	time.Sleep(time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")

	if err != nil {
		fmt.Println("client err:", err)
		return
	}

	for {
		// 连接调用Write,写数据
		dp := znet.NewDataPkg()
		bins, err := dp.Pack(znet.NewMessage(0, []byte("zinxv0.5 client test messgage")))
		if err != nil {
			fmt.Println("client pack err", err)
			return 
		}
		_, err = conn.Write(bins)
		if err != nil {
			fmt.Println("conn.write err", err)
			return 
		}
		// 服务器回显
		binHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binHead); err != nil {
			fmt.Println("readfull err", err)
		}
		msgh, err := dp.Unpack(binHead)

		if err != nil {
			fmt.Println("unpack err", err)
		}
		
		if msgh.GetMsgLen() > 0 && 	msgh.GetMsgLen() < utils.Global_obj.MaxPackageSize {
			msg := msgh.(*znet.Message)
			msg.Data = make([]byte, msg.DataLen)
			if _, err := io.ReadFull(conn, msg.Data); err != nil{
				fmt.Println("read msg data err", err)
				return
			}

			fmt.Println("succ", msg.Id," ", msg.GetMsgLen() ," ", string(msg.GetData()))
		}

		// 阻塞
		time.Sleep(time.Second)
	}

}
