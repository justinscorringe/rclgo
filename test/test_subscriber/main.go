package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	rclgo "github.com/justinscorringe/rclgo/ros2"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	msgChan := make(chan string, 1)
	go func() {
		// Receive input in a loop
		for {
			var s string
			fmt.Scan(&s)
			// Send what we read over the channel
			msgChan <- s
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
	msgType, err := rclgo.NewDynamicMessageType("std_msgs/String")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Creating the subscriber! \n")
	err = mySub.Init(mySubOpts, myNode, "/myGoTopic", msgType)
	if err != nil {
		log.Fatalf("SubscriptionsInit: %s", err)
	}

	// //Creating the msg type
	// var myMsg types.GeometryMsgsTwist
	// myMsg.InitMessage()

	time.Sleep(100 * time.Millisecond)
	i := 0

loop:
	for {
		fmt.Println("Subscriber run loop!")
		msg, err := mySub.TakeMessageRaw(msgType)
		if err == nil {
			fmt.Printf("(Suscriber) Received %v\n", msg)
			i = 0
		} else {
			fmt.Println(err)
			time.Sleep(100 * time.Millisecond)
			if i > 10 {
				log.Fatal("dead")
			}
			i++
			goto loop
		}

		time.Sleep(100 * time.Millisecond)

		select {
		case <-sigs:
			fmt.Println("Got shutdown, exiting")
			break loop
		case <-msgChan:
		}
	}

	fmt.Printf("Shutting down!! \n")

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
