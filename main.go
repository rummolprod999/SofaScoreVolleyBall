package main

import "time"

func init() {
	CreateEnv()
}

func main() {
	defer SaveStack()
	currentTime := time.Now()
	if currentTime.Hour() > 19 || currentTime.Hour() < 12 {
		return
	}
	server := SofaVolleyBall{}
	server.run()
}
