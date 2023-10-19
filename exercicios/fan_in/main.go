package main

import (
	"fmt"
	"sync"
)

func main() {

	ch1 := request_stream()
	ch2 := request_stream()

	go populate(3, ch1)
	go populate(4, ch2)
	// fanin Pattern, processo de merge de channels
	// ______
	//       \
	// ______/ merged chan

	chm := make(chan string, 7)
	done := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(7)
	go func() {

		for {
			select {
			case v1 := <-ch1:
				chm <- v1
				wg.Done()
			case v2 := <-ch2:
				chm <- v2
				wg.Done()
			}

		}
	}()

	// So pode chamar ela, quando as outras estiverem terminadas
	wg.Wait()
	close(chm)
	go ingest(chm, done)

	// feio, trocar para wg
	<-done

}

// Fica escrevendo sempre que puder, até bater o limite maximo N de mensagens no canal
func populate(n int, ch chan string) {

	i := 0
	for i < n {
		ch <- fmt.Sprintf("String number %d", i)
		i++
	}

}

// So cria um canal -- nasce da assinatura
func request_stream() chan string {
	return make(chan string)
}

// Recebe um canal unico, que deve ser resultado do merge dos ch1, e ch2. @Consumer
func ingest(in chan string, ctrl chan int) {

	fmt.Println("GET INTO INGEST")
	for v := range in {
		fmt.Println(v + " NO INGEST")
	}

	// Devem ser printado 7 linhas.

	ctrl <- 0

}
