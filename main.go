package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Definição de constantes do programa
const nome = "Eduardo" // Coloque seu Nome
const versao = 1.1
const quantidade_monitoramento = 3 // Número de vezes que o monitoramento será executado
const delay = 5                    // Tempo de espera entre os testes (em minutos). Caso você queira mudar para segundos ou horas, vá até a função monitoramento e modifique o time.Sleep.

func main() {
	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão", versao)

	for {
		// Exibe o menu de opções
		fmt.Println("1 - Iniciar Monitoramento")
		fmt.Println("2 - Exibir Logs")
		fmt.Println("0 - Sair")

		var escolha int
		fmt.Scan(&escolha)
		fmt.Println("")

		// Verifica a escolha do usuário
		switch escolha {
		case 1:
			monitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
		}
	}
}

// Função que executa o monitoramento dos sites
func monitoramento() {
	fmt.Print("Monitorando...\n\n")

	// Lê os sites do arquivo "sites.txt"
	sites := leArquivoSites()

	// Realiza o teste de cada site a quantidade definida
	for i := 0; i < quantidade_monitoramento; i++ {
		fmt.Println("Teste", i+1, ":")

		for i, site := range sites {
			fmt.Println("Testando site", i+1, ":", site)
			testaSite(site) // Chama a função que testa o site
			fmt.Println("")
		}

		println("------------------------------------\n")

		// Aguarda o tempo definido antes do próximo teste
		time.Sleep(delay * time.Minute) // time.Minute (Minutos) || time.Second (Segundos) || time.Hour (Horas)
	}
}

// Função que faz a requisição HTTP para testar se o site está online
func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	} else {
		// Verifica se a resposta foi bem-sucedida (código 200)
		if resp.StatusCode == 200 {
			fmt.Println("Site:", site, "foi carregado com sucesso!")
			registraLogs(site, true)
		} else {
			fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
			registraLogs(site, false)
		}
	}
}

// Função que lê os sites do arquivo "sites.txt"
func leArquivoSites() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt") // Abre o arquivo de sites
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		os.Exit(0)
	}

	// Lê o arquivo linha por linha
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha) // Remove espaços e quebras de linha

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

// Função que registra os logs do monitoramento no arquivo "logs.txt"
func registraLogs(site string, status bool) {
	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	// Escreve a entrada de log com data, hora, site e status
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

// Função que exibe o conteúdo do arquivo de logs "logs.txt"
func imprimeLogs() {
	arquivo, err := os.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
