package main

import "time"

func addlog(msg string) {
	thisTime := time.Now().String()[:19] + " "
	writeFile(logPath, thisTime+msg+"\n")
}
