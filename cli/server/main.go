package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	w.WriteHeader(http.StatusOK)
}

func handleRequests() {

	ready := &atomic.Value{}
	ready.Store(0)
	h := NewReadyHandler(ready)

	go FakeFunctionToSimulateTimeTakenToBeReady(ready)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/up", h.UpAndRunningHandler)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}

func (h *ReadyHandler) UpAndRunningHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpAndRunningHandler")
	if val, ok := h.ready.Load().(int); ok && val == 1 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

type ReadyHandler struct {
	ready *atomic.Value
}

func NewReadyHandler(ready *atomic.Value) *ReadyHandler {
	return &ReadyHandler{
		ready: ready,
	}
}

func FakeFunctionToSimulateTimeTakenToBeReady(ready *atomic.Value) {
	time.Sleep(20*time.Second)
	ready.Store(1)
}
