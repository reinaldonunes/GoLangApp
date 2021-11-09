package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	showIntro()
	for {
		showMenu()
		opt := readMenu()
		switch opt {
		case 1:
			initMonitor()
		case 2:
			fmt.Println("Comando do 2")
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Comando inválido")
		}
	}
}

// FUNCTIONS
func showIntro() {
	fmt.Println("=========PROGRAMA DE MONITORAMENTO DE SITES=========")
	nome := "Reinaldo"
	versao := 1.0

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está em sua versao", versao)
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Encerrar")
}

func readMenu() int {
	var optSelected int
	fmt.Scan(&optSelected)

	return optSelected
}

func initMonitor() {
	fmt.Println("Iniciando monitoramento...")
	sites := []string{"https://google.com.br", "https://github.com", "https://alura.com.br"}

	for i := 0; i < len(sites); i++ {
		resp, _ := http.Get(sites[i])

		if resp.StatusCode == 200 {
			fmt.Println("Site:", sites[i], "foi carregado com sucesso!!!")
		} else {
			fmt.Println("Site:", sites[i], "apresentou falha no carregamento!!!")
		}
	}
}
