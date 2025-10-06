# 🔢 Enums em Go - Casa Aposta

Este documento explica como implementar e usar enums em Go no projeto Casa Aposta.

## 📋 Índice

- [O que são Enums](#o-que-são-enums)
- [Tipos de Enums em Go](#tipos-de-enums-em-go)
- [Implementações no Projeto](#implementações-no-projeto)
- [Exemplos de Uso](#exemplos-de-uso)
- [Boas Práticas](#boas-práticas)

## 🤔 O que são Enums

Enums (enumerações) são um tipo de dados que consiste em um conjunto de constantes nomeadas. Em Go, não existe um tipo `enum` nativo, mas podemos simular enums usando constantes e tipos customizados.

## 🛠️ Tipos de Enums em Go

### 1. Enum com `iota` (Números Sequenciais)

```go
type ConnectionStatus int

const (
    Disconnected ConnectionStatus = iota  // 0
    Connected                             // 1
    Connecting                            // 2
    Error                                 // 3
)
```

**Vantagens:**
- Valores únicos automáticos
- Eficiente em memória
- Comparações rápidas

**Desvantagens:**
- Valores podem mudar se a ordem for alterada
- Não são legíveis em logs sem método String()

### 2. Enum com Strings

```go
type DatabaseOperation string

const (
    OperationConnect    DatabaseOperation = "Connect"
    OperationDisconnect DatabaseOperation = "Disconnect"
    OperationQuery      DatabaseOperation = "Query"
)
```

**Vantagens:**
- Valores legíveis
- Estáveis (não mudam com reordenação)
- Fáceis de debugar

**Desvantagens:**
- Maior uso de memória
- Comparações mais lentas

## 🏗️ Implementações no Projeto

### ConnectionStatus (Enum com iota)

```go
// Definição do tipo
type ConnectionStatus int

// Constantes do enum
const (
    Disconnected ConnectionStatus = iota
    Connected
    Connecting
    Error
)

// Método String para legibilidade
func (cs ConnectionStatus) String() string {
    switch cs {
    case Disconnected:
        return "Disconnected"
    case Connected:
        return "Connected"
    case Connecting:
        return "Connecting"
    case Error:
        return "Error"
    default:
        return "Unknown"
    }
}
```

### DatabaseOperation (Enum com Strings)

```go
type DatabaseOperation string

const (
    OperationConnect    DatabaseOperation = "Connect"
    OperationDisconnect DatabaseOperation = "Disconnect"
    OperationQuery      DatabaseOperation = "Query"
    OperationInsert     DatabaseOperation = "Insert"
    OperationUpdate     DatabaseOperation = "Update"
    OperationDelete     DatabaseOperation = "Delete"
)
```

## 💻 Exemplos de Uso

### Uso Básico

```go
// Declarando variáveis
var status ConnectionStatus = Connected
var operation DatabaseOperation = OperationQuery

// Comparações
if status == Connected {
    fmt.Println("Banco conectado!")
}

if operation == OperationQuery {
    fmt.Println("Executando consulta...")
}
```

### Em Funções

```go
func processOperation(op DatabaseOperation) error {
    switch op {
    case OperationConnect:
        return connectToDatabase()
    case OperationDisconnect:
        return disconnectFromDatabase()
    case OperationQuery:
        return executeQuery()
    default:
        return errors.New("operação não suportada")
    }
}

// Uso
err := processOperation(OperationConnect)
```

### Com Switch Statements

```go
func getStatusMessage(status ConnectionStatus) string {
    switch status {
    case Disconnected:
        return "Banco de dados desconectado"
    case Connected:
        return "Banco de dados conectado"
    case Connecting:
        return "Conectando ao banco de dados..."
    case Error:
        return "Erro na conexão com o banco"
    default:
        return "Status desconhecido"
    }
}
```

### Validação de Enums

```go
func isValidOperation(op DatabaseOperation) bool {
    switch op {
    case OperationConnect, OperationDisconnect, OperationQuery,
         OperationInsert, OperationUpdate, OperationDelete:
        return true
    default:
        return false
    }
}

func isValidStatus(status ConnectionStatus) bool {
    return status >= Disconnected && status <= Error
}
```

## 🎯 Exemplos Práticos no Projeto

### 1. Gerenciamento de Status de Conexão

```go
package main

import (
    "fmt"
    "casa-aposta/database"
)

func main() {
    // Status inicial
    fmt.Printf("Status inicial: %s\n", database.CurrentConnectionStatus)
    
    // Simulando conexão
    database.CurrentConnectionStatus = database.Connecting
    fmt.Printf("Conectando: %s\n", database.CurrentConnectionStatus)
    
    // Conexão estabelecida
    database.CurrentConnectionStatus = database.Connected
    fmt.Printf("Conectado: %s\n", database.CurrentConnectionStatus)
}
```

### 2. Sistema de Logs com Enums

```go
func logDatabaseOperation(op database.DatabaseOperation, status database.ConnectionStatus) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    fmt.Printf("[%s] Operação: %s | Status: %s\n", timestamp, op, status)
}

// Uso
logDatabaseOperation(database.OperationConnect, database.Connecting)
logDatabaseOperation(database.OperationQuery, database.Connected)
```

### 3. Configuração Baseada em Enums

```go
type DatabaseConfig struct {
    Operation DatabaseOperation
    Status    ConnectionStatus
    Retries   int
}

func createConfig(op DatabaseOperation) DatabaseConfig {
    return DatabaseConfig{
        Operation: op,
        Status:    database.Disconnected,
        Retries:   3,
    }
}
```

## 📝 Boas Práticas

### 1. Sempre Implemente o Método String()

```go
func (cs ConnectionStatus) String() string {
    switch cs {
    case Disconnected:
        return "Disconnected"
    // ... outros casos
    default:
        return "Unknown"
    }
}
```

### 2. Use Tipos Customizados

```go
// ✅ Bom - tipo customizado
type Status int

// ❌ Ruim - apenas constantes
const (
    StatusOK = 1
    StatusError = 2
)
```

### 3. Agrupe Enums Relacionados

```go
// ✅ Bom - enums relacionados juntos
type DatabaseStatus int
const (
    DBDisconnected DatabaseStatus = iota
    DBConnected
    DBError
)

type DatabaseOperation string
const (
    DBConnect DatabaseOperation = "connect"
    DBQuery   DatabaseOperation = "query"
)
```

### 4. Use Validação

```go
func SetConnectionStatus(status ConnectionStatus) error {
    if status < Disconnected || status > Error {
        return errors.New("status de conexão inválido")
    }
    CurrentConnectionStatus = status
    return nil
}
```

### 5. Documente os Enums

```go
// ConnectionStatus representa o estado atual da conexão com o banco de dados
type ConnectionStatus int

const (
    Disconnected ConnectionStatus = iota // Banco desconectado
    Connected                            // Banco conectado e pronto
    Connecting                           // Estabelecendo conexão
    Error                                // Erro na conexão
)
```

## 🧪 Testando Enums

```go
func TestConnectionStatus(t *testing.T) {
    // Teste de valores
    if Disconnected != 0 {
        t.Errorf("Disconnected deveria ser 0, mas é %d", Disconnected)
    }
    
    // Teste do método String
    if Connected.String() != "Connected" {
        t.Errorf("String() deveria retornar 'Connected'")
    }
    
    // Teste de validação
    if !isValidStatus(Connected) {
        t.Errorf("Connected deveria ser um status válido")
    }
}
```

## 🔄 Migração de Constantes para Enums

Se você tem constantes simples:

```go
// Antes
const (
    CONNECT = "connect"
    DISCONNECT = "disconnect"
)
```

Migre para enums:

```go
// Depois
type Operation string

const (
    OperationConnect    Operation = "connect"
    OperationDisconnect Operation = "disconnect"
)
```

## 📊 Comparação de Performance

| Tipo | Memória | Velocidade | Legibilidade |
|------|---------|------------|--------------|
| iota | Baixa   | Alta       | Média        |
| string | Alta  | Média      | Alta         |

## 🚀 Executando os Exemplos

```bash
# Execute o exemplo de uso dos enums
go run examples/enum_usage.go

# Saída esperada:
# === Exemplos de Uso dos Enums ===
# 
# 1. Enum ConnectionStatus:
# Status 1: Disconnected (valor: 0)
# Status 2: Connected (valor: 1)
# ...
```

---

Os enums tornam o código mais legível, type-safe e fácil de manter. Use-os sempre que tiver um conjunto fixo de valores relacionados!