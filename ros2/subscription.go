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
func RclGetZeroInitializedSubscription() RclSubscription {
	var ret C.rcl_subscription_t = C.rcl_get_zero_initialized_subscription()
	return RclSubscription(ret)
}

//
func RclSubscriptionInit(
	subscription *RclSubscription,
	node *RclNode,
	typeSupport *ROSIdlMessageTypeSupport,
	topicName string,
	options *RclSubscriptionOptions,
) int {

	cTopicName := C.CString(topicName)
	defer C.free(unsafe.Pointer(cTopicName))

	var ret C.int32_t = C.rcl_subscription_init(
		(*C.rcl_subscription_t)(subscription),
		(*C.struct_rcl_node_t)(node),
		(*C.rosidl_message_type_support_t)(typeSupport),
		cTopicName,
		(*C.rcl_subscription_options_t)(options),
	)

	return int(ret)
}

//
func RclSubscriptionFini(subscription *RclSubscription, node *RclNode) int {
	var ret C.int32_t = C.rcl_subscription_fini(
		(*C.rcl_subscription_t)(subscription),
		(*C.rcl_node_t)(node),
	)
	return int(ret)
}

//
func RclSubscriptionGetDefaultOptions() RclSubscriptionOptions {
	var ret C.rcl_subscription_options_t = C.rcl_subscription_get_default_options()
	return RclSubscriptionOptions(ret)
}

type Subscription struct {
	rclSubscription *RclSubscription
}

type SubscriptionOptions struct {
	rclSubscriptionOptions *RclSubscriptionOptions
}

func NewZeroInitializedSubscription() Subscription {
	zeroSubscription := RclGetZeroInitializedSubscription()
	return Subscription{&zeroSubscription}
}

func NewSubscriptionDefaultOptions() SubscriptionOptions {
	defOpts := RclSubscriptionGetDefaultOptions()
	return SubscriptionOptions{&defOpts}
}

func (s *Subscription) Init(
	subscriptionOptions SubscriptionOptions,
	node Node,
	topicName string,
	msg MessageType,
) error {

	ret := RclSubscriptionInit(
		s.rclSubscription,
		node.rclNode,
		msg.RosType(),
		topicName,
		subscriptionOptions.rclSubscriptionOptions,
	)

	if ret != Ok {
		return NewErr("RclSubscriptionInit", ret)
	}

	return nil
}

func (s *Subscription) SubscriptionFini(node Node) error {
	ret := RclSubscriptionFini(
		s.rclSubscription,
		node.rclNode,
	)

	if ret != Ok {
		return NewErr("RclSubscriptionFini", ret)
	}

	return nil
}

//
func (s *Subscription) TakeMessageSerialized(msgType MessageType) (Message, error) {
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
