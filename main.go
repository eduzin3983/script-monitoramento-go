package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	nome := "Eduardo"
	versao := 1.1

	fmt.Println("Olá, ", nome)
	fmt.Println("Este programa está na versão", versao)

	for {
		fmt.Println("1 - Iniciar Monitoramento")
		fmt.Println("2 - Exibir Logs")
		fmt.Println("0 - Sair")

		var escolha int
		fmt.Scan(&escolha)

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
	site := "https://www.alura.com.br"

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!\n")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)

	}
}
