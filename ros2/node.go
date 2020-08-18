package ros2

// #cgo CFLAGS: -I/opt/ros/dashing/include
// #include <rcl/rcl.h>
import "C"
import "unsafe"

//
type RclNode C.rcl_node_t

//
type RclNodeOptions C.rcl_node_options_t

//
type RmwNode C.rmw_node_t

//
type RclGuardCondition C.rcl_guard_condition_t

// Node is a structure that encapsulates a ROS Node.
type Node struct {
	rclNode *RclNode
}

// NodeOptions is a structure that encapsulates the options for creating an
// RclNode.
type NodeOptions struct {
	rclNodeOptions *RclNodeOptions
}

// NewZeroInitializedNode returns an RclNode with members initialized to `NULL`.
func NewZeroInitializedNode() Node {
	zeroNode := RclGetZeroInitializedNode()
	return Node{rclNode: &zeroNode}
}

// NewNodeDefaultOptions returns the default node options in a RclNodeOptions.
func NewNodeDefaultOptions() NodeOptions {
	defOpts := RclNodeGetDefaultOptions()
	return NodeOptions{rclNodeOptions: &defOpts}
}

// Init initialize a ROS node.
func (n *Node) Init(name string, namespace string, ctx Context, nodeOptions NodeOptions) error {
	ret := RclNodeInit(
		n.rclNode,
		name,
		namespace,
		ctx.rclContext,
		nodeOptions.rclNodeOptions,
	)
	if ret != Ok {
		return NewErr("RclNodeInit", ret)
	}

	return nil
}

// Fini finalizes an RclNode.
func (n *Node) Fini() error {
	ret := RclNodeFini(n.rclNode)
	if ret != Ok {
		return NewErr("RclNodeFini", ret)
	}

	return nil
}

// IsValid returns `true` if the node is valid, else `false`.
func (n *Node) IsValid() bool {
	ret := RclNodeIsValid(n.rclNode)
	return ret
}

// IsValidExceptContext returns true if node is valid, except for the context
// being valid.
func (n *Node) IsValidExceptContext() bool {
	ret := RclNodeIsValidExceptContext(n.rclNode)
	return ret
}

// GetName returns the name of the node.
func (n *Node) GetName() string {
	ret := RclNodeGetName(n.rclNode)
	return ret
}

// GetNamespace returns the namespace of the node.
func (n *Node) GetNamespace() string {
	ret := RclNodeGetNamespace(n.rclNode)
	return ret
}

// GetFullyQualifiedName returns the fully qualified name of the node.
func (n *Node) GetFullyQualifiedName() string {
	ret := RclNodeGetFullyQualifiedName(n.rclNode)
	return ret
}

// GetOptions returns the rcl node options.
func (n *Node) GetOptions() NodeOptions {
	opts := RclNodeGetOptions(n.rclNode)
	return NodeOptions{opts}
}

// GetDomainID returns the ROS domain ID that the node is using.
func (n *Node) GetDomainID() (uint, error) {
	var domainID uint
	ret := RclNodeGetDomainID(n.rclNode, &domainID)
	if ret != Ok {
		return domainID, NewErr("RclNodeGetDomainID", ret)
	}

	return domainID, nil
}

// AssertLiveliness manually asserts that this node is alive
// (for RMW_QOS_POLICY_LIVELINESS_MANUAL_BY_NODE)
func (n *Node) AssertLiveliness() error {
	ret := RclNodeAssertLiveliness(n.rclNode)
	if ret != Ok {
		return NewErr("RclNodeAssertLiveliness", ret)
	}

	return nil
}

// GetRmwHandle returns the rmw node handle.
// func (n *Node) GetRmwHandle() RmwHandle {
// 	ret := RclNodeGetRmwHandle(n.rclNode)
// 	return RmwHandle{ret}
// }

// GetRclInstanceID returns the associated rcl instance id.
func (n *Node) GetRclInstanceID() uint64 {
	var ret uint64 = RclNodeGetRclInstanceID(n.rclNode)
	return ret
}

// GetGraphGuardCondition returns a guard condition which is triggered when the
// ROS graph changes.
// func (n *Node) GetGraphGuardCondition() GuardCondition {
// 	guard := RclNodeGetGraphGuardCondition(n.rclNode)
// 	return GuardCondition{guard}
// }

// GetLoggerName returns the logger name of the node.
func (n *Node) GetLoggerName() string {
	var ret string = RclNodeGetLoggerName(n.rclNode)
	return ret
}

//
func RclNodeGetDefaultOptions() RclNodeOptions {
	var defOpts C.rcl_node_options_t = C.rcl_node_get_default_options()
	return RclNodeOptions(defOpts)
}

//
func RclGetZeroInitializedNode() RclNode {
	var zeroNode C.rcl_node_t = C.rcl_get_zero_initialized_node()
	return RclNode(zeroNode)
}

//
func RclNodeInit(node *RclNode, name string, namespace string, ctx RclContextPtr, options *RclNodeOptions) int {
	var cName *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var cNamespace *C.char = C.CString(namespace)
	defer C.free(unsafe.Pointer(cNamespace))

	var ret C.int32_t = C.rcl_node_init(
		(*C.rcl_node_t)(node),
		cName,
		cNamespace,
		(*C.rcl_context_t)(ctx),
		(*C.rcl_node_options_t)(options),
	)

	return int(ret)
}

//
func RclNodeFini(node *RclNode) int {
	cNode := (*C.rcl_node_t)(node)
	ret := C.rcl_node_fini(cNode)
	return int(ret)
}

//
func RclNodeIsValid(node *RclNode) bool {
	var ret C.bool = C.rcl_node_is_valid(
		(*C.rcl_node_t)(node),
	)
	return bool(ret)
}

//
func RclNodeIsValidExceptContext(node *RclNode) bool {
	var ret C.bool = C.rcl_node_is_valid_except_context(
		(*C.rcl_node_t)(node),
	)
	return bool(ret)
}

//
func RclNodeGetName(node *RclNode) string {
	var cStr *C.char = C.rcl_node_get_name(
		(*C.rcl_node_t)(node),
	)
	return C.GoString(cStr)
}

//
func RclNodeGetNamespace(node *RclNode) string {
	var cStr = C.rcl_node_get_namespace(
		(*C.rcl_node_t)(node),
	)
	return C.GoString(cStr)
}

//
func RclNodeGetFullyQualifiedName(node *RclNode) string {
	var cStr = C.rcl_node_get_fully_qualified_name(
		(*C.rcl_node_t)(node),
	)
	return C.GoString(cStr)
}

//
func RclNodeGetOptions(node *RclNode) *RclNodeOptions {
	var opts *C.rcl_node_options_t = C.rcl_node_get_options(
		(*C.rcl_node_t)(node),
	)
	return (*RclNodeOptions)(opts)
}

//
func RclNodeGetDomainID(node *RclNode, domainID *uint) int {
	var dom C.ulong
	var ret C.int = C.rcl_node_get_domain_id(
		(*C.rcl_node_t)(node),
		&dom,
	)
	*domainID = uint(dom)
	return int(ret)
}

//
func RclNodeAssertLiveliness(node *RclNode) int {
	ret := C.rcl_node_assert_liveliness(
		(*C.rcl_node_t)(node),
	)
	return int(ret)
}

//
func RclNodeGetRmwHandle(node *RclNode) *RmwNode {
	ret := C.rcl_node_get_rmw_handle(
		(*C.rcl_node_t)(node),
	)
	return (*RmwNode)(ret)
}

//
func RclNodeGetRclInstanceID(node *RclNode) uint64 {
	ret := C.rcl_node_get_rcl_instance_id(
		(*C.rcl_node_t)(node),
	)
	return uint64(ret)
}

//
func RclNodeGetGraphGuardCondition(node *RclNode) *RclGuardCondition {
	ret := C.rcl_node_get_graph_guard_condition(
		(*C.rcl_node_t)(node),
	)
	return (*RclGuardCondition)(ret)
}

//
func RclNodeGetLoggerName(node *RclNode) string {
	var cStr *C.char = C.rcl_node_get_logger_name(
		(*C.rcl_node_t)(node),
	)
	return C.GoString(cStr)
}
