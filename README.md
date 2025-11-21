# vehicle-platform-payments
Este repositório contém um serviço de simulação de um gateway de pagamentos.

## Funcionalidades

- **Simula um gateway de pagamento:** Realiza a simulação de um gateway de pagamento retornando algumas informações sobre o pagamento realizado.

## Tecnologias Utilizadas

- **Go (Golang):** Para o desenvolvimento da API de vendas.
- **Gin:** Framework web para o desenvolvimento da API.
- **Docker Compose:** Para o setup do serviço via Docker.

## Como Rodar o Projeto Localmente

### 1. Pré-requisitos

Certifique-se de que você tem as seguintes dependências instaladas:

- **Go (Golang)** versão 1.18 ou superior
- **Git** para clonar o repositório
- **Docker** e **Docker Compose**

### 2. Configuração para rodar o serviço localmente com Docker Compose

1. Clone o repositório:

    ```bash
    git clone git@github.com:caiiomp/vehicle-platform-payments.git
    ```

2. Na raiz do projeto instale as dependências do Go:

    ```bash
    go mod tidy
    ```

3. Na raiz do projeto, inicie o serviço e suas dependências `docker`:

    ```bash
    docker compose up -d
    ```

    Isso irá iniciar o serviço e as suas dependências localmente via contêiner. O serviço estará disponível em `http://localhost:4003`.

    ⚠️ Para que consigamos rodar todos os serviços integrados, devemos criar uma rede compartilhada no docker. Caso não tenha criada, podemos criar com o seguinte comando:

    ```bash
    docker network create shared_network
    ```

### 3. Testando o serviço

Use **Postman**, **Insomnia**, **cURL** ou qualquer outro cliente **HTTP** para testar os endpoints:

- `POST /payments` - Simula um pagamento

Os testes unitários e os testes de integração podem ser executados da seguinte forma respectivamente:
```bash
    go test ./... -v
    go test -tags=integration -v ./...
```

## Documentação (Swagger)

Para acessar a documentação do serviço, acessar o seguinte endpoint: 
```
http://localhost:4003/swagger/index.html
```