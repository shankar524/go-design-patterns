package main

import "fmt"

type IBuilder interface {
	Build() (*Notification, error)
}

func main() {
	// TODO: Create NotificationBuilder
	var builder = newNotificationBuilder()

	// TODO: Use builder to set some properties
	builder.SetTitle("New notification")
	builder.SetSubtitle("First notification")
	builder.SetMessage("Hello")
	builder.SetImage("image.jpg")
	builder.SetIcon("icon.png")
	builder.SetPriority(10)
	builder.SetNotificationType("alert")

	// print
	printBuilder(builder)

	// fix error and print
	builder.SetPriority(SUCCESS)
	printBuilder(builder)

}

func printBuilder(builder IBuilder) {
	// TODO: Use build function to create a finished product
	notif, err := builder.Build()
	if err != nil {
		fmt.Printf("Error creating notification. Error: %+v\n", err)
	} else {
		fmt.Printf("Notification: %+v\n", notif)
	}
}
