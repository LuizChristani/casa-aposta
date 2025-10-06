# 🚨 README de Erros - Casa Aposta

Este documento contém todos os erros encontrados no projeto Casa Aposta e suas respectivas soluções. Use este guia como referência para resolver problemas similares.

## 📋 Índice

1. [Erros de Compilação](#erros-de-compilação)
2. [Erros de Runtime](#erros-de-runtime)
3. [Erros de Configuração](#erros-de-configuração)
4. [Erros Futuros](#erros-futuros)
5. [Como Reportar Novos Erros](#como-reportar-novos-erros)

---

## 🔧 Erros de Compilação

### ❌ Erro 1: `undefined: GetStatusConnection`

**Arquivo:** `database/database.go:79`

**Mensagem de Erro:**
```
database\database.go:79:5: undefined: GetStatusConnection
```

**Problema:**
- Função `GetStatusConnection()` estava sendo chamada como função global
- Mas estava definida como método de `*ConnectionStatus`

**Código Problemático:**
```go
func ConnectSimulationDatabase(db Database, data map[string] interface{}) (string, error) {
    if GetStatusConnection() != Disconnected {  // ❌ Erro aqui
        return "", errors.New("database: connection is not disconnected")
    }
}
```

**✅ Solução:**
```go
// Adicionada variável global
var currentConnectionStatus = Disconnected

func ConnectSimulationDatabase(db Database, data map[string] interface{}) (string, error) {
    if currentConnectionStatus != Disconnected {  // ✅ Corrigido
        return "", errors.New("database: connection is not disconnected")
    }
}

// Funções auxiliares adicionadas
func GetCurrentConnectionStatus() ConnectionStatus {
    return currentConnectionStatus
}

func SetConnectionStatus(status ConnectionStatus) {
    currentConnectionStatus = status
}
```

---

### ❌ Erro 2: Ordem de Definição de Tipos

**Arquivo:** `database/database.go:60-76`

**Mensagem de Erro:**
```
undefined: ConnectionStatus (used before declaration)
```

**Problema:**
- Constantes usando `ConnectionStatus` definidas antes do tipo
- Go requer que tipos sejam definidos antes de serem usados

**Código Problemático:**
```go
const (
    Disconnected ConnectionStatus = iota  // ❌ ConnectionStatus não definido ainda
    Connected
    Connecting
    Error
)

type ConnectionStatus int  // ❌ Definido depois das constantes
```

**✅ Solução:**
```go
// Definindo os tipos primeiro
type ConnectionStatus int
type DatabaseOperation string

// Enum para Status de Conexão
const (
    Disconnected ConnectionStatus = iota  // ✅ Agora funciona
    Connected
    Connecting
    Error
)
```

---

### ❌ Erro 3: Assignment Mismatch

**Arquivo:** `runtime/runtime.go:10`

**Mensagem de Erro:**
```
runtime\runtime.go:10:10: assignment mismatch: 1 variable but database.GetTableName returns 2 values
runtime\runtime.go:10:32: cannot use model (variable of type interface{}) as string value in argument to database.GetTableName
```

**Problema:**
- `GetTableName` retorna 2 valores (string, error) mas só 1 estava sendo capturado
- Type assertion incorreta de `interface{}` para `string`

**Código Problemático:**
```go
func RuntimeVerificationTables() error {
    for _, model := range models.AllModels {
        err := database.GetTableName(model)  // ❌ Só captura 1 valor, função retorna 2
        // ❌ model é interface{}, não string
        if err != nil {
            return err
        }
    }
    return nil
}
```

**✅ Solução:**
```go
import "reflect"

func RuntimeVerificationTables() error {
    for _, model := range models.AllModels {
        // Obtendo o nome do tipo usando reflection
        modelType := reflect.TypeOf(model)
        modelName := modelType.Name()
        
        _, err := database.GetTableName(modelName)  // ✅ Captura ambos os valores
        if err != nil {
            return err
        }
    }
    return nil
}
```

---

## 🏃‍♂️ Erros de Runtime

### ❌ Erro 4: `panic: database: operation is nil`

**Arquivo:** `cmd/api.go:14`

**Mensagem de Erro:**
```
panic: database: operation is nil

goroutine 1 [running]:
main.main()
        D:/GitClone/casa-aposta/cmd/api.go:14 +0x231
exit status 2
```

**Problema:**
- `SetDatabase` chamado com mapa vazio
- `ConnectSimulationDatabase` espera chave "operation" no mapa

**Código Problemático:**
```go
func main() {
    err := database.SetDatabase(map[string]interface{}{
        // ❌ Mapa vazio - falta "operation"
    })
    
    if err != nil {
        panic(err)  // ❌ Panic aqui
    }
}
```

**✅ Solução:**
```go
func main() {
    err := database.SetDatabase(map[string]interface{}{
        "operation": database.OperationConnect,  // ✅ Operação adicionada
    })
    
    if err != nil {
        panic(err)
    }
}
```

---

### ❌ Erro 5: `no Go files in directory`

**Comando:** `go run .`

**Mensagem de Erro:**
```
no Go files in D:\GitClone\casa-aposta
```

**Problema:**
- Tentativa de executar `go run .` no diretório raiz
- Arquivo `main.go` está em `cmd/api.go`

**Comando Problemático:**
```bash
go run .  # ❌ Não há main.go no diretório raiz
```

**✅ Solução:**
```bash
go run cmd/api.go  # ✅ Executa o arquivo correto
```

---

## ⚙️ Erros de Configuração

### ❌ Erro 6: Estrutura de Diretórios Inconsistente

**Problema:**
- Arquivo `database.json` duplicado em locais diferentes
- Confusão entre configuração e dados

**Estrutura Problemática:**
```
database/
├── database.json          # ❌ Duplicado
├── database/
│   └── database.json      # ❌ Duplicado
└── games_data.json        # ✅ Correto
```

**✅ Solução:**
```
database/
├── database/
│   └── database.json      # ✅ Configuração do banco
└── tables/
    └── games_data.json    # ✅ Dados dos jogos
```

---

---

### ❌ Erro 7: Fluxo Lógico Incorreto - Verificação de Tabelas Antes da Conexão

**Arquivo:** `cmd/api.go:10-16`

**Problema:**
- `RuntimeVerificationTables()` era chamado **ANTES** da conexão com o banco
- `SetDatabase()` (conexão) era chamado **DEPOIS** da verificação
- Isso não faz sentido lógico - como verificar tabelas sem estar conectado?

**Código Problemático:**
```go
func main() {
    errorRuntime := runtime.RuntimeVerificationTables()  // ❌ Verifica tabelas SEM conexão
    
    if errorRuntime != nil {
        panic(errorRuntime)
    }
    
    err := database.SetDatabase(map[string]interface{}{  // ❌ Conecta DEPOIS
        "operation": database.OperationConnect,
    })
}
```

**✅ Solução:**
```go
func main() {
    // Primeiro conecta ao banco
    err := database.SetDatabase(map[string]interface{}{  // ✅ Conecta PRIMEIRO
        "operation": database.OperationConnect,
    })

    if err != nil {
        panic(err)
    }
    
    // Depois verifica as tabelas (agora que está conectado)
    errorRuntime := runtime.RuntimeVerificationTables()  // ✅ Verifica DEPOIS da conexão

    if errorRuntime != nil {
        panic(errorRuntime)
    }
}
```

**Impacto:**
- ✅ Fluxo lógico correto: Conectar → Verificar → Operar
- ✅ Evita erros de acesso a recursos não inicializados
- ✅ Melhora a legibilidade e manutenibilidade do código

---

### ❌ Erro 8: Função GetTableName com Assinatura Incorreta e Lógica Invertida

**Arquivo:** `database/database.go:143-158`

**Problemas:**
1. **Assinatura incorreta**: Retornava apenas `error` em vez de `(string, error)`
2. **Lógica invertida**: Retornava erro quando a tabela **existia** em vez de quando **não existia**
3. **Inconsistência**: Nome da função sugere retornar string, mas só retornava error

**Código Problemático:**
```go
func GetTableName(model string) (error){  // ❌ Deveria retornar (string, error)
    fmt.Println("database: get table name", model)
    if model == "" {
        return errors.New("database: model is empty")
    }

    _, err := os.Stat("database/tables/" + model + ".json")
    if err == nil {  // ❌ Lógica invertida - erro quando existe!
        return errors.New("database: table " + model + " already exist")
    }

    // Se chegou até aqui, a tabela existe
    fmt.Println("database: table", model, "exists")

    return nil
}
```

**✅ Solução:**
```go
func GetTableName(model string) (string, error){  // ✅ Assinatura correta
    fmt.Println("database: get table name", model)
    if model == "" {
        return "", errors.New("database: model is empty")
    }

    tablePath := "database/tables/" + model + ".json"
    _, err := os.Stat(tablePath)
    if err != nil {  // ✅ Lógica correta - erro quando NÃO existe
        return "", errors.New("database: table " + model + " does not exist")
    }

    // Se chegou até aqui, a tabela existe
    fmt.Println("database: table", model, "exists")

    return tablePath, nil  // ✅ Retorna o caminho da tabela
}
```

**Correção no runtime.go:**
```go
// Antes (incorreto)
err := database.GetTableName(modelName)  // ❌ Só captura 1 valor

// Depois (correto)  
_, err := database.GetTableName(modelName)  // ✅ Captura ambos os valores
```

**Impacto:**
- ✅ Função agora retorna informação útil (caminho da tabela)
- ✅ Lógica correta: erro quando tabela não existe
- ✅ Assinatura consistente com o nome da função
- ✅ Compatível com padrões Go de retorno múltiplo

---

## 🔮 Erros Futuros

Esta seção será atualizada conforme novos erros forem reportados.

### Como Adicionar Novos Erros

Quando encontrar um novo erro, adicione-o seguindo este formato:

```markdown
### ❌ Erro X: [Título Descritivo]

**Arquivo:** `caminho/do/arquivo.go:linha`

**Mensagem de Erro:**
```
[Mensagem exata do erro]
```

**Problema:**
- Descrição clara do que causou o erro

**Código Problemático:**
```go
// Código que causa o erro
```

**✅ Solução:**
```go
// Código corrigido
```
```

---

## 📝 Como Reportar Novos Erros

Quando encontrar um erro:

1. **Copie a mensagem exata** do erro
2. **Identifique o arquivo e linha** onde ocorre
3. **Descreva o contexto** - o que você estava tentando fazer
4. **Inclua o código problemático** se possível
5. **Adicione neste documento** seguindo o formato acima

### Comandos Úteis para Diagnóstico

```bash
# Verificar erros de compilação
go build ./...

# Executar com informações detalhadas
go run -v cmd/api.go

# Verificar sintaxe
go fmt ./...

# Executar testes
go test ./...
```

---

## 🛠️ Ferramentas de Debug

### Verificação de Tipos
```go
import "reflect"

func debugType(v interface{}) {
    fmt.Printf("Type: %T, Value: %v\n", v, v)
    fmt.Printf("Reflect Type: %s\n", reflect.TypeOf(v).Name())
}
```

### Log de Conexão
```go
func logConnectionStatus() {
    fmt.Printf("Current Connection Status: %s\n", currentConnectionStatus.String())
}
```

---

## 🔄 Erro 9: Lógica Invertida na Função CreateTablesIfNotExist

**📁 Arquivo:** `database/database.go`  
**🔍 Tipo:** Erro de Lógica Invertida  
**⚠️ Descrição:** A função `CreateTablesIfNotExist` retornava `true` quando o arquivo **já existia**, causando confusão na lógica de controle.

### ❌ Código Problemático:
```go
func CreateTablesIfNotExist(data map[string] interface{}, filename string) bool {
    _, err := os.Stat(filename)
    if err == nil {
        fmt.Println("database: tables already exist")
        return true  // ❌ ERRO: Retorna TRUE quando arquivo JÁ EXISTE
    }
    
    err = os.WriteFile(filename, []byte("[]"), 0644)
    if err != nil {
        return false
    }
    return true
}
```

### ✅ Código Corrigido:
```go
func CreateTablesIfNotExist(data map[string] interface{}, filename string) bool {
    _, err := os.Stat(filename)
    if err == nil {
        // Arquivo já existe, não precisa criar
        return false
    }

    // Arquivo não existe, vamos criar
    err = os.WriteFile(filename, []byte("[]"), 0644)
    if err != nil {
        fmt.Printf("database: error creating file %s: %v\n", filename, err)
        return false
    }
    return true // Sucesso na criação
}
```

### 🎯 Impacto Positivo:
- ✅ Lógica clara: `true` = criou arquivo, `false` = não criou
- ✅ Melhor controle de fluxo
- ✅ Mensagens de erro mais informativas

---

## 🔄 Erro 10: Lógica Invertida na Função GetTableName

**📁 Arquivo:** `database/database.go`  
**🔍 Tipo:** Erro de Lógica Invertida  
**⚠️ Descrição:** A função `GetTableName` retornava erro quando `CreateTablesIfNotExist` tinha **sucesso**, causando panic mesmo quando o arquivo era criado corretamente.

### ❌ Código Problemático:
```go
if CreateTablesIfNotExist(map[string]interface{}{
    "operation": OperationConnect,
}, tablePath) {
    return "", errors.New("database: table " + model + " not created")  // ❌ ERRO!
}
```

### ✅ Código Corrigido:
```go
fmt.Printf("database: table %s does not exist, creating...\n", model)
if !CreateTablesIfNotExist(map[string]interface{}{
    "operation": OperationConnect,
}, tablePath) {
    return "", errors.New("database: failed to create table " + model)
}
fmt.Printf("database: table %s created successfully\n", model)
```

### 🎯 Impacto Positivo:
- ✅ Não há mais panic quando arquivo é criado com sucesso
- ✅ Mensagens informativas sobre criação de tabelas
- ✅ Fluxo lógico correto: erro apenas quando falha

---

## 📊 Estatísticas dos Erros

- **Total de Erros Documentados:** 10
- **Erros de Compilação:** 3
- **Erros de Runtime:** 2
- **Erros de Configuração:** 1
- **Erros de Fluxo Lógico:** 1
- **Erros de Design de Função:** 1
- **Erros de Lógica Invertida:** 2
- **Status:** ✅ Todos Resolvidos

---

## 🎯 Prevenção de Erros

### Boas Práticas

1. **Sempre compile antes de executar:**
   ```bash
   go build ./... && go run cmd/api.go
   ```

2. **Use ferramentas de análise:**
   ```bash
   go vet ./...
   go fmt ./...
   ```

3. **Teste regularmente:**
   ```bash
   go test ./...
   ```

4. **Verifique tipos com reflection quando necessário**

5. **Mantenha documentação atualizada**

---

*Última atualização: $(date)*
*Versão do Go: 1.25.1*
*Projeto: Casa Aposta*