package types

import (
	"unsafe"

	cwrap "github.com/justinscorringe/rclgo/ros2"
)

type GeometryMsgsTwist struct {
	data *cwrap.GeometryMsgs_MsgTwist
	MessageBase
}

func (msg *GeometryMsgsTwist) Type() MessageTypeSupport {
	return msg.MsgType
}

func (msg *GeometryMsgsTwist) Data() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *GeometryMsgsTwist) GetDataMap() interface{} {
	return msg.data
}

func (msg *GeometryMsgsTwist) InitMessage() {
	msg.data = cwrap.InitGeometryMsgsMsgTwist()
	msg.MsgType = GetMessageTypeFromGeometryMsgsMsgTwist()
}

func (msg *GeometryMsgsTwist) DestroyMessage() {
	cwrap.DestroyGeometryMsgsMsgTwist(msg.data)
}

func (msg *GeometryMsgsTwist) Set(data map[string]interface{}) {
	msg.data.Set(data)
}

func GetMessageTypeFromGeometryMsgsMsgTwist() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromGeometryMsgsMsgTwist()
	return MessageTypeSupport{ret}
}
