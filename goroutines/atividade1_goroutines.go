package main

import (
	"fmt"
	"sync"
)

// WaitGroup é uma estrutura de sincronização que permite aguardar um grupo de goroutines terminar.
//usamos &wg, pois ele é o ponteiro da variável, entao trabalhamos usando sempre o mesmo, e nao a copia

func escritora(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		valor := id*100 + i
		ch <- valor
	}
	fmt.Printf("Escritora %d finalizada\n", id)
}

func leitora(ch <-chan int) {
	fmt.Println("Leitora iniciada")

	contador := 0
	for valor := range ch {
		contador++
		fmt.Printf("Leitora recebeu [%d]: %d\n", contador, valor)
	}

	fmt.Printf("Leitora finalizada (total: %d valores)\n", contador)
}

func iniciarComunicacao() {
	canal := make(chan int)

	// aguarda as escritoras terminarem
	var wg sync.WaitGroup

	fmt.Println("Iniciando comunicação entre goroutines...\n")

	wg.Add(2)
	go escritora(1, canal, &wg)
	go escritora(2, canal, &wg)
	go leitora(canal)
	wg.Wait()
	close(canal)

	fmt.Println("\nTodas as escritoras finalizaram, canal fechado")
}

func main() {
	iniciarComunicacao()
	fmt.Println("Programa finalizado!")
}
