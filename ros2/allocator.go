package ros2

type Allocator struct {
	rclAllocator *RclAllocator
}

func NewZeroInitializedAllocator() Allocator {
	zeroAllocator := RcutilsGetZeroInitializedAllocator()
	return Allocator{&zeroAllocator}
}

func NewDefaultAllocator() Allocator {
	defAllocator := RcutilsGetDefaultAllocator()
	return Allocator{&defAllocator}
}
