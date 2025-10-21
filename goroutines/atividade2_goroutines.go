package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Estrutura(classe) para representar uma leitura do sensor
type Leitura struct {
	Sensor string
	Valor  float64
	Hora   string
}

func sensorTemperatura(canal chan<- Leitura) {
	for {
		temp := 15.0 + rand.Float64()*20.0
		leitura := Leitura{
			Sensor: "Temperatura",
			Valor:  temp,
			Hora:   time.Now().Format("15:04:05"),
		}
		canal <- leitura
		time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second)
	}
}

func sensorPressao(canal chan<- Leitura) {
	for {
		pressao := 980.0 + rand.Float64()*60.0
		leitura := Leitura{
			Sensor: "Pressão",
			Valor:  pressao,
			Hora:   time.Now().Format("15:04:05"),
		}
		canal <- leitura
		time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second)
	}
}

func sensorUmidade(canal chan<- Leitura) {
	for {
		umidade := 30.0 + rand.Float64()*60.0
		leitura := Leitura{
			Sensor: "Umidade",
			Valor:  umidade,
			Hora:   time.Now().Format("15:04:05"),
		}

		canal <- leitura

		time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	canalLeituras := make(chan Leitura)

	fmt.Println("Iniciando sistema de monitoramento de sensores...\n")

	go sensorTemperatura(canalLeituras)
	go sensorPressao(canalLeituras)
	go sensorUmidade(canalLeituras)
	contador := 0
	for {
		select {
		case leitura := <-canalLeituras:
			contador++

			// Formata a unidade de medida baseada no sensor
			var unidade string
			switch leitura.Sensor {
			case "Temperatura":
				unidade = "°C"
			case "Pressão":
				unidade = "hPa"
			case "Umidade":
				unidade = "%"
			}

			fmt.Printf("[%s] %s enviou: %.2f %s (Leitura #%d)\n",
				leitura.Hora,
				leitura.Sensor,
				leitura.Valor,
				unidade,
				contador)

			if contador >= 20 {
				fmt.Println("\nSistema finalizado após 20 leituras")
				return
			}
		}
	}
}
