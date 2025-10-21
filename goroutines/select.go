package main

import (
	"fmt"
	"time"
)

func main() {
	// Declaração das variáveis e canais
	done := make(chan bool)
	ch := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan bool)
	x := 42

	// Goroutine para enviar valores periodicamente
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
		done <- true // sinaliza o fim
	}() //executa a funcao imediatamente

	for {
		select {
		case <-done:
			fmt.Println("Encerrando loop.")
			return
		case v := <-ch:
			fmt.Println("Recebeu de ch:", v)
		case v := <-ch2:
			fmt.Println("Recebeu de ch2:", v)
		case ch3 <- x:
			fmt.Println("Escreveu em ch3:", x)
		case <-ch4:
			fmt.Println("Obteve valor de ch4, mas ignorou ele")
		default:
			fmt.Println("Nada a fazer")
			time.Sleep(200 * time.Millisecond)
		}
	}
}
