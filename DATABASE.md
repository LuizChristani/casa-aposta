# 💾 Documentação do Banco de Dados JSON

## Visão Geral

O projeto Casa Aposta utiliza um sistema de banco de dados baseado em arquivos JSON, onde cada arquivo representa uma "tabela" de dados. Esta abordagem oferece simplicidade, portabilidade e facilidade de manutenção.

## 🗂️ Estrutura de Diretórios

```
database/
├── configuration_db.go      # Funções de configuração do banco
├── database.go             # Operações CRUD do banco
├── database/               # Configurações de conexão
│   └── database.json      # Arquivo de configuração principal
└── tables/                # Tabelas de dados
    └── games_data.json    # Dados dos jogos
```

## 📋 Tabelas (Arquivos JSON)

### 1. Configuração do Banco (`database/database/database.json`)

**Propósito**: Armazena as configurações de conexão e metadados do banco de dados.

**Estrutura**:
```json
{
  "name": {
    "URL": "localhost:5432",
    "DBName": "casa_aposta_db", 
    "DBSenha": "senha123"
  }
}
```

**Campos**:
- `URL`: Endereço do servidor de banco de dados
- `DBName`: Nome do banco de dados
- `DBSenha`: Senha de acesso ao banco

**Uso**: Este arquivo é lido pela função `ConnectSimulationDatabase()` para estabelecer a "conexão" simulada.

### 2. Dados dos Jogos (`database/tables/games_data.json`)

**Propósito**: Armazena informações sobre os jogos disponíveis para apostas.

**Estrutura**:
```json
[
  {
    "name": "Tigrinho",
    "score": 100,
    "tipo": "esports"
  },
  {
    "name": "Futebol Virtual",
    "score": 85,
    "tipo": "esports"
  }
]
```

**Campos**:
- `name`: Nome do jogo
- `score`: Pontuação ou rating do jogo
- `tipo`: Categoria do jogo (esports, casino, etc.)

**Uso**: Este arquivo é manipulado pela função `Insert()` para adicionar novos jogos.

## 🔧 Operações Disponíveis

### Conexão com Banco de Dados

```go
func ConnectSimulationDatabase(db Database) (Database, error)
```

**Descrição**: Simula a conexão com o banco de dados lendo o arquivo de configuração.

**Parâmetros**:
- `db`: Estrutura Database vazia

**Retorno**:
- `Database`: Estrutura preenchida com dados de configuração
- `error`: Erro caso a operação falhe

**Arquivo utilizado**: `database/database/database.json`

### Inserção de Dados

```go
func Insert() error
```

**Descrição**: Insere novos dados de jogos na tabela games_data.

**Comportamento**:
1. Lê o arquivo `database/tables/games_data.json`
2. Adiciona novos dados ao array existente
3. Salva o arquivo atualizado

**Arquivo utilizado**: `database/tables/games_data.json`

### Configuração do Banco

```go
func SetDatabase() error
```

**Descrição**: Inicializa e configura o banco de dados.

**Comportamento**:
1. Cria uma instância Database vazia
2. Chama `ConnectSimulationDatabase()` para carregar configurações
3. Retorna erro se a operação falhar

## 📊 Estruturas de Dados

### Database

```go
type Database struct {
    Name DatabaseConfig `json:"name"`
}

type DatabaseConfig struct {
    URL     string `json:"URL"`
    DBName  string `json:"DBName"`
    DBSenha string `json:"DBSenha"`
}
```

### Game (Implícito)

```go
type Game struct {
    Name  string `json:"name"`
    Score int    `json:"score"`
    Tipo  string `json:"tipo"`
}
```

## 🚀 Como Adicionar Novas Tabelas

### 1. Criar o Arquivo JSON

```bash
# Crie um novo arquivo na pasta tables/
touch database/tables/nova_tabela.json

# Inicialize com um array vazio
echo "[]" > database/tables/nova_tabela.json
```

### 2. Definir a Estrutura

```json
[
  {
    "campo1": "valor1",
    "campo2": "valor2",
    "campo3": 123
  }
]
```

### 3. Implementar Operações

```go
func InsertNovaTabelaData() error {
    dataFile := "database/tables/nova_tabela.json"
    
    // Ler dados existentes
    data, err := os.ReadFile(dataFile)
    if err != nil {
        return err
    }
    
    var records []NovoTipo
    json.Unmarshal(data, &records)
    
    // Adicionar novos dados
    newRecord := NovoTipo{
        Campo1: "novo valor",
        Campo2: "outro valor",
        Campo3: 456,
    }
    records = append(records, newRecord)
    
    // Salvar dados atualizados
    updatedData, _ := json.MarshalIndent(records, "", "  ")
    return os.WriteFile(dataFile, updatedData, 0644)
}
```

## 🔍 Exemplos de Uso

### Lendo Dados de uma Tabela

```go
func ReadGamesData() ([]Game, error) {
    data, err := os.ReadFile("database/tables/games_data.json")
    if err != nil {
        return nil, err
    }
    
    var games []Game
    err = json.Unmarshal(data, &games)
    return games, err
}
```

### Atualizando um Registro

```go
func UpdateGame(gameName string, newScore int) error {
    // Ler dados existentes
    games, err := ReadGamesData()
    if err != nil {
        return err
    }
    
    // Encontrar e atualizar o jogo
    for i, game := range games {
        if game.Name == gameName {
            games[i].Score = newScore
            break
        }
    }
    
    // Salvar dados atualizados
    data, _ := json.MarshalIndent(games, "", "  ")
    return os.WriteFile("database/tables/games_data.json", data, 0644)
}
```

### Deletando um Registro

```go
func DeleteGame(gameName string) error {
    games, err := ReadGamesData()
    if err != nil {
        return err
    }
    
    // Filtrar jogos (remover o especificado)
    var filteredGames []Game
    for _, game := range games {
        if game.Name != gameName {
            filteredGames = append(filteredGames, game)
        }
    }
    
    // Salvar dados filtrados
    data, _ := json.MarshalIndent(filteredGames, "", "  ")
    return os.WriteFile("database/tables/games_data.json", data, 0644)
}
```

## ⚠️ Considerações Importantes

### Vantagens do Sistema JSON

- **Simplicidade**: Fácil de entender e implementar
- **Portabilidade**: Funciona em qualquer ambiente
- **Versionamento**: Pode ser versionado com Git
- **Debugging**: Dados facilmente legíveis
- **Backup**: Simples cópia de arquivos

### Limitações

- **Performance**: Não adequado para grandes volumes de dados
- **Concorrência**: Sem controle de transações simultâneas
- **Integridade**: Sem validações automáticas de dados
- **Relacionamentos**: Sem suporte nativo a foreign keys
- **Indexação**: Sem otimizações de busca

### Boas Práticas

1. **Backup Regular**: Faça backup dos arquivos JSON regularmente
2. **Validação**: Sempre valide dados antes de inserir
3. **Tratamento de Erros**: Implemente tratamento robusto de erros
4. **Estrutura Consistente**: Mantenha estruturas de dados consistentes
5. **Documentação**: Documente mudanças na estrutura dos dados

## 🔧 Troubleshooting

### Arquivo JSON Corrompido

```bash
# Verificar sintaxe JSON
python -m json.tool database/tables/games_data.json

# Ou usando jq (se instalado)
jq . database/tables/games_data.json
```

### Restaurar Backup

```bash
# Copiar backup
cp database/tables/games_data.json.backup database/tables/games_data.json
```

### Resetar Tabela

```bash
# Limpar dados (manter estrutura)
echo "[]" > database/tables/games_data.json
```

## 📈 Migração Futura

Quando o projeto crescer, considere migrar para:

- **SQLite**: Para melhor performance local
- **PostgreSQL**: Para aplicações em produção
- **MongoDB**: Para dados não-relacionais complexos

O sistema atual facilita essa migração, pois os dados já estão estruturados e as operações são bem definidas.

---

Esta documentação deve ser atualizada sempre que novas tabelas ou operações forem adicionadas ao sistema.