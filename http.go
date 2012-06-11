package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func mesureServer(w http.ResponseWriter, req *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "text/event-stream")
	h.Set("Transfer-Encoding", "chunked")
	h.Set("Cache-Control", "no-cache")
	h.Set("Connection", "keep-alive")
	h.Set("X-Accel-Buffering", "no")
	w.WriteHeader(200)
	w.Write([]byte("data: "))
	j, _ := json.Marshal(Mesures)
	w.Write(j)
	w.Write([]byte("\r\n\r\n"))
	channel := Suscribe(func(m msg) {})
	f, f_ok := w.(http.Flusher)
	if f_ok {
		f.Flush()
	}
	for {
		m := <-channel.channel
		msg := fmt.Sprintf("data: {\"%s\": %d}\r\n\r\n", m.key, m.value)
		w.Write([]byte(msg))
		if f_ok {
			f.Flush()
		}
	}
}

func startHttp() {
	http.HandleFunc("/events", mesureServer)
	log.Printf("About to start http://localhost:8000")
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		panic(err)
	}
}
