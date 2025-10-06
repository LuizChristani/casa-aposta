# 🎰 Casa Aposta

Uma aplicação backend robusta desenvolvida em Go, utilizando arquivos JSON para armazenamento de dados e o framework Gin para API REST, com arquitetura em camadas bem definida.

## 📋 Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Estado Atual](#estado-atual)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Instalação](#instalação)
- [Como Usar](#como-usar)
- [API Disponível](#api-disponível)
- [Modelos de Dados](#modelos-de-dados)
- [Arquitetura](#arquitetura)
- [Sistema de IDs](#sistema-de-ids)
- [Licença](#-licença)

## 🎯 Sobre o Projeto

Casa Aposta é um projeto backend para gerenciamento de jogos com **arquitetura em camadas completa**. Implementa CRUD funcional para jogos, sistema de validações, geração automática de IDs e uma API REST robusta.

### ✅ Estado Atual - Funcionalidades Implementadas

- ✅ **Arquitetura em camadas completa** (Repository → Service → Handler)
- ✅ **CRUD completo para Games** (Create, Read, Update, Delete)
- ✅ **Sistema de banco de dados JSON** com persistência
- ✅ **Geração automática de IDs** únicos
- ✅ **Validações de negócio** (nomes únicos, valores válidos)
- ✅ **API REST funcional** com endpoints GET e POST
- ✅ **Injeção de dependências** adequada
- ✅ **Modelos de dados estruturados**
- ✅ **Sistema de operações genéricas** (Insert, Select)

### 🚧 Em Desenvolvimento/Planejado

- 🔄 Endpoints PUT e DELETE para Games
- 🔄 Sistema completo de usuários
- 🔄 Autenticação e autorização
- 🔄 Sistema de apostas
- 🔄 Middleware de validação
- 🔄 Testes automatizados
- 🔄 Documentação da API (Swagger)

## 📁 Estrutura do Projeto

```
casa-aposta/
├── cmd/
│   ├── api.go                    # Servidor HTTP principal com injeção de dependências
│   └── endpoints/
│       └── handler.go            # Handlers HTTP (Controllers)
├── contract/
│   └── games/
│       └── games_request.go      # Contratos de requisição
├── database/
│   ├── configuration_db.go      # Configuração do banco
│   ├── database.go              # Operações básicas
│   ├── connections/
│   │   ├── connect.go           # Conexão com "banco"
│   │   └── status.go            # Status da conexão
│   ├── operations/
│   │   ├── generateId.go        # ✨ Geração e consulta de IDs
│   │   ├── insert.go            # Operações de inserção (genéricas)
│   │   └── select.go            # Operações de seleção
│   └── tables/
│       ├── Games.json           # Dados dos jogos (com dados reais)
│       └── User.json            # Dados dos usuários
├── examples/
│   └── id_usage.go              # ✨ Exemplos de uso das funções de ID
├── models/
│   ├── all_models.go            # Lista de modelos
│   ├── games.go                 # Modelo Games completo
│   └── user.go                  # Modelo User
├── repository/
│   └── games.go                 # ✨ Repository completo com CRUD
├── service/
│   └── games.go                 # ✨ Service layer com validações de negócio
├── validator/
│   └── validator.go             # Sistema de validações
├── go.mod                       # Dependências Go
├── go.sum                       # Lock de dependências
└── LICENSE                      # Licença restritiva
```

### 🆕 Novos Arquivos e Funcionalidades

- **✨ generateId.go**: Sistema completo de geração e consulta de IDs
- **✨ examples/**: Pasta com exemplos práticos de uso
- **✨ endpoints/**: Handlers organizados em pasta separada
- **✨ contract/**: Contratos de API para validação
- **🔄 Repository**: Agora com CRUD completo implementado
- **🔄 Service**: Camada de negócio com validações funcionais
- **🔄 Handler**: Controllers HTTP com injeção de dependências

## 🚀 Instalação

### Pré-requisitos

- Go 1.25.1 ou superior
- Git

### Passos de Instalação

1. **Clone o repositório**
   ```bash
   git clone <url-do-repositorio>
   cd casa-aposta
   ```

2. **Instale as dependências**
   ```bash
   go mod tidy
   ```

3. **Execute a aplicação**
   ```bash
   go run cmd/api.go
   ```

## 🎮 Como Usar

### Executando o Servidor

```bash
go run cmd/api.go
```

### Saída Esperada

```
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
database: get table name User
database: table User exists
database: get table name Games
database: table Games exists
[GIN-debug] GET    /games                    --> casa-aposta/cmd/endpoints.(*Handler).GetAllGames-fm (3 handlers)
[GIN-debug] POST   /games                    --> casa-aposta/cmd/endpoints.(*Handler).CreateGame-fm (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

### Testando a API

#### 📋 Listar Jogos (GET)
```bash
# PowerShell
Invoke-WebRequest -Uri "http://localhost:8080/games" -Method GET

# Curl (se disponível)
curl http://localhost:8080/games
```

**Resposta:**
```json
[
  {
    "id": 1,
    "name": "Mega Sena",
    "description": "Loteria com 6 números",
    "min_bet": 4.5,
    "max_bet": 500
  }
]
```

#### ➕ Criar Jogo (POST)
```bash
# PowerShell
Invoke-WebRequest -Uri "http://localhost:8080/games" -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"name":"Lotofácil","description":"Loteria com 15 números","min_bet":2.50,"max_bet":300.0}'

# Curl (se disponível)
curl -X POST http://localhost:8080/games \
  -H "Content-Type: application/json" \
  -d '{"name":"Lotofácil","description":"Loteria com 15 números","min_bet":2.50,"max_bet":300.0}'
```

**Resposta:**
```json
{
  "id": 2,
  "name": "Lotofácil",
  "description": "Loteria com 15 números",
  "min_bet": 2.5,
  "max_bet": 300
}
```

## 🔧 API Disponível

### ✅ Endpoints Implementados

| Método | Endpoint | Descrição | Status | Validações |
|--------|----------|-----------|--------|------------|
| GET | `/games` | Lista todos os jogos | ✅ Funcionando | - |
| POST | `/games` | Cria um novo jogo | ✅ Funcionando | Nome único, valores válidos |

### 🚧 Endpoints Planejados

- `GET /games/:id` - Obter jogo específico
- `PUT /games/:id` - Atualizar jogo
- `DELETE /games/:id` - Deletar jogo
- `GET /users` - Listar usuários
- `POST /users` - Criar usuário
- Sistema de autenticação
- Sistema de apostas

### 🔒 Validações Implementadas

#### POST /games
- ✅ **Nome obrigatório** e não pode estar vazio
- ✅ **Nome único** - não permite jogos com nomes duplicados
- ✅ **min_bet** não pode ser negativo
- ✅ **max_bet** não pode ser negativo
- ✅ **min_bet** não pode ser maior que **max_bet**
- ✅ **ID gerado automaticamente** de forma sequencial

## 📊 Modelos de Dados

### Games

```go
type Games struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    MinBet      float64 `json:"min_bet"`
    MaxBet      float64 `json:"max_bet"`
}
```

### User

```go
type User struct {
    ID       int
    Name     string
    Password string
    Saldo    float32
}
```

**Nota:** O modelo User não possui tags JSON implementadas.

## 🏗️ Arquitetura

### Arquitetura Atual (Camadas Completas)

```
HTTP Request → Gin Router → Handler → Service → Repository → JSON Files
     ↓              ↓           ↓         ↓          ↓
  JSON Body    Route Mapping  Validation  Business   Data Access
                              & Binding    Logic     & Persistence
```

### 🎯 Camadas Implementadas

#### **1. Handler Layer (Controllers)**
- **Localização**: `cmd/endpoints/handler.go`
- **Responsabilidade**: Receber requisições HTTP, validar entrada, chamar services
- **Funcionalidades**: 
  - Binding de JSON para structs
  - Validação de entrada
  - Tratamento de erros HTTP
  - Injeção de dependências

#### **2. Service Layer (Business Logic)**
- **Localização**: `service/games.go`
- **Responsabilidade**: Lógica de negócio, validações de domínio
- **Funcionalidades**:
  - Validação de regras de negócio
  - Verificação de duplicatas
  - Validação de valores (min_bet, max_bet)
  - Orquestração de operações

#### **3. Repository Layer (Data Access)**
- **Localização**: `repository/games.go`
- **Responsabilidade**: Acesso aos dados, operações CRUD
- **Funcionalidades**:
  - Operações de leitura e escrita
  - Abstração do sistema de persistência
  - Mapeamento de dados

#### **4. Database Layer (Persistence)**
- **Localização**: `database/operations/`
- **Responsabilidade**: Operações de baixo nível com arquivos
- **Funcionalidades**:
  - Geração de IDs únicos
  - Operações genéricas de Insert/Select
  - Manipulação de arquivos JSON

### 🔄 Fluxo de Dados

#### **POST /games (Criar Jogo)**
```
1. HTTP Request → Handler.CreateGame()
2. JSON Binding → models.Games struct
3. Handler → Service.CreateGame()
4. Service → Validações de negócio
5. Service → Repository.CreateGame()
6. Repository → operations.InsertModel()
7. Operations → Geração de ID + Persistência
8. Response ← JSON com jogo criado
```

#### **GET /games (Listar Jogos)**
```
1. HTTP Request → Handler.GetAllGames()
2. Handler → Service.GetAllGames()
3. Service → Repository.GetAllGames()
4. Repository → operations.Select()
5. Operations → Leitura do arquivo JSON
6. Response ← JSON com lista de jogos
```

### 🔧 Injeção de Dependências

```go
// cmd/api.go
gamesRepo := repository.NewGamesRepository()
gamesService := service.NewGamesService(gamesRepo)
handler := endpoints.NewHandler(gamesService)

router.GET("/games", handler.GetAllGames)
router.POST("/games", handler.CreateGame)
```

## 🆔 Sistema de IDs

O projeto implementa um sistema robusto de geração e consulta de IDs únicos:

### 📋 Funções Disponíveis

#### **`GenerateID(tableName string) int64`**
- **Propósito**: Gera o próximo ID sequencial para uma tabela
- **Uso**: Criação de novos registros
- **Exemplo**: `nextID := operations.GenerateID("Games")` → `2`

#### **`GetLastID(tableName string) int64`**
- **Propósito**: Retorna o maior ID existente na tabela
- **Uso**: Verificar o último ID usado (independente da ordem)
- **Exemplo**: `lastID := operations.GetLastID("Games")` → `1`

#### **`GetLastCreatedID(tableName string) int64`**
- **Propósito**: Retorna o ID do último registro inserido (último na lista)
- **Uso**: Verificar o último registro criado cronologicamente
- **Exemplo**: `lastCreated := operations.GetLastCreatedID("Games")` → `1`

### 💡 Exemplo Prático

```go
import "casa-aposta/database/operations"

// Verificar último ID criado
lastID := operations.GetLastCreatedID("Games")
fmt.Printf("Último jogo criado: ID %d", lastID)

// Gerar próximo ID
nextID := operations.GenerateID("Games")
fmt.Printf("Próximo ID será: %d", nextID)
```

### 🔧 Como Funciona

1. **Leitura**: Carrega dados existentes do arquivo JSON
2. **Análise**: Percorre registros para encontrar IDs
3. **Cálculo**: Determina o próximo ID baseado na lógica escolhida
4. **Retorno**: Fornece o ID calculado

## 💾 Sistema de "Banco de Dados"

O projeto utiliza arquivos JSON como banco de dados com persistência real:

- **Games.json**: Armazena dados dos jogos (com dados persistentes)
- **User.json**: Armazena dados dos usuários

### ✅ Funcionalidades Implementadas

- ✅ **Persistência real** de dados entre reinicializações
- ✅ **Geração automática de IDs** únicos e sequenciais
- ✅ **Operações CRUD** funcionais
- ✅ **Validações** de integridade de dados
- ✅ **Sistema genérico** de inserção e consulta

### ⚠️ Limitações Conhecidas

- Performance limitada para grandes volumes
- Não há transações ACID
- Não há índices ou otimizações
- Concorrência limitada (single-threaded)
- Sem backup automático

## 🔧 Dependências

### Principais

- **Gin**: Framework web para Go
- **Go 1.25.1**: Versão do Go utilizada

### Completas (go.mod)

```
github.com/gin-gonic/gin v1.11.0
github.com/go-playground/validator/v10 v10.28.0
github.com/goccy/go-json v0.10.5
// ... outras dependências transitivas
```

## 🚧 Status do Desenvolvimento

### ✅ Funcionalidades Completadas

1. ✅ **Arquitetura em camadas** - Repository → Service → Handler
2. ✅ **CRUD para Games** - Create e Read implementados
3. ✅ **Persistência de dados** - dados salvos entre reinicializações
4. ✅ **Validações de negócio** - regras de domínio implementadas
5. ✅ **Sistema de IDs** - geração automática e consultas
6. ✅ **API REST funcional** - endpoints GET e POST
7. ✅ **Injeção de dependências** - arquitetura desacoplada

### 🔄 Em Desenvolvimento

1. 🔄 **Endpoints PUT e DELETE** para Games
2. 🔄 **Sistema completo de usuários** (User CRUD)
3. 🔄 **Middleware de validação** avançado
4. 🔄 **Tratamento de erros** mais robusto

### 📋 Próximos Passos Planejados

1. **Completar CRUD de Games**
   - Implementar `GET /games/:id`
   - Implementar `PUT /games/:id`
   - Implementar `DELETE /games/:id`

2. **Sistema de Usuários**
   - Criar endpoints para User
   - Implementar validações específicas
   - Adicionar relacionamentos

3. **Melhorias de Qualidade**
   - Adicionar testes automatizados
   - Implementar logging estruturado
   - Criar documentação da API (Swagger)

4. **Funcionalidades Avançadas**
   - Sistema de autenticação JWT
   - Sistema de apostas
   - Middleware de autorização
   - Rate limiting

### ⚠️ Limitações Atuais

- **Concorrência**: Sistema single-threaded
- **Performance**: Não otimizado para grandes volumes
- **Segurança**: Sem autenticação implementada
- **Testes**: Sem cobertura de testes automatizados
- **Monitoramento**: Sem métricas ou observabilidade

## 📄 Licença

**LICENÇA RESTRITIVA - USO NÃO COMERCIAL**

Copyright (c) 2024 Casa Aposta

### Termos e Condições

Este software e a documentação associada (o "Software") são fornecidos sob os seguintes termos:

#### ✅ **PERMITIDO:**
- ✅ Uso pessoal e educacional
- ✅ Modificação para uso próprio
- ✅ Distribuição para fins não comerciais
- ✅ Estudo e aprendizado
- ✅ Contribuições para o projeto original

#### ❌ **PROIBIDO:**
- ❌ **USO COMERCIAL** de qualquer natureza
- ❌ Venda do software ou de versões modificadas
- ❌ Uso em produtos ou serviços comerciais
- ❌ Monetização direta ou indireta
- ❌ Licenciamento para terceiros com fins lucrativos

#### 📋 **CONDIÇÕES:**
- O aviso de copyright deve ser mantido em todas as cópias
- Esta licença deve ser incluída em todas as distribuições
- Modificações devem ser claramente indicadas
- O software é fornecido "COMO ESTÁ", sem garantias

#### ⚖️ **VIOLAÇÕES:**
O uso comercial não autorizado resultará em ação legal imediata.

Para uso comercial, entre em contato para licenciamento especial.

---

**Casa Aposta** - Projeto em desenvolvimento inicial 🚧