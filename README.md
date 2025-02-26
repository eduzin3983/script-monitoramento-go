# Monitor de Sites em Go

Este programa foi criado para estudar um pouco sobre a linguagem **Go** e suas funções. Ele realiza o monitoramento de sites, verificando se estão online e salvando logs das verificações.

## 📋 Requisitos
- Go instalado ([Download Go](https://go.dev))
- Criar um arquivo chamado **`sites.txt`** na mesma pasta do código, contendo a lista de sites a serem monitorados.

## 📂 Como configurar?
1. **Criar o arquivo `sites.txt`**
   - Adicione os sites que deseja monitorar, um por linha. Exemplo:
     ```
     https://www.google.com
     https://www.spotify.com
     https://www.github.com
     ```

2. **Configurar a frequência do monitoramento**
   - Você pode alterar as seguintes constantes no código:
     ```go
     const quantidade_monitoramento = 3 // Quantas vezes cada site será testado
     const delay = 5                    // Intervalo entre os testes (em minutos)
     ```

## 🚀 Como executar?
Abra o terminal na pasta onde está o arquivo **`main.go`** e execute no terminal:

```sh
go run main.go
```

## 🎯 Como funciona?

O programa exibe um menu:

```bash
1 - Iniciar Monitoramento
2 - Exibir Logs
0 - Sair
```
- Quando você escolhe a opção 1, ele começa a monitorar os sites da lista.

- Quando você escolhe a opção 2, ele exibe os logs gravados no arquivo logs.txt.

Os logs ficam registrados no formato:

```bash
26/02/2025 14:30:00 - https://www.google.com - Online: true
26/02/2025 14:35:00 - https://www.github.com - Online: true
26/02/2025 14:40:00 - https://www.siteinexistente.com - Online: false
```

- A opção 0 encerra o programa.
