package main

import (
	"github.com/0xAX/notificator"
	"log"
	"time"
)

func main() {
	log.Print("Started running timer")
	var numberOfHours int
	var message string
	isWeekend := false
	var currentTime time.Time
	for {
		numberOfHours += 1
		currentTime = time.Now()
		today := currentTime.Weekday()
		hour := currentTime.Hour()
		isWeekend = today == time.Sunday || today == time.Friday || today == time.Saturday
		switch {
		case (hour >= 22 || hour <= 3) && !isWeekend:
			message = "Late night on a weekday. Sleep!."
		case numberOfHours == 1:
			message = "One hour! Take a break"
		case numberOfHours == 2:
			message = "Two hours! Take a break now."
		case numberOfHours > 2: // or default:
			message = "Take a break now. Seriously."
		}

		log.Print("Sleeping program for one hour")
		time.Sleep(1 * time.Hour)

		showNotification(message)
	}

}

func showNotification(message string) {
	notify := notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "Timer",
	})

	notify.Push("Take a Break", message, "", notificator.UR_CRITICAL)

	// Check errors
	//	err := notify.Push("error", "ops =(", "", notificator.UR_CRITICAL)

	//if err != nil {
	//	log.Fatal(err)
	//}

}
