package main
import (
	"fmt"
	"time"
)
func main(){
	fmt.Println("Iniciando o programa")

	mensagens :=[]string{"Ola", "do", "Go", "com", "gourotines!"}

	for i, msg := range mensagens {
		//a palavra go antes da chamada de uma funcao inicia uma goroutine. esta func sera exec em uma nova goroutine
		//para cada mensagem, cria uma goroutine usando uma closure function
		go func(indice int, texto string){
			fmt.Printf("Goroutine %d: %s \n", indice, texto)
		} (i, msg)
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Fim do programa")
	//A saida aparecera em ordens diferentes devido a forma concorrente de execucao.
}