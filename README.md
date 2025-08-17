# 📇 API REST de Contatos

Uma API REST em Go com arquitetura limpa para gerenciar contatos usando o framework [Fiber](https://gofiber.io/) e SQLite como banco de dados.  
Inclui funcionalidades completas de CRUD, validação de CPF/CNPJ (padrões brasileiros), formatação de número de telefone, sanitização contra XSS e SQL injection.

---

## ✨ Funcionalidades
- **Arquitetura Limpa** (separação de camadas entre handlers, use cases e repositórios)
- **Banco SQLite** (embutido, sem necessidade de configuração)
- **Framework Fiber** para alta performance em HTTP
- **CRUD de Contatos**
- **Validação de CPF & CNPJ** (algoritmo oficial brasileiro)
- **Formatação de telefone** para `(XX) X XXXX-XXXX`
- **Sanitização SQL** em todas as entradas do usuário
- **Prevenção XSS** em strings

---

## 🗂 Estrutura do Projeto
```
.
├── cmd
│   └── main.go               # Ponto de entrada da aplicação
├── internal
│   ├── entity                # Entidades (modelos de domínio)
│   ├── repository            # Persistência de dados
│   ├── usecase               # Lógica de negócio
│   └── handler               # Handlers HTTP
├── pkg
│   ├── database              # Conexão ao banco dados
│   ├── formatters            # Formatação de telefone, etc.
│   ├── middlwares            # Aplica sanitização de json
│   ├── sanitizer             # Sanitização SQL & XSS
│   ├── validators            # Validação de CPF/CNPJ
│   └── seed                  # Seed de contatos de exemplo
├── Dockerfile
├── docker-compose.yml        # Containers de seed e aplicação
├── go.mod
├── go.sum
└── README.md
```

---

## 📦 Requisitos
- [Go 1.21+](https://go.dev/dl/)
- Docker e Compose
- GCC

---

## 🚀 Executando Localmente

### Clonar o repositório
```bash
git clone https://github.com/bergsantana/go-contacts.git 
cd go-contacts
```

 

###  Subir o containers
Executa seed -> go-api
```bash
docker compose up -d
```
 

Você verá um **log de inicialização** com todos os endpoints disponíveis.

---

## 📡 Endpoints da API

| Método | Endpoint                  | Descrição                     |
|--------|---------------------------|-------------------------------|
| GET    | `/contacts`               | Listar todos os contatos      |
| GET    | `/contacts/:id`           | Obter um contato por ID       |
| GET    | `/contacts/cpf/:cpf`      | Obter um contato por CPF      |
| GET    | `/contacts/cnpj/:cnpj`    | Obter um contato por CNPJ     |
| GET    | `/contacts/email/:email`  | Obter um contato por Email    |
| POST   | `/contacts`               | Criar um novo contato         |
| PUT    | `/contacts/:id`           | Atualizar um contato          |
| DELETE | `/contacts/:id`           | Deletar um contato            |

---

## 📝 Exemplos de JSON para criar um contato

**Somente CPF**
```json
{
  "name": "João da Silva",
  "email": "joao.silva@example.com",
  "cpf": "123.456.789-09",
  "cnpj": null,
  "phone": "11987654321",
  "address": "Rua das Flores, 123, São Paulo - SP"
}
```

**Somente CNPJ**
```json
{
  "name": "Empresa XYZ Ltda",
  "email": "contato@empresa.xyz",
  "cpf": null,
  "cnpj": "12.345.678/0001-99",
  "phone": "21912345678",
  "address": "Avenida Paulista, 1500, São Paulo - SP"
}
```

---

## 🔹 Exemplos de GET

**Por CPF**
```
GET /contacts/cpf/12345678909
```
**Resposta**
```json
{
  "name": "João Pereira",
  "email": "joao.pereira@example.com",
  "cpf": "123.456.789-09",
  "cnpj": null,
  "phone": "(11) 9 8765-4321",
  "address": "Rua das Palmeiras, 123, São Paulo - SP"
}
```

**Por CNPJ**
```
GET /contacts/cnpj/45987654000132
```
**Resposta**
```json
{
  "name": "Tech Solutions Ltda",
  "email": "contato@techsolutions.com.br",
  "cpf": null,
  "cnpj": "45.987.654/0001-32",
  "phone": "(11) 3 2567-890",
  "address": "Avenida Paulista, 1500, São Paulo - SP"
}
```

 
---

## 🔒 Recursos de Segurança
- **Proteção contra SQL Injection**: Prepared statements + regex
- **Prevenção XSS**: Tags HTML e scripts removidos das entradas
- **Validações**:
  - CPF segue o algoritmo brasileiro oficial
  - CNPJ segue o algoritmo brasileiro oficial
  - Telefone deve ter 11 dígitos e será formatado
  - Email é único no banco de dados

---

## 🛠 Desenvolvimento

 ```bash
$env:CGO_ENABLED = 1
```

 ```bash
go run cmd/main.go seed
 ```
 ```bash

go run cmd/main.go
```

 

## 📄 Licença
Licença MIT – pode usar o projeto para fins pessoais ou comerciais.
