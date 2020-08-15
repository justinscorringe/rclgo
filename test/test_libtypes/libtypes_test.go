package libtypes_test

import (
	"fmt"
	"os"
	"testing"

	types "github.com/edwinhayes/rosgo/libgengo"
)

func TestLibTypes(t *testing.T) {
	context, err := types.NewMsgContext([]string{os.Getenv("ROS_PACKAGE_PATH")})
	if err != nil {
		t.Fatal("could not get ros 2 message types")
	} else {
		fmt.Print(context)
	}

}
