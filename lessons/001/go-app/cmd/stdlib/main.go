package main

import (
	"booster/first-api/device"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type server struct {
	state int

	// prometheus metrics
	// metrics *metrics
}

func NewServer() *server {
	state := 0
	return &server{state: state}
}

func JSON(writer http.ResponseWriter, s interface{}) {
	b, err := json.Marshal(s)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(b)
}

func (s *server) getDevices(w http.ResponseWriter, req *http.Request) {
	devices := []device.Device{
		{Id: 0, Mac: "5F-33-CC-1F-43-82", Firmware: "2.1.6"},
		{Id: 1, Mac: "EF-2B-C4-F5-D6-34", Firmware: "2.1.5"},
		{Id: 2, Mac: "62-46-13-B7-B3-A1", Firmware: "3.0.0"},
		{Id: 3, Mac: "96-A8-DE-5B-77-14", Firmware: "1.0.1"},
		{Id: 4, Mac: "7E-3B-62-A6-09-12", Firmware: "3.5.6"},
	}

	JSON(w, devices)
}

func (s *server) getHealth(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK"))
}

func (s *server) getInfo(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("stdlib"))
}

func (s *server) getReady(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`OK`))
}

func main() {
	mux := http.NewServeMux()
	s := NewServer()
	
	mux.HandleFunc("GET /devices", s.getDevices)
	mux.HandleFunc("GET /healthz", s.getHealth)
	mux.HandleFunc("GET /about", s.getInfo)
	mux.HandleFunc("GET /ready", s.getReady)
	mux.Handle("GET /metrics", promhttp.Handler())
	fmt.Println("Server is running on 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
