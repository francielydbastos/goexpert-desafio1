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
git clone https://github.com/francielydbastos/goexpert-desafio1
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

Abra o terminal e execute o comando abaixo:

``` bash
go run ./server
```

O servidor está disponível no endpoint abaixo e aceita requisições GET:

``` bash
http://localhost:8080/cotacao
```

Em caso de sucesso, o endpoint retorna o status 200 OK com a cotação bid no corpo da resposta:

``` bash
{
    "bid": "5.1546"
}
```
Em cenários de erro, o endpoint retorna o status 500 StatusInternalServerError com a descrição do erro no corpo da resposta.

### 5️⃣ Executar o cliente

Abra uma nova janela no terminal e execute o comando abaixo:

``` bash
go run ./client
```

Ao ser executado, o cliente automaticamente fará uma chamada ao endpoint /cotacao e gravará o valor no arquivo cotacao.txt, na raíz do projeto. Ao executar a aplicação pela 1ª vez, caso o arquivo não exista ele será criado; em execuções posteriores o dado será apenas adicionado ao arquivo já existente, mantendo o arquivo dos valores anteriores.

## 🗃️ Banco de Dados

O SQLite cria automaticamente o arquivo na raiz do projeto:

    goexpert.db

Os dados inseridos podem ser visualizados utilizando o sqlite3 ou um plugin de database na IDE.


