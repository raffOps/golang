package main

import (
	"fmt"
	"net/http"
	"os"
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

func initMonitoring(url string) error {
	get, err := http.Get(url)
	println(url, get.Status)
	return err
}

func main() {
	url := "https://alura.com.br"
	printIntroduction()
	option, err := getOption()
	if err != nil {
		os.Exit(1)
	}
	switch option {
	case 1:
		err := initMonitoring(url)
		if err != nil {
			println(err)
			os.Exit(1)
		}
	default:
		println("Opção inválida")
	}
}
