package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	quantidadeMonitoramento int = 10
	delayMonitoramento          = 500 // em ms
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

func initMonitoring(urls *[]string) error {
	println("Monitoring")
	for i := 0; i <= quantidadeMonitoramento; i++ {
		for _, value := range *urls {
			resp, err := http.Get(value)
			if err != nil {
				println(err)
				return err
			}
			saveLog(time.Now().Format(time.DateTime), value, resp.Status)
		}
		time.Sleep(delayMonitoramento * time.Millisecond)
	}
	return nil
}

func saveLog(timestamp string, site string, status string) error {
	log := timestamp + " - " + site + " - " + status
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	file.WriteString(log + "\n")
	file.Close()
	return nil
}

func showLogs(filename string) {
	logs, _ := readFile(filename)
	for _, value := range logs {
		println(value)
	}
	println()

}

func readFile(filename string) ([]string, error) {
	// file, err := os.ReadFile(filename) // lê tudo de uma vez
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var sites []string
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		readString = strings.TrimSpace(readString)
		sites = append(sites, readString)
		if err == io.EOF {
			break
		}
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	return sites, err
}

func main() {
	printIntroduction()
	for {
		option, err := getOption()
		if err != nil {
			println("Opção inválida")
			continue
		}
		switch option {
		case 1:
			urls, err := readFile("sites.txt")
			if err != nil {
				println(err)
				os.Exit(1)
			}
			errMonitoring := initMonitoring(&urls)
			if errMonitoring != nil {
				println(errMonitoring)
				os.Exit(1)
			}
		case 2:
			showLogs("logs.txt")
		case 3:
			os.Exit(0)
		default:
			println("Opção inválida")
		}
	}
}
