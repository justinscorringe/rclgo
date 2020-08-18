package ros2

type WaitSet struct {
	rclWaitSet *RclWaitSet
}

func NewZeroInitializedWaitSet() WaitSet {
	zeroWaitset := RclGetZeroInitializedWaitSet()
	return WaitSet{&zeroWaitset}
}

func (w *WaitSet) WaitSetInit(
	numSubs, numGuards, numTimers, numClients, numServices, numEvents int,
	ctx RclContextPtr,
	allo RclAllocator,
) error {
	ret := RclWaitSetInit(
		w.rclWaitSet,
		numSubs,
		numGuards,
		numTimers,
		numClients,
		numServices,
		numEvents,
		ctx,
		allo,
	)
	if ret != Ok {
		return NewErr("RclWaitSetInit", ret)
	}

	return nil
}

func (w *WaitSet) Fini() error {
	ret := RclWaitSetFini(w.rclWaitSet)
	if ret != Ok {
		return NewErr("RclWaitSetFini", ret)
	}

	return nil
}

func (w *WaitSet) GetAllocator(allocator *RclAllocator) error {
	ret := RclWaitSetGetAllocator(w.rclWaitSet, allocator)
	if ret != Ok {
		return NewErr("RclWaitSetGetAllocator", ret)
	}

	return nil
}

func (w *WaitSet) WaitSetAddsubscription(subscription Subscription) error {
	ret := RclWaitSetAddSubscription(w.rclWaitSet, subscription.rclSubscription, nil)
	if ret != Ok {
		return NewErr("RclWaitSetAddSubscription", ret)
	}

	return nil
}

func (w *WaitSet) WaitSetClearSubscriptions() error {
	ret := RclWaitSetClear(w.rclWaitSet)
	if ret != Ok {
		return NewErr("RclWaitSetClear", ret)
	}

	return nil
}
