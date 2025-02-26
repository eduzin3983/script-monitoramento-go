package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const nome = "Eduardo"
const versao = 1.1
const quantidade_monitoramento = 5
const delay = 5

func main() {
	fmt.Println("Olá, ", nome)
	fmt.Println("Este programa está na versão", versao)

	var site_escolha string
	var sites []string

	for {
		fmt.Print("Qual site(s) deseja monitorar? Digite 0 para sair: ")
		fmt.Scan(&site_escolha)
		if site_escolha != "0" {
			sites = append(sites, site_escolha)
		} else {
			fmt.Println("")
			break
		}
	}

	for {
		fmt.Println("1 - Iniciar Monitoramento")
		fmt.Println("2 - Exibir Logs")
		fmt.Println("0 - Sair")

		var escolha int
		fmt.Scan(&escolha)
		fmt.Println("")

		switch escolha {
		case 1:
			monitoramento(sites)
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
		}
	}

}

func monitoramento(sites []string) {
	fmt.Println("Monitorando...")
	fmt.Println("")

	for i := 0; i < quantidade_monitoramento; i++ {
		fmt.Println("Teste ", i+1, ":")

		for i, site := range sites {
			fmt.Println("Testando site", i+1, ":", site)

			testaSite(site)

			fmt.Println("")
		}

		println("------------------------------------")
		fmt.Println("")

		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
