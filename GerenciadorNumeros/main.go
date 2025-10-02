package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func adicionar(numeros []int) ([]int, error) {
	var n int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&n)
	if n < 1 {
		return numeros, errors.New("entrada invalida")
	}
	numeros = append(numeros, n)
	return numeros, nil
}
func listar(numeros []int) error {
	if len(numeros) == 0 {
		return errors.New("array vazio")
	}
	for _, valor := range numeros {
		fmt.Println(valor)
	}
	return nil
}

func ordenar(numeros []int) ([]int, error) {
	if len(numeros) == 0 {
		return nil, errors.New("array vazio")
	}
	opcao := 0
	fmt.Println("Para ordenar crescente digite 1, decrescente digite 2: ")
	fmt.Scanln(&opcao)
	if opcao == 1 {
		return crescente(numeros), nil
	}
	return decrescente(numeros), nil
}

func decrescente(numeros []int) []int {
	slices.SortFunc(numeros, func(a, b int) int {
		return b - a
	})
	return numeros
}

func crescente(numeros []int) []int {
	slices.Sort(numeros)
	return numeros
}

func remover(numeros []int) ([]int, error) {
	var indice int
	fmt.Println("Digite o indice do numero a ser removido: ")
	_, err := fmt.Scanln(&indice)
	if indice < 0 || indice >= len(numeros) || err != nil {
		return numeros, errors.New("entrada invalida")
	}
	numeros = append(numeros[:indice], numeros[indice+1:]...)
	return numeros, nil
}
func estatisticas(numeros []int) (int, int, int, error) {
	if len(numeros) == 0 {
		return 0, 0, 0, errors.New("slice vazio")
	} else {
		return minimo(numeros), maximo(numeros), media(numeros), nil
	}
}

func minimo(numeros []int) int {
	var minimo int
	minimo = slices.Min(numeros)
	return minimo
}

func maximo(numeros []int) int {
	var maximo int
	maximo = slices.Max(numeros)
	return maximo
}

func media(numeros []int) int {
	var media int
	total := 0
	for _, valor := range numeros {
		total += valor
	}
	media = total / len(numeros)
	return media
}

func divisao() (int, error) {
	n1 := 0
	n2 := 0
	fmt.Println("Digite o primeiro numero: ")
	fmt.Scanln(&n1)
	fmt.Println("Digite o segundo numero: ")
	fmt.Scanln(&n2)
	if n2 <= 0 {
		return 0, errors.New("divisor nao pode ser 0")
	}
	return n1 / n2, nil
}

func exibirPares(numeros []int) ([]int, error) {
	var result []int
	for _, n := range numeros {
		if n%2 == 0 {
			result = append(result, n)
		}
	}
	if len(result) == 0 {
		return nil, errors.New("Nao existem numeros pares")
	}
	return result, nil
}
func exportar(numeros []int) error {
	var strin []string
	for _, n := range numeros {
		strin = append(strin, strconv.Itoa(n))
	}
	texto := strings.Join(strin, ",")
	err := os.WriteFile("slice.txt", []byte(texto), 0644)
	if err != nil {
		return fmt.Errorf("erro durante exportar: %w", err)
	}
	return nil
}

func esvaziar(numeros []int) ([]int, error) {
	if len(numeros) == 0 {
		return numeros, errors.New("array vazio")
	}
	return numeros[:0], nil
}

func main() {
	var numeros []int
	opcao := 10
	fmt.Println("\n***** MENU *****")
	for {
		fmt.Print("1 - Adicionar um numero inteiro \n")
		fmt.Print("2 - Listar o array \n")
		fmt.Print("3 - Remover por índice \n")
		fmt.Print("4 - Visualizar estatisticas \n")
		fmt.Print("5 - Dividir \n")
		fmt.Print("6 - Esvaziar a lista \n")
		fmt.Print("7 - Ordenar \n")
		fmt.Print("8 - Exibir pares \n")
		fmt.Print("9 - Exportar array \n")
		fmt.Print("0 - Encerrar \n")
		_, err := fmt.Scanln(&opcao)
		if err != nil {
			errors.New("Entrada invalida")
		}
		switch opcao {
		case 1:
			numeros, err = adicionar(numeros)
			if err != nil {
				fmt.Println("Erro:", err)
			}
		case 2:
			listar(numeros)
			if err != nil {
				fmt.Println("Erro:", err)
			}
		case 3:
			numeros, err = remover(numeros)
			if err != nil {
				fmt.Println("Erro:", err)
			}
		case 4:
			minim, maxim, med, err := estatisticas(numeros)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("Mínimo: %d | Máximo: %d | Média: %d\n\n", minim, maxim, med)
			}
		case 5:
			resultado, err := divisao()
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Resultado da divisao: ", resultado)
			}
		case 6:
			numeros, err = esvaziar(numeros)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Resultado da esvaziar: ", numeros)
			}
		case 7:
			numeros, err = ordenar(numeros)
			if err != nil {
				fmt.Println("Erro:", err)
			}
		case 8:
			pares, err := exibirPares(numeros)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Pares", pares)
			}
		case 9:
			err = exportar(numeros)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Array exportado para este mesmo diretorio com sucesso!")
			}
		case 0:
			if opcao == 0 {
				fmt.Println("Encerrando o programa...")
				return
			}
		}
	}
}
