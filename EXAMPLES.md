# 🚀 Exemplos Práticos - Casa Aposta

Este arquivo contém exemplos práticos de como usar e estender o sistema Casa Aposta.

## 📋 Índice

- [Executando o Projeto](#executando-o-projeto)
- [Exemplos de Operações](#exemplos-de-operações)
- [Adicionando Novas Funcionalidades](#adicionando-novas-funcionalidades)
- [Exemplos de Código](#exemplos-de-código)
- [Cenários de Uso](#cenários-de-uso)

## 🎮 Executando o Projeto

### Execução Básica

```bash
# Navegue até o diretório do projeto
cd casa-aposta

# Execute a aplicação
go run cmd/api.go
```

**Saída esperada:**
```
Database connected successfully!
Data inserted successfully!
```

### Verificando os Dados

```bash
# Visualizar dados dos jogos
cat database/tables/games_data.json

# Visualizar configuração do banco
cat database/database/database.json
```

## 🔧 Exemplos de Operações

### 1. Conectando ao Banco de Dados

```go
package main

import (
    "fmt"
    "casa-aposta/database"
)

func main() {
    // Configurar banco de dados
    err := database.SetDatabase()
    if err != nil {
        fmt.Printf("Erro ao configurar banco: %v\n", err)
        return
    }
    
    fmt.Println("Banco configurado com sucesso!")
}
```

### 2. Inserindo Dados de Jogos

```go
package main

import (
    "fmt"
    "casa-aposta/database"
)

func main() {
    // Inserir dados
    err := database.Insert()
    if err != nil {
        fmt.Printf("Erro ao inserir dados: %v\n", err)
        return
    }
    
    fmt.Println("Dados inseridos com sucesso!")
}
```

### 3. Verificação de Runtime

```go
package main

import (
    "fmt"
    "casa-aposta/runtime"
)

func main() {
    // Verificar modelos
    err := runtime.RunTimeVerificationTables()
    if err != nil {
        fmt.Printf("Erro na verificação: %v\n", err)
        return
    }
    
    fmt.Println("Verificação concluída!")
}
```

## 🆕 Adicionando Novas Funcionalidades

### Exemplo 1: Criando Tabela de Usuários

#### 1. Criar arquivo JSON para usuários

```bash
# Criar arquivo de usuários
echo "[]" > database/tables/users_data.json
```

#### 2. Implementar função de inserção

```go
// Adicionar em database/database.go

func InsertUser(name, password string) error {
    dataFile := "database/tables/users_data.json"
    
    // Ler dados existentes
    data, err := os.ReadFile(dataFile)
    if err != nil {
        return err
    }
    
    var users []map[string]interface{}
    json.Unmarshal(data, &users)
    
    // Criar novo usuário
    newUser := map[string]interface{}{
        "id":       len(users) + 1,
        "name":     name,
        "password": password,
    }
    
    users = append(users, newUser)
    
    // Salvar dados atualizados
    updatedData, _ := json.MarshalIndent(users, "", "  ")
    return os.WriteFile(dataFile, updatedData, 0644)
}
```

#### 3. Usar a nova função

```go
// Em cmd/api.go
func main() {
    // Configurar banco
    database.SetDatabase()
    
    // Inserir usuário
    err := database.InsertUser("João", "senha123")
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }
    
    fmt.Println("Usuário criado com sucesso!")
}
```

### Exemplo 2: Sistema de Apostas

#### 1. Criar tabela de apostas

```bash
echo "[]" > database/tables/bets_data.json
```

#### 2. Implementar estrutura de aposta

```go
// Adicionar em models/
type Bet struct {
    ID       int     `json:"id"`
    UserID   int     `json:"user_id"`
    GameName string  `json:"game_name"`
    Amount   float64 `json:"amount"`
    Status   string  `json:"status"` // "pending", "won", "lost"
}
```

#### 3. Implementar função de aposta

```go
func PlaceBet(userID int, gameName string, amount float64) error {
    dataFile := "database/tables/bets_data.json"
    
    data, err := os.ReadFile(dataFile)
    if err != nil {
        return err
    }
    
    var bets []Bet
    json.Unmarshal(data, &bets)
    
    newBet := Bet{
        ID:       len(bets) + 1,
        UserID:   userID,
        GameName: gameName,
        Amount:   amount,
        Status:   "pending",
    }
    
    bets = append(bets, newBet)
    
    updatedData, _ := json.MarshalIndent(bets, "", "  ")
    return os.WriteFile(dataFile, updatedData, 0644)
}
```

## 💻 Exemplos de Código Completos

### Exemplo Completo: Sistema de Usuários

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
    "casa-aposta/models"
)

func CreateUser(name, password string) error {
    dataFile := "database/tables/users_data.json"
    
    // Ler usuários existentes
    data, err := os.ReadFile(dataFile)
    if err != nil {
        // Se arquivo não existe, criar array vazio
        data = []byte("[]")
    }
    
    var users []models.Usuario
    json.Unmarshal(data, &users)
    
    // Criar novo usuário
    newUser := models.Usuario{
        ID:       len(users) + 1,
        Name:     name,
        Password: password,
    }
    
    users = append(users, newUser)
    
    // Salvar
    updatedData, _ := json.MarshalIndent(users, "", "  ")
    return os.WriteFile(dataFile, updatedData, 0644)
}

func GetAllUsers() ([]models.Usuario, error) {
    dataFile := "database/tables/users_data.json"
    
    data, err := os.ReadFile(dataFile)
    if err != nil {
        return nil, err
    }
    
    var users []models.Usuario
    err = json.Unmarshal(data, &users)
    return users, err
}

func main() {
    // Criar usuários
    CreateUser("João", "senha123")
    CreateUser("Maria", "senha456")
    
    // Listar usuários
    users, err := GetAllUsers()
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }
    
    fmt.Println("Usuários cadastrados:")
    for _, user := range users {
        fmt.Printf("ID: %d, Nome: %s\n", user.ID, user.Name)
    }
}
```

### Exemplo: API REST Simples

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "casa-aposta/database"
    "casa-aposta/models"
)

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
    users, err := GetAllUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }
    
    var user models.Usuario
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    err = CreateUser(user.Name, user.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Usuário criado com sucesso!")
}

func main() {
    // Configurar banco
    database.SetDatabase()
    
    // Rotas
    http.HandleFunc("/users", getUsersHandler)
    http.HandleFunc("/users/create", createUserHandler)
    
    fmt.Println("Servidor rodando em http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

## 🎯 Cenários de Uso

### Cenário 1: Cadastro de Novo Jogo

```bash
# 1. Execute a aplicação
go run cmd/api.go

# 2. Verifique os dados inseridos
cat database/tables/games_data.json

# 3. Resultado esperado:
# [
#   {
#     "name": "Tigrinho",
#     "score": 100,
#     "tipo": "esports"
#   }
# ]
```

### Cenário 2: Backup dos Dados

```bash
# Criar backup
mkdir backup
cp -r database/tables/ backup/
cp -r database/database/ backup/

# Verificar backup
ls backup/
```

### Cenário 3: Restaurar Dados

```bash
# Restaurar do backup
cp -r backup/tables/ database/
cp -r backup/database/ database/

# Verificar restauração
go run cmd/api.go
```

### Cenário 4: Limpar Dados

```bash
# Limpar tabela de jogos
echo "[]" > database/tables/games_data.json

# Verificar limpeza
cat database/tables/games_data.json
```

## 🔧 Comandos Úteis

### Desenvolvimento

```bash
# Executar aplicação
go run cmd/api.go

# Compilar aplicação
go build -o casa-aposta cmd/api.go

# Executar compilado
./casa-aposta
```

### Verificação de Dados

```bash
# Verificar sintaxe JSON
python -m json.tool database/tables/games_data.json

# Contar registros
cat database/tables/games_data.json | jq length

# Filtrar dados (se jq estiver instalado)
cat database/tables/games_data.json | jq '.[] | select(.tipo=="esports")'
```

### Manutenção

```bash
# Backup automático
tar -czf backup-$(date +%Y%m%d).tar.gz database/

# Verificar tamanho dos arquivos
du -h database/tables/*

# Limpar logs (se houver)
> logs/application.log
```

## 📊 Monitoramento

### Script de Monitoramento Simples

```bash
#!/bin/bash
# monitor.sh

echo "=== Status do Casa Aposta ==="
echo "Data: $(date)"
echo

echo "Arquivos de dados:"
ls -la database/tables/

echo
echo "Tamanho dos arquivos:"
du -h database/tables/*

echo
echo "Último jogo adicionado:"
tail -n 5 database/tables/games_data.json
```

### Uso do Script

```bash
chmod +x monitor.sh
./monitor.sh
```

## 🚨 Troubleshooting

### Problema: Arquivo JSON corrompido

```bash
# Verificar sintaxe
python -m json.tool database/tables/games_data.json

# Se corrompido, restaurar backup
cp backup/tables/games_data.json database/tables/
```

### Problema: Permissões de arquivo

```bash
# Corrigir permissões
chmod 644 database/tables/*.json
chmod 644 database/database/*.json
```

### Problema: Aplicação não inicia

```bash
# Verificar Go
go version

# Verificar módulos
go mod tidy

# Executar com verbose
go run -v cmd/api.go
```

---

Estes exemplos devem cobrir a maioria dos casos de uso do sistema Casa Aposta. Para mais informações, consulte o README.md e DATABASE.md.