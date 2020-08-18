package ros2

// #cgo CFLAGS: -I/opt/ros/dashing/include
// #cgo LDFLAGS: -L/opt/ros/dashing/lib -Wl,-rpath=/opt/ros/dashing/lib -lrcl -lrcutils
// #include <rcl/rcl.h>
import "C"

import (
	"fmt"
)

type Allocator struct {
	rclAllocator *RclAllocator
}

type RcutilsErrorState struct {
	msg  string
	file string
	line uint64
}

func (es *RcutilsErrorState) Error() string {
	return es.String()
}

func (es *RcutilsErrorState) String() string {
	if es == nil {
		return ""
	}

	return " msg: " + es.msg +
		"\nfile: " + es.file +
		fmt.Sprintf("\nline: %d\n", es.line)
}

//
func RcutilsGetErrorState() *RcutilsErrorState {
	var errState = C.rcutils_get_error_state()
	if errState == nil {
		return nil
	}
	var ret = RcutilsErrorState{
		msg:  C.GoString(&errState.message[0]),
		file: C.GoString(&errState.file[0]),
		line: uint64(errState.line_number),
	}

	C.rcutils_reset_error()
	return &ret
}

func NewZeroInitializedAllocator() Allocator {
	ret := C.rcutils_get_zero_initialized_allocator()
	zeroAllocator := RclAllocator(ret)
	return Allocator{&zeroAllocator}
}

func NewDefaultAllocator() Allocator {
	ret := C.rcutils_get_default_allocator()
	defAllocator := RclAllocator(ret)
	return Allocator{&defAllocator}
}
