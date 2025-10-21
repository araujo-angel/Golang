package main

import (
	"fmt"
)

//jobs <-chan int: um canal de entrada, que recebe números para serem processados. apenas leitura
//results chan<- int: um canal de saída, onde o resultado será enviado após o processamento. faz escrita

func worker(id int, jobs <-chan int, results chan<- int){
	for n := range jobs{
		results <- n * n
	}
}

func main(){
	jobs := make(chan int)
	results := make(chan int)

	//start worker in a Goroutine and close channel 'results' in the end
	go func(){
		worker(1, jobs, results)
		close(results)
	}()

	//producer: send 1..5 and close 'jobs' in the end
	//envia dados para o canal de jobs.
	go func(){
		for i := 1; i <= 5; i++ {
			jobs <- i}
		close(jobs)
	}()

	//consumer: reads all the results until the channel is closed
	//é responsável por consumir e imprimir os resultados após o processamento.
	for r := range results{
		fmt.Println("Resultado: ", r)
	}
	fmt.Println("Fim!")
}