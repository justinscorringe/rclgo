package ros2

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
// #include <rmw/rmw.h>
// #include <rmw/error_handling.h>
// #include <rmw/validate_full_topic_name.h>
import "C"

import (
	"bytes"
	"fmt"
	"unsafe"
)

//
type RclSubscription C.rcl_subscription_t

//
type RclSubscriptionOptions C.rcl_subscription_options_t

//
type Subscription struct {
	rclSubscription *RclSubscription
}

//
type SubscriptionOptions struct {
	rclSubscriptionOptions *RclSubscriptionOptions
}

func NewZeroInitializedSubscription() Subscription {
	var ret C.rcl_subscription_t = C.rcl_get_zero_initialized_subscription()
	zeroSubscription := RclSubscription(ret)
	return Subscription{&zeroSubscription}
}

func NewSubscriptionDefaultOptions() SubscriptionOptions {
	var ret C.rcl_subscription_options_t = C.rcl_subscription_get_default_options()
	defOpts := RclSubscriptionOptions(ret)
	return SubscriptionOptions{&defOpts}
}

//
func (s *Subscription) Init(subscriptionOptions SubscriptionOptions, node Node, topicName string, msg MessageType) error {

	cTopicName := C.CString(topicName)
	defer C.free(unsafe.Pointer(cTopicName))

	var ret C.int32_t = C.rcl_subscription_init(
		(*C.rcl_subscription_t)(s.rclSubscription),
		(*C.struct_rcl_node_t)(node.rclNode),
		(*C.rosidl_message_type_support_t)(msg.RosType()),
		cTopicName,
		(*C.rcl_subscription_options_t)(subscriptionOptions.rclSubscriptionOptions),
	)

	//,

	if ret != Ok {
		return NewErr("RclSubscriptionInit", int(ret))
	}

	return nil
}

//
func (s *Subscription) SubscriptionFini(node Node) error {

	var ret C.int32_t = C.rcl_subscription_fini(
		(*C.rcl_subscription_t)(s.rclSubscription),
		(*C.rcl_node_t)(node.rclNode),
	)

	if ret != Ok {
		return NewErr("RclSubscriptionFini", int(ret))
	}

	return nil
}

func NewRawMessage() SerializedMsg {
	ret := C.rmw_serialized_message_t{}
	ret.allocator = C.rcutils_get_default_allocator()
	return SerializedMsg(ret)
}

//
func (s *Subscription) TakeRawMessage(msgType MessageType) (Message, error) {
	if msgType == nil || s.rclSubscription == nil {
		return nil, NewErr("nil", Error)
	}

	msgEvent := NewRawMessage()

	ret := C.rcl_take_serialized_message(
		(*C.rcl_subscription_t)(s.rclSubscription),
		(*C.rmw_serialized_message_t)(&msgEvent),
		(*C.rmw_message_info_t)(msgType.RosInfo()),
		nil)

	if ret != Ok {
		return nil, NewErr("RclTake", int(ret))
	}

	msg := msgType.NewMessage()

	buf := (*[]byte)(unsafe.Pointer(&msgEvent.buffer))
	fmt.Printf("length of buffer: %v    buffer: %s    msgEvent: %s\n", msgEvent.buffer_length, buf, msgEvent)
	reader := bytes.NewReader(*buf)
	err := msg.Deserialize(reader, int(msgEvent.buffer_length))
	if err != nil {
		return nil, err
	}

	return msg, nil
}

//
func (s *Subscription) TakeMessage(msg Message) error {
	if s.rclSubscription == nil {
		return NewErr("nil", Error)
	}

	rosMessage := msg.RosMessage()

	var ret = C.rcl_take(
		(*C.rcl_subscription_t)(s.rclSubscription),
		rosMessage,
		//(*C.rmw_message_info_t)(msg.Type().RosInfo()),
		nil,
		nil,
	)

	if ret != Ok {
		return NewErr("RclTake", int(ret))
	}

	return nil
}
