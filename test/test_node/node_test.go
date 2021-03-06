package rclgo_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/justinscorringe/rclgo"
)

func TestNodeCreation(t *testing.T) {
	// Initialization
	ctx := rclgo.NewZeroInitializedContext()
	err := ctx.Init()
	if err != nil {
		t.Fatalf("rcl.Init failed: %s\n", err)
	}

	myNode := rclgo.NewZeroInitializedNode()
	myNodeOpts := rclgo.NewNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	err = myNode.Init("fakeNameForNode", "", ctx, myNodeOpts)
	if err != nil {
		t.Fatalf("NodeInit failed: %s\n", err)
	}

	time.Sleep(5 * time.Second) // or runtime.Gosched() or similar per @misterbee

	err = myNode.Fini()
	if err != nil {
		t.Fatalf("NodeFini failed: %s\n", err)
	}

	err = ctx.Shutdown()
	if err != nil {
		t.Fatalf("rcl.Shutdown failed: %s\n", err)
	}
}
