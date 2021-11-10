package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 2
const delay = 5

func main() {
	showIntro()
	for {
		showMenu()
		opt := readMenu()
		switch opt {
		case 1:
			initMonitor()
		case 2:
			fmt.Println("Exibindo logs...")
			printLogs()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Comando inválido")
			os.Exit(-1)
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
	sites := readUrls()
	//sites := []string{"https://google.com.br", "https://github.com", "https://alura.com.br"}

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando o site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!!!")
		logs(site, true)
	} else {
		fmt.Println("Site:", site, "apresentou falha no carregamento!!!")
		logs(site, false)
	}
}

func readUrls() []string {
	var sites []string
	file, err := os.Open("urls.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro!", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}
	file.Close()
	return sites
}

func logs(site string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func printLogs() {
	file, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}
