package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/justinscorringe/rclgo"
	"github.com/justinscorringe/rclgo/types"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	msg := make(chan string, 1)
	go func() {
		// Receive input in a loop
		for {
			var s string
			fmt.Scan(&s)
			// Send what we read over the channel
			msg <- s
		}
	}()
	// Initialization
	ctx := rclgo.NewZeroInitializedContext()
	err := ctx.Init()
	if err != nil {
		log.Fatalf("rcl.Init: %s", err)
	}
	myNode := rclgo.NewZeroInitializedNode()
	myNodeOpts := rclgo.NewNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	err = myNode.Init("GoSubscriber", "", ctx, myNodeOpts)
	if err != nil {
		log.Fatalf("NodeInit: %s", err)
	}

	//Create the subscriptor
	mySub := rclgo.NewZeroInitializedSubscription()
	mySubOpts := rclgo.NewSubscriptionDefaultOptions()

	//Creating the type
	msgType := rclgo.NewDynamicMessageType("std_msgs/String")

	fmt.Printf("Creating the subscriber! \n")
	err = mySub.Init(mySubOpts, myNode, "/myGoTopic", msgType.cSpec)
	if err != nil {
		log.Fatalf("SubscriptionsInit: %s", err)
	}

	//Creating the msg type
	var myMsg types.GeometryMsgsTwist
	myMsg.InitMessage()

	time.Sleep(100 * time.Millisecond)

loop:
	for {
		fmt.Println("Subscriber run loop!")
		err = mySub.TakeMessage(&myMsg.MsgInfo, myMsg.Data())
		if err == nil {
			fmt.Printf("(Suscriber) Received %v\n", myMsg.GetDataMap())
		} else {
			fmt.Print(err)
			time.Sleep(100 * time.Millisecond)
			goto loop
		}

		time.Sleep(100 * time.Millisecond)
		select {
		case <-sigs:
			fmt.Println("Got shutdown, exiting")
			break loop
		case <-msg:
		}
	}

	fmt.Printf("Shutting down!! \n")

	myMsg.DestroyMessage()
	err = mySub.SubscriptionFini(myNode)
	if err != nil {
		log.Fatalf("SubscriptionFini: %s", err)
	}

	err = myNode.Fini()
	if err != nil {
		log.Fatalf("NodeFini: %s", err)
	}

	err = ctx.Shutdown()
	if err != nil {
		log.Fatalf("rcl.Shutdown: %s", err)
	}
}
