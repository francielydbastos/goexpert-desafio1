# 💱 Client-Server-API - Consulta de Cotação USD/BRL

Esse projeto foi desenvolvido como parte do desafio proposto pelo curso Go Expert, da FullCycle.

Essa aplicação é dividida em duas partes principais: 
- Um server que disponibiliza o endpoint /cotacao na porta :8080 e consulta a cotação USD/BRL na API economia.awesomeapi. Os dados são persistidos em banco de dados SQLite utilizando GORM.
- Um client que consome esse server e salva os dados em um arquivo .txt

As duas partes estão contidas nesse repositório, porém são executadas separadamente.

------------------------------------------------------------------------

## 🧠 Tecnologias Utilizadas

-   Go
-   GORM
-   SQLite (driver modernc - pure Go)
-   Context API (timeouts controlados)

------------------------------------------------------------------------

## ⚙️ Como Executar o Projeto

### 1️⃣ Clonar o repositório

``` bash
git clone <url-do-repositorio>
cd Client-Server-API
```

### 2️⃣ Instalar dependências

``` bash
go mod tidy
```

### 3️⃣ Garantir que CGO está desabilitado

Windows (PowerShell):

``` powershell
$env:CGO_ENABLED="0"
```

Linux/Mac:

``` bash
export CGO_ENABLED=0
```

Essa verificação é necessária devido à forma como o GORM executa o SQLite utilizando a linguagem C, o que pode gerar erros na inicialização do banco de dados.

### 4️⃣ Executar o servidor

``` bash
go run ./server
```

------------------------------------------------------------------------

## 🗃️ Banco de Dados

O SQLite cria automaticamente o arquivo na raiz do projeto:

    goexpert.db

Os dados inseridos podem ser visualizados utilizando o sqlite3 ou um plugin de database na IDE.

------------------------------------------------------------------------

## ⏱️ Controle de Timeout

O projeto utiliza contextos com timeout:

-   200ms → chamada à API externa
-   10ms → inserção no banco

------------------------------------------------------------------------

## 📌 Observações

