package main

import (
	"fmt"
)

const maxCapacity = 10

// Fan-out Pattern, a ideia é que n workers estejam lendo do canal principal.
// Cada mensagem do canal principal é posta no canal que apenas um worker ouve.
// Dessa forma distribuimos a carga

func main() {
	// Na medida que o channel é buferizado, garantimos a criacao de apenas maxCapacity req
	req_chan := make(chan string, maxCapacity)
	workers := []chan string{}

	// Kick-off 4 workers
	for i := 0; i < 4; i++ {
		c := make(chan string)
		workers = append(workers, c)
		name := fmt.Sprintf("Worker %d", i)
		go exec_req(name, c)
	}

	for {
		// Para evitar que sejam criadas mais que maxCapacity reqs @Producer
		go create_req(req_chan)
		for _, w := range workers {
			w <- <-req_chan
		}
	}

}

// @Consumer, should be just listening to the channel -- WORKER
func exec_req(name string, req_chan chan string) {
	fmt.Println("Worker %s", name)
	for req := range req_chan {
		fmt.Println("Worker --" + name + "got message: " + req)
	}
}

// @Producer
func create_req(reqChan chan string) {

	i := 0
	for {
		reqChan <- fmt.Sprintf("webReq: %d \n", i)
		i++
	}
}
