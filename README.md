# Billing-Ingestor

O **Billing-Ingestor** é um componente de ingestão de pulsos de uso, desenvolvido em **Go**, responsável por processar e agregar dados de consumo enviados por sistemas externos para fins de **cobrança (billing)**.

---

## Descrição

Este serviço coleta e armazena pulsos de uso enviados via requisições HTTP, realiza agregações diárias e, periodicamente, simula o envio desses dados a um sistema processador e armazenador.

A aplicação foi construída utilizando **arquitetura limpa (Clean Architecture)**, permitindo maior organização, testabilidade e separação clara de responsabilidades entre camadas (handler, usecase, service, repository).

---

## Tecnologias Utilizadas

- **Go**: Linguagem principal
- **Echo**: Framework web para criação da API REST
- **Docker**: Containerização da aplicação
- **Testify**: Testes unitários com mocks
---

## Como rodar o projeto

### 1. Clonar o repositório

Clone o repositório para a sua máquina local:
```bash

git  clone  https://github.com/laianemuckler/billing-ingestor

cd  billing-ingestor
```
### 2. Rodar o projeto localmente

Passo 1: Instalar as dependências
```bash
go mod tidy
```
Passo 2: Executar o projeto
```bash
go run cmd/api/main.go
```
### 3. Subir o projeto com Docker

Passo 1: Construir a imagem do Docker
```bash
docker-compose build
```
Passo 2: Subir os containers
```bash
docker-compose up
```
Passo 3: Reconstruir a imagem após modificações
```bash
docker-compose up --build
```
### 4. Rodar os testes

Rodar todos os testes unitários (de forma detalhada)
```bash
go test -v ./...
```
Rodar os testes com cobertura
```bash
go test -cover ./...
```

## Endpoints

### **GET** `/aggregates` 
Retorna uma lista de pulsos agregados.

### **POST** `/pulses`
Registra um novo pulso de uso.

### **POST** `/agreggates/commit`
Envia manualmente um lote de dados agregados com base em uma data.

### Exemplo de requisição com cURL

Você pode adicionar um pulso para ser processado usando o `curl` com o seguinte comando:

```bash
curl -X POST http://localhost:8081/pulses \
-H "Content-Type: application/json" \
-d '{
  "tenant_id": "tenant1",
  "product_sku": "sku1",
  "use_unit": "unit1",
  "used_amount": 100
}'
```

## Melhorias
- Substituir a fila fake por um serviço de mensageria
- Substituir o cache em memória pelo Redis 
- Adicionar um arquivo de variáveis de  ambiente e uma `config`
- Adicionar testes unitários em todas as camadas e aumentar a cobertura
- Adicionar Actions que permitem rodar todos os testes e fazer deploy para dev
- Adicionar instrumentação para melhoria da Observabilidade