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

const nome = "Eduardo"
const versao = 1.1
const quantidade_monitoramento = 3
const delay = 5

func main() {
	fmt.Println("Olá, ", nome)
	fmt.Println("Este programa está na versão", versao)

	for {
		fmt.Println("1 - Iniciar Monitoramento")
		fmt.Println("2 - Exibir Logs")
		fmt.Println("0 - Sair")

		var escolha int
		fmt.Scan(&escolha)
		fmt.Println("")

		switch escolha {
		case 1:
			monitoramento()
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

func monitoramento() {
	fmt.Println("Monitorando...")
	fmt.Println("")

	sites := leArquivoSites()

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
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	} else {
		if resp.StatusCode == 200 {
			fmt.Println("Site:", site, "foi carregado com sucesso!")
		} else {
			fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		}
	}
}

func leArquivoSites() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {

			break
		}
	}

	arquivo.Close()

	return sites
}
