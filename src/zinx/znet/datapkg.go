package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zx/src/zinx/utils"
	"zx/src/zinx/ziface"
)

type DataPkg struct{}

func NewDataPkg() *DataPkg {
	return &DataPkg{}
}

func (dp *DataPkg) GetHeadLen() uint32 {
	// DataLen uint32(4字节) + ID uint32(4字节)
	return 8
}
func (dp *DataPkg) Pack(msg ziface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})

	// datalen 和 id 写入databuff
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	// 数据写入databuff
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

func (dp *DataPkg) Unpack(bd []byte) (ziface.IMessage, error) {
	br := bytes.NewReader(bd)

	msg := &Message{}
	if err := binary.Read(br, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(br, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	// 判断datalen是否超过允许最大
	if utils.Global_obj.MaxPackageSize > 0 && msg.DataLen > utils.Global_obj.MaxPackageSize {
		return nil, errors.New("too large msg data recv")
	}

	// 读数据
	if err := binary.Read(br, binary.LittleEndian, &msg.Data); err != nil {
		return nil, err
	}

	return msg, nil
}
