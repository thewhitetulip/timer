package main

import (
	"log"
	"os/exec"
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
		if today == 0 || today == 5 || today == 6 {
			isWeekend = true
		}
		if (hour >= 22 || hour <= 3) && (isWeekend == false) {
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
		time.Sleep(1 * time.Hour)

		cmd := exec.Command("notify-send", message)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

	}

}
