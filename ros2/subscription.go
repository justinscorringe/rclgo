package ros2

// #cgo CFLAGS: -I/opt/ros/dashing/include
// #include <rcl/rcl.h>
import "C"

import (
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

//
func (s *Subscription) TakeMessageRaw(msgType MessageType) (Message, error) {
	if msgType == nil || s.rclSubscription == nil {
		return nil, NewErr("nil", Error)
	}

	msgRawBytes := SerializedMsg{}

	ret := C.rcl_take_serialized_message(
		(*C.rcl_subscription_t)(s.rclSubscription),
		(*C.rcl_serialized_message_t)(&msgRawBytes),
		(*C.rmw_message_info_t)(msgType.RosInfo()),
		nil)

	if ret != Ok {
		return nil, NewErr("RclTake", int(ret))
	}

	msg := msgType.NewMessage()

	return msg, nil
}

//
func (s *Subscription) TakeMessage(msg Message) error {
	if msg == nil || s.rclSubscription == nil {
		return NewErr("nil", Error)
	}
	var ret = C.rcl_take(
		(*C.rcl_subscription_t)(s.rclSubscription),
		msg.RosData(),
		(*C.rmw_message_info_t)(msg.Type().RosInfo()),
		nil,
	)

	if ret != Ok {
		return NewErr("RclTake", int(ret))
	}

	return nil
}
