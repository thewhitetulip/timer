package main

import (
	"log"
	"os/exec"
	"time"
)

func main() {
	var numberOfHours int
	var message string
	for {
		numberOfHours += 1
		hour := time.Now().Hour()
		if hour >= 22 || hour <= 3 {
			message = "You should seriously sleep now."
		} else {
			if numberOfHours == 1 {
				message = "One hour! Take a break"
			} else if numberOfHours == 2 {
				message = "Two hours! Take a break now."
			} else {
				message = "Take a break now. Seriously."
			}
		}
		time.Sleep(2 * time.Hour)

		cmd := exec.Command("notify-send", message)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

	}

}
