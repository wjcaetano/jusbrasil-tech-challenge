# Guia do Usuário - JusBrasil Tech Challenge

Bem-vindo ao guia do **JusBrasil Tech Challenge**! Aqui você encontrará instruções detalhadas para configurar, executar e utilizar o projeto.

## Índice
1. [Visão Geral do Projeto](#visão-geral-do-projeto)
2. [Configuração do Ambiente](#configuração-do-ambiente)
3. [Uso do Makefile](#uso-do-makefile)
4. [Documentação da API](#documentação-da-api)
5. [Configuração Goland](docs/guide/dev-env/goland.md)

---

## Visão Geral do Projeto

O **JusBrasil Tech Challenge** é um sistema desenvolvido para realizar scraping de dados legais e disponibilizá-los via uma API REST.  
Principais tecnologias utilizadas:
- **Linguagem**: Go (Golang)
- **Frameworks**: Uber FX, Go-Chi
- **Banco de Dados**: MySQL
- **Documentação da API**: Swagger UI e OpenAPI

---

## Configuração do Ambiente

Este guia descreve como configurar o ambiente para rodar o projeto.

### Pré-requisitos
- **Go**: Versão 1.20 ou superior.
- **Docker** e **Docker Compose**.
- **MySQL**: Para o banco de dados.

### Passos para Configuração
1. **Clone o Repositório**
   ```bash
   git clone https://github.com/seu-usuario/jusbrasil-tech-challenge.git
   cd jusbrasil-tech-challenge

2. **Configure as Variáveis de Ambiente**

    Crie um arquivo `.env` na raiz do projeto com o seguinte

    ```
    DB_USER=root
    DB_PASSWORD=password
    DB_NAME=jusbrasil
    DB_HOST=localhost
    DB_PORT=3306
   ```
3. **Suba os Serviços do Docker**

    Execute o comando `make fs` para subir os serviços do Docker


4.  **Instale as Dependências do Projeto**

    `go mod tidy`

5. **Teste a Configuração**

    * Acesse a documentação do Swagger da API:

       `make specs_serve`
    * Abra o navegador em: [http//localhost:8080](http://localhost:8080/)

## Uso do Makefile

O `Makefile` simplifica o gerenciamento do projeto. Aqui estão os comandos mais importantes:

### Principais Comandos

1. **Executar a aplicação**
   
    `make run_local`


2. **Gerar a Documentação Swagger**

    `make specs_generate`


3. **Servir a Documentação Swagger**

    `make specs_serve`


4. **Executar testes**

    `make test`


5. **Listar todos os comandos disponíveis**

    `make help`

## Documentação da API

A API fornece endpoints para realizar scraping e consultar dados legais.

### Endpoints Disponíveis

1. **Saúde da aplicação**

    * **Método:** `GET`
    * **URL:** `/health`
    * **Resposta:**
   ```
    {
        "status": "OK"
    }
   ```

2. **Consulta de Dados Extraídos**

   * **Método:** `GET`
   * **URL:** `/health`
   * **Parâmetros**:
     * `url` (string, obrigatório): 
     
       URL do site para scrapping: 
     ```
     https://storage.googleapis.com/jus-challenges/challenge-crawler.html
     * ``` 
   * **Exemplo:**
   ```
    curl -X GET "http://localhost:8080/scrapper?url=https://storage.googleapis.com/jus-challenges/challenge-crawler.html"
   ```
   * **Resposta:**

   ```
      [
       {
        "caseNumber": "1234567-89.2023.8.00.0000",
        "summary": "Resumo do caso",
        "reporter": "Nome do relator",
        "court": "Tribunal X",
        "judgingBody": "Órgão julgador Y",
        "judgementDate": "2024-11-27",
        "caseClass": "Classe do caso",
        "publicationDate": "2024-11-28"
       }
      ]
   ```
   
## Links Úteis
   * [Swagger UI - Documentação Interativa da API](http://localhost:8080/) (Após rodar make specs_serve)
   * [Repositório no GitHub](https://github.com/seu-usuario/jusbrasil-tech-challenge)