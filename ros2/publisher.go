package ros2

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include "rcl/rcl.h"
import "C"
import "unsafe"

//
type RclPublisher C.rcl_publisher_t

//
type RclPublisherOptions C.rcl_publisher_options_t

//
type Publisher struct {
	rclPublisher *RclPublisher
}

//
type PublisherOptions struct {
	rclPublisherOptions *RclPublisherOptions
}

func NewZeroInitializedPublisher() Publisher {
	ret := C.rcl_get_zero_initialized_publisher()
	zeroPublisher := RclPublisher(ret)
	return Publisher{&zeroPublisher}
}

func NewPublisherDefaultOptions() PublisherOptions {
	ret := C.rcl_publisher_get_default_options()
	defOpts := RclPublisherOptions(ret)
	return PublisherOptions{&defOpts}
}

func (p *Publisher) GetTopicName() string {
	ret := C.rcl_publisher_get_topic_name((*C.rcl_publisher_t)(p.rclPublisher))
	return C.GoString(ret)
}

func (p *Publisher) Init(publisherOptions PublisherOptions, node Node, topicName string, msg MessageType) error {
	tName := C.CString(topicName)
	defer C.free(unsafe.Pointer(tName))

	ret := C.rcl_publisher_init(
		(*C.rcl_publisher_t)(p.rclPublisher),
		(*C.rcl_node_t)(node.rclNode),
		(*C.rosidl_message_type_support_t)(msg.RosType()),
		tName,
		(*C.rcl_publisher_options_t)(publisherOptions.rclPublisherOptions),
	)

	if ret != Ok {
		return NewErr("RclInitOptionsInit", int(ret))
	}

	return nil
}

func (p *Publisher) PublisherFini(node Node) error {

	ret := C.rcl_publisher_fini(
		(*C.rcl_publisher_t)(p.rclPublisher),
		(*C.rcl_node_t)(node.rclNode),
	)
	if ret != Ok {
		return NewErr("RclPublisherFini", int(ret))
	}

	return nil
}

func (p *Publisher) Publish(msg Message) error {

	//(*C.rosidl_message_type_support_t)(typeSupport).data = data

	rosMessage := msg.RosMessage()

	ret := C.rcl_publish(
		(*C.rcl_publisher_t)(p.rclPublisher),
		unsafe.Pointer(&rosMessage),
		nil,
	)

	if ret != Ok {
		return NewErr("RclPublish", int(ret))
	}

	return nil
}

func (p *Publisher) PublishRaw(msg Message) error {

	//(*C.rosidl_message_type_support_t)(typeSupport).data = data

	// ret := C.rcl_publish_serialized_message(
	// 	(*C.rcl_publisher_t)(p.rclPublisher),
	// 	msg.RosMessage(),
	// 	nil,
	// )

	// if ret != Ok {
	// 	return NewErr("RclPublish", int(ret))
	// }

	return nil
}

func (p *Publisher) IsValid() bool {
	var ret C.bool = C.rcl_publisher_is_valid((*C.rcl_publisher_t)(p.rclPublisher))
	return bool(ret)
}
