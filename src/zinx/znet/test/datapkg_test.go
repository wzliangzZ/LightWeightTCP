package test

import (
	"fmt"
	"io"
	"net"
	"os"
	"testing"
	"zx/src/zinx/znet"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
// 测试前 把 Global_obj.Reload() 注释掉
func TestDataPack(t *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:7777")
	check(err)

	go func() {
		for {
			conn, err := lis.Accept()
			check(err)

			go func(conn net.Conn) {
				dp := znet.NewDataPkg()
				for {
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil{
						fmt.Println("read head error")
						break
					}

					msgHead, err := dp.Unpack(headData)
					check(err)

					if msgHead.GetMsgLen() > 0{
						fmt.Println("succ ", msgHead.GetMsgLen())
						msg := msgHead.(*znet.Message)
						msg.Data = make([]byte, msg.GetMsgLen())
						_, err := io.ReadFull(conn, msg.Data)
						check(err)


						fmt.Println("succ id:", msg.GetMsgId(), "len:", msg.GetMsgLen(), "data:", string(msg.GetData()))
					}

				}

			}(conn)

		}
	}()


	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	check(err)

	dp := znet.NewDataPkg()
	msg1 := znet.Message{
		Id: 1,
		DataLen: 5,
		Data: []byte("12345"),
	}
	bin1, err := dp.Pack(&msg1)
	check(err)

	msg2 := znet.Message{
		Id: 2,
		DataLen: 7,
		Data: []byte("7654321"),
	}
	bin2, err := dp.Pack(&msg2)
	check(err)
	bin := append(bin1, bin2...)

	conn.Write(bin)

	select{}
}
