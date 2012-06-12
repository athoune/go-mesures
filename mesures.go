package main

var Mesures map[string]int64 = make(map[string]int64)

func main() {
	go startSocket()
	startHttp()
}
