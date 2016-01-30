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

		if (hour >= 22 || hour <= 6) && (isWeekend == false) {
			message = "Late night on a weekday. Sleep!."
		} else {
			if numberOfHours == 1 {
				message = "One hour! Take a break"
			} else if numberOfHours == 2 {
				message = "Two hours! Take a break now."
			} else {
				message = "Take a break now. Seriously."
			}
		}

		log.Print("Sleeping program for one hour")
		time.Sleep(1 * time.Second)

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
