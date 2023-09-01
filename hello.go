package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	logs []string
	urls []string{"https://alura.com.br", "google.com"}
)

func printIntroduction() {
	name := "user"
	version := 1.2
	fmt.Println("Olá", name, "\nVersão: ", version)
}

func getOption() (int, error) {
	var option int
	fmt.Println("1- iniciar monitoramento\n2- Exibir log\n3- sair")
	_, err := fmt.Scan(&option)
	return option, err
}

func initMonitoring() {
	for _, value := range urls {
		resp, _ := http.Get(value)
		logs = append(logs, strings.Join([]string{value, resp.Status}, " - "))
	}
	println("Monitoring")
}

func showLogs() {
	for _, value := range logs {
		println(value)
	}
}

func main() {
	printIntroduction()
	for {
		option, err := getOption()
		if err != nil {
			os.Exit(1)
		}
		switch option {
		case 1:
			initMonitoring(urls)
		case 2:
			showLogs()
		case 3:
			os.Exit(0)
		default:
			println("Opção inválida")
		}
	}
}
