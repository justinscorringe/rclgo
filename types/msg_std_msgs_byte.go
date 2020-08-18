package types

import (
	"unsafe"

	cwrap "github.com/justinscorringe/rclgo/ros2"
)

type StdMsgsByte struct {
	data    *cwrap.StdMsgs_MsgByte
	MsgType MessageTypeSupport
}

func (msg *StdMsgsByte) Type() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsByte) Data() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsByte) InitMessage() {
	msg.data = cwrap.InitStdMsgsMsgByte()
	msg.MsgType = GetMessageTypeFromStdMsgsByte()
}

func (msg *StdMsgsByte) SetByte(data byte) {
	//TODO: to implement the setter
	msg.data.Set(data)
}

func (msg *StdMsgsByte) GetByte() byte {
	return byte(msg.data.Data())
}

func (msg *StdMsgsByte) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue := ""
	return myRetValue
}

func (msg *StdMsgsByte) DestroyMessage() {
	cwrap.DestroyStdMsgsMsgByte(msg.data)
}

func GetMessageTypeFromStdMsgsByte() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromStdMsgsMsgByte()
	return MessageTypeSupport{ret}
}
