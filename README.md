# Billing-Ingestor
###Contexto:
####Parte 1:
Fazer uso de boas práticas no desenvolvimento de software traz ganhos muito significativos quando pensamos em manutenibilidade, qualidade e evolução contínua de nossas soluções.
Os princípios SOLID e princípios de design funcional, olhando aqui para a os princípios de programação funcional aplicados a forma como estruturamos  um sistema, têm papel fundamental auxiliando na estruturação de sistemas modulares, coesos e resilientes à mudança. Princípios como o da responsabilidade única e o de estar aberto para extensão, mas fechado para modificação ajudam a isolar responsabilidades e facilitam a introdução de novas funcionalidades sem impactar o que já existe. Já a inversão de dependência permite desacoplar regras de negócio das implementações, o que torna o sistema mais flexível a mudanças tecnológicas. Do lado funcional, práticas como o uso de funções mais específicas, com um único objetivo, imutabilidade e composição de funções tornam o código mais previsível, testável e fácil de manter. Isso reduz possíveis efeitos colaterais de  mudanças e aumenta a confiança das equipes para evoluir o sistema.
Ao combinar esses princípios, criamos sistemas que favorecem a integração, extensão e evolução de código, podendo evoluir arquiteturas complexas de forma sustentável.

####Parte 2:
Implementação de uma das partes do sistema de bilhetagem de consumo: ingestor. 

---

## Descrição

O **Billing-Ingestor** é um componente de ingestão desenvolvido em Golang, responsável por processar dados de consumo (pulsos) enviados via requisições HTTP, realizando agregações em lotes e simulando o envio desses dados a um sistema processador e armazenador.

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

