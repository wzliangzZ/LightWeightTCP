package znet

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		Id: id,
		Data: data,
		DataLen: uint32(len(data)),
	}
}

func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}
func (msg *Message) GetMsgLen() uint32 {
	return msg.DataLen
}
func (msg *Message) GetData() []byte {
	return msg.Data
}

func (msg *Message) SetMsgId(id uint32) {
	msg.Id = id
}
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}
