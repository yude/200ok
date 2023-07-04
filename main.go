package main

func main() {
	LoadConfig()
	InitStatuses()
	Greeting()
	StartRoutine()

	select {}
}
