package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	args := os.Args[1:]
	n := args[0]

	in, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println("Err during conversion")
	}

	wg := &sync.WaitGroup{}
	wg.Add(in)

	for i := 0; i < in; i++ {
		go func() {
			randomInt := rand.Intn(in) + 1 // Generating a random number from 1 - N
			sleep := time.Duration(randomInt) * time.Second

			fmt.Println("This routine will sleep", sleep)
			time.Sleep(sleep)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Main informs - the N value is %d", in)
}
