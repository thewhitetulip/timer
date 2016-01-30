#Timer

Timer is a simple Go application written to remind the user to take breaks every hour. It was created to remind me at odd time in the night that I should sleep.

There is another alternative to this solution, to use `crontab`, but that approach didn't work for me since if I start my machine at 10:59PM and the crontab is scheduled to run every one hour, then it tells me to take a break in one minute.

In addition with a message every hour, this application also checks if today is a weekday or weekend. If it is a weekday and it is after 11PM it tells you to sleep rather than take breaks.

It is cross platform as we use [notificator] (github.com/0xAX/notificator) for notifications.

##Build instructions:

1. `go get github.com/thewhitetulip/timer`
1. `go build`
1. './timer'

##Hacking
If you want more features please raise an issue and we'll talk. PRs to extend the application are welcome!
 
License: MIT
