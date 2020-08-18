package ros2

// Err implements the Error interface for rcl errors.
type Err struct {
	msg   string
	ret   RCLRetT
	state *RcutilsErrorState
}

// NewErr gets the error state from rcutils and wraps it in an Err.
// Intended for internal use only.
func NewErr(msg string, ret int) Err {
	return Err{
		msg:   msg,
		ret:   RCLRetT(ret),
		state: RcutilsGetErrorState(),
	}
}

func (e Err) Error() string {
	if e.ret == Ok {
		return e.msg
	}
	return e.msg + ": " + e.ret.String() + "\n" + e.state.Error()
}

// Unwrap returns the underlying rcl error code.
func (e Err) Unwrap() error { return e.ret }
