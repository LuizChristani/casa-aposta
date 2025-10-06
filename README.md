# 🎰 Casa Aposta

Uma aplicação backend simples desenvolvida em Go, utilizando arquivos JSON para armazenamento de dados e o framework Gin para API REST.

## 📋 Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Estado Atual](#estado-atual)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Instalação](#instalação)
- [Como Usar](#como-usar)
- [API Disponível](#api-disponível)
- [Modelos de Dados](#modelos-de-dados)
- [Arquitetura](#arquitetura)
- [Licença](#-licença)

## 🎯 Sobre o Projeto

Casa Aposta é um projeto backend **em desenvolvimento inicial** para gerenciamento de jogos. Atualmente implementa apenas funcionalidades básicas de listagem de jogos através de uma API REST simples.

### ⚠️ Estado Atual - Funcionalidades Implementadas

- ✅ Sistema básico de banco de dados JSON
- ✅ Modelo de dados para Games e User
- ✅ API REST com endpoint GET /games
- ✅ Estrutura de projeto organizada
- ✅ Configuração com Gin framework

### 🚧 Em Desenvolvimento/Planejado

- ❌ Sistema de apostas (não implementado)
- ❌ Autenticação de usuários (não implementado)
- ❌ CRUD completo (apenas leitura implementada)
- ❌ Lógica de negócio (service layer vazio)
- ❌ Validações de dados
- ❌ Testes automatizados

## 📁 Estrutura do Projeto

```
casa-aposta/
├── cmd/
│   └── api.go                    # Ponto de entrada - servidor HTTP
├── database/
│   ├── configuration_db.go      # Configuração do banco
│   ├── database.go              # Operações básicas
│   ├── connections/
│   │   ├── connect.go           # Conexão com "banco"
│   │   └── status.go            # Status da conexão
│   ├── operations/
│   │   ├── insert.go            # Operação de inserção
│   │   └── select.go            # Operação de seleção
│   └── tables/
│       ├── Games.json           # Dados dos jogos (vazio: [])
│       └── User.json            # Dados dos usuários (vazio: [])
├── models/
│   ├── all_models.go            # Lista de modelos
│   ├── games.go                 # Modelo Games
│   └── user.go                  # Modelo User
├── repository/
│   └── games.go                 # Repository para games (básico)
├── service/
│   └── games.go                 # Service layer (vazio)
├── runtime/
│   └── runtime.go               # Verificações de runtime
├── go.mod                       # Dependências Go
├── go.sum                       # Lock de dependências
└── LICENSE                      # Licença restritiva
```

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
[GIN-debug] GET    /games                    --> main.main.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

### Testando a API

```bash
# Listar jogos (retorna array vazio por padrão)
curl http://localhost:8080/games
```

**Resposta:**
```json
{
  "games": []
}
```

## 🔧 API Disponível

### Endpoints Implementados

| Método | Endpoint | Descrição | Status |
|--------|----------|-----------|--------|
| GET | `/games` | Lista todos os jogos | ✅ Funcionando |

### Endpoints Planejados (Não Implementados)

- `POST /games` - Criar jogo
- `GET /games/:id` - Obter jogo específico
- `PUT /games/:id` - Atualizar jogo
- `DELETE /games/:id` - Deletar jogo
- `GET /users` - Listar usuários
- `POST /users` - Criar usuário
- Sistema de autenticação
- Sistema de apostas

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

### Arquitetura Atual (Simples)

```
HTTP Request → Gin Router → Repository → JSON Files
```

### Camadas Implementadas

- **CMD**: Servidor HTTP com Gin
- **Repository**: Acesso básico aos dados (apenas leitura)
- **Models**: Estruturas de dados
- **Database**: Sistema de arquivos JSON

### Camadas Planejadas (Não Implementadas)

- **Service**: Lógica de negócio (arquivo vazio)
- **Middleware**: Autenticação, validação
- **Tests**: Testes automatizados

## 💾 Sistema de "Banco de Dados"

O projeto utiliza arquivos JSON como banco de dados:

- **Games.json**: Armazena dados dos jogos (atualmente vazio: `[]`)
- **User.json**: Armazena dados dos usuários (atualmente vazio: `[]`)

### Limitações Atuais

- Não há persistência real de dados
- Não há validações
- Não há transações
- Performance limitada para grandes volumes
- Não há índices ou otimizações

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

## 🚧 Limitações e TODOs

### Limitações Conhecidas

1. **Dados não persistem** entre reinicializações
2. **Sem validação** de entrada
3. **Sem autenticação** ou autorização
4. **Sem testes** automatizados
5. **Service layer vazio** - sem lógica de negócio
6. **Apenas leitura** - sem operações de escrita via API

### Próximos Passos Sugeridos

1. Implementar CRUD completo para Games
2. Adicionar validações de dados
3. Implementar persistência real dos dados
4. Criar testes automatizados
5. Adicionar sistema de usuários
6. Implementar autenticação
7. Desenvolver lógica de apostas

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