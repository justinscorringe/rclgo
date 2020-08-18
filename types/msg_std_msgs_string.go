package types

import (
	"unsafe"

	cwrap "github.com/justinscorringe/rclgo/ros2"
)

type StdMsgsString struct {
	data *cwrap.StdMsgs_MsgString
	MessageBase
}

func (msg *StdMsgsString) Type() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsString) Data() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsString) GetDataAsString() string {
	return msg.data.String()
}

func (msg *StdMsgsString) InitMessage() {
	msg.data = cwrap.InitStdMsgsMsgString()
	msg.MsgType = GetMessageTypeFromStdMsgsString()
}

func (msg *StdMsgsString) DestroyMessage() {
	cwrap.DestroyStdMsgsMsgString(msg.data)
}

func (msg *StdMsgsString) SetText(text string) {
	msg.data.Set(text)
}

func GetMessageTypeFromStdMsgsString() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromStdMsgsMsgString()
	return MessageTypeSupport{ret}
}
