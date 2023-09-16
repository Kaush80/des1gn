package main

import "fmt"

func main() {
	var bldr = newNotificationBuilder()
	bldr.SetTitle("New Notification")
	bldr.SetIcon("icon.png")
	bldr.SetImage("image.jpg")
	bldr.SetPriority(10)
	bldr.SetMessage("This is a noti")
	bldr.SetNoType("Alert")
	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("Error creating notification object ", err)
	} else {
		fmt.Println("Notification:", notif)
	}

}
