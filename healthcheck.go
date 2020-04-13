package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	for {
		initMenu()

		var enviroment int
		fmt.Println("")
		fmt.Scanln(&enviroment)
		fmt.Println("")

		switch enviroment {
		case 1, 2:
			initMonitoring(enviroment)
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido")
			os.Exit(-1)
		}
	}
}

func initMenu() {
	fmt.Println("1 - Iniciar Monitoramento [Produção]")
	fmt.Println("2 - Iniciar Monitoramento [Sandbox]")
	fmt.Println("0 - Sair do Programa")
}

func initMonitoring(enviroment int) {

	sites := fileReading(enviroment)

	var typeEnviroment string

	if enviroment == 1 {
		typeEnviroment = "Produção"
	} else {
		typeEnviroment = "Sandbox"
	}

	for i, site := range sites {
		fmt.Println("Testando site", i+1, "em", typeEnviroment, ":", site)
		isHealthcheck := testSite(site)
		fmt.Println("isHealthcheck =", isHealthcheck)
		fmt.Println("")
	}
}

func fileReading(enviroment int) []string {

	const (
		prodution = "sitesProducao.txt"
		sandbox   = "sitesSandbox.txt"
	)
	var sites []string

	var siteEnviroment string

	if enviroment == 1 {
		siteEnviroment = prodution
	} else {
		siteEnviroment = sandbox
	}

	file, err := os.Open(siteEnviroment)

	if err != nil {
		fmt.Println("Deu ruim!!!", err)
	}

	reading := bufio.NewReader(file)

	for {
		lineFile, err := reading.ReadString('\n')
		lineFile = strings.TrimSpace(lineFile)

		sites = append(sites, lineFile)

		if err == io.EOF {
			break
		}
	}
	file.Close()

	return sites
}

func testSite(site string) bool {

	var isHealthCheck = false

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Deu ruim!!!!", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "foi carregado com sucesso!!")
		isHealthCheck = true
	} else {
		fmt.Println("O site:", site, "deu ruim!! O status code é", resp.StatusCode)
	}
	return isHealthCheck
}
