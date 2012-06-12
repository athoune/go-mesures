package main

func main() {
	go startMemento()
	go startSocket()
	startHttp()
}
