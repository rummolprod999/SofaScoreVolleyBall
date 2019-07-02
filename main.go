package main

func init() {
	CreateEnv()
}

func main() {
	defer SaveStack()
	server := SofaVolleyBall{}
	server.run()
}
