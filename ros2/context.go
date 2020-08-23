package ros2

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

//
func RclContextFini(ctx RclContextPtr) int {
	var ret C.int32_t = C.rcl_context_fini(
		(*C.rcl_context_t)(ctx),
	)
	return int(ret)
}

//
func RclContextIsValid(ctx RclContextPtr) bool {
	var ret C.bool = C.rcl_context_is_valid(
		(*C.rcl_context_t)(ctx),
	)
	return bool(ret)
}

// Context encapsulates the non-global state of an init/shutdown cycle.
type Context struct {
	rclContext RclContextPtr
}

// NewZeroInitializedContext returns a zero initialization context object.
func NewZeroInitializedContext() Context {
	ctxPtr := GetZeroInitializedContextPtr()
	return Context{rclContext: ctxPtr}
}

// Init finalizes a context.
func (ctx *Context) Init() error {
	var ret int

	var opts = RclGetZeroInitializedInitOptions()
	alloc := RclGetDefaultAllocator()

	ret = RclInitOptionsInit(&opts, alloc)
	if ret != Ok {
		return NewErr("RclInitOptionsInit", ret)
	}

	ret = RclInit(
		0,
		[]string{},
		&opts,
		ctx.rclContext,
	)
	if ret != Ok {
		return NewErr("RclInit", ret)
	}

	ret = RclInitOptionsFini(&opts)
	if ret != Ok {
		return NewErr("RclInitOptionsFini", ret)
	}

	return nil
}

// Fini finalizes a context.
func (ctx *Context) Fini() error {
	ret := RclContextFini(ctx.rclContext)
	if ret != Ok {
		return NewErr("RclContextFini", ret)
	}

	return nil
}

// Shutdown shuts down the context.
func (ctx *Context) Shutdown() error {
	ret := RclShutdown(ctx.rclContext)
	if ret != Ok {
		return NewErr("RclShutdown", ret)
	}

	return nil
}

// IsValid returns `true` if the context is currently valid, otherwise `false`.
func (ctx *Context) IsValid() bool {
	ret := RclContextIsValid(ctx.rclContext)
	return bool(ret)
}
