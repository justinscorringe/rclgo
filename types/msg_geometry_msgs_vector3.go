package types

import (
	"unsafe"

	cwrap "github.com/justinscorringe/rclgo/ros2"
)

type GeometryMsgsVector3 struct {
	data *cwrap.GeometryMsgs_MsgVector3
	MessageBase
}

func (msg *GeometryMsgsVector3) Type() MessageTypeSupport {
	return msg.MsgType
}

func (msg *GeometryMsgsVector3) Data() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *GeometryMsgsVector3) GetDataMap() interface{} {
	return msg.data
}

func (msg *GeometryMsgsVector3) InitMessage() {
	msg.data = cwrap.InitGeometryMsgsMsgVector3()
	msg.MsgType = GetMessageTypeFromGeometryMsgsMsgVector3()
}

func (msg *GeometryMsgsVector3) DestroyMessage() {
	cwrap.DestroyGeometryMsgsMsgVector3(msg.data)
}

func (msg *GeometryMsgsVector3) Set(data map[string]interface{}) {
	msg.data.Set(data)
}

func GetMessageTypeFromGeometryMsgsMsgVector3() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromGeometryMsgsMsgVector3()
	return MessageTypeSupport{ret}
}
