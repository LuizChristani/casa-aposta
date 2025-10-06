# 🎰 Casa Aposta

Uma aplicação de apostas esportivas desenvolvida em Go, utilizando um sistema de banco de dados baseado em arquivos JSON para armazenamento de dados.

## 📋 Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Arquitetura](#arquitetura)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Banco de Dados JSON](#banco-de-dados-json)
- [Instalação](#instalação)
- [Como Usar](#como-usar)
- [Modelos de Dados](#modelos-de-dados)
- [API](#api)
- [Contribuição](#contribuição)
- [Licença](#-licença)

## 🎯 Sobre o Projeto

Casa Aposta é uma aplicação backend para gerenciamento de apostas esportivas. O projeto utiliza uma arquitetura simples e eficiente, com armazenamento de dados em arquivos JSON, proporcionando uma solução leve e de fácil manutenção.

### Características Principais

- ✅ Sistema de banco de dados baseado em JSON
- ✅ Arquitetura modular e escalável
- ✅ Gerenciamento de usuários
- ✅ Sistema de jogos e apostas
- ✅ API RESTful (em desenvolvimento)
- ✅ Configuração flexível de banco de dados

## 🏗️ Arquitetura

O projeto segue uma arquitetura em camadas bem definida:

```
Casa Aposta
├── Camada de Apresentação (cmd/)
├── Camada de Modelos (models/)
├── Camada de Serviços (service/) - Em desenvolvimento
├── Camada de Repositório (repository/) - Em desenvolvimento
├── Camada de Dados (database/)
└── Camada de Runtime (runtime/)
```

### Componentes Principais

- **CMD**: Ponto de entrada da aplicação
- **Models**: Definição das estruturas de dados
- **Database**: Gerenciamento do banco de dados JSON
- **Runtime**: Verificações e validações em tempo de execução
- **Service**: Lógica de negócio (planejado)
- **Repository**: Camada de acesso aos dados (planejado)

## 📁 Estrutura do Projeto

```
casa-aposta/
├── cmd/
│   └── api.go                 # Ponto de entrada da aplicação
├── database/
│   ├── configuration_db.go   # Configuração do banco de dados
│   ├── database.go          # Operações do banco de dados
│   ├── database/
│   │   └── database.json    # Configurações de conexão
│   └── tables/
│       └── games_data.json  # Dados dos jogos
├── models/
│   ├── all_models.go        # Registro de todos os modelos
│   └── usuario.go           # Modelo de usuário
├── repository/              # Camada de repositório (vazio)
├── runtime/
│   └── runtime.go           # Verificações de runtime
├── service/                 # Camada de serviços (vazio)
├── go.mod                   # Dependências do Go
└── README.md               # Este arquivo
```

## 💾 Banco de Dados JSON

O projeto utiliza um sistema de banco de dados baseado em arquivos JSON, onde cada arquivo representa uma "tabela":

### Estrutura do Banco de Dados

#### 1. Configuração (`database/database/database.json`)
```json
{
  "name": {
    "URL": "localhost:5432",
    "DBName": "casa_aposta_db",
    "DBSenha": "senha123"
  }
}
```

#### 2. Dados dos Jogos (`database/tables/games_data.json`)
```json
[
  {
    "name": "Tigrinho",
    "score": 100,
    "tipo": "esports"
  }
]
```

### Vantagens do Sistema JSON

- **Simplicidade**: Fácil de entender e modificar
- **Portabilidade**: Funciona em qualquer ambiente
- **Versionamento**: Pode ser versionado com Git
- **Debugging**: Fácil visualização dos dados
- **Performance**: Adequado para aplicações pequenas e médias

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

2. **Inicialize o módulo Go** (se necessário)
   ```bash
   go mod tidy
   ```

3. **Verifique a estrutura do banco de dados**
   ```bash
   # Certifique-se de que os arquivos JSON existem:
   # - database/database/database.json
   # - database/tables/games_data.json
   ```

## 🎮 Como Usar

### Executando a Aplicação

```bash
# Execute a aplicação principal
go run cmd/api.go
```

### Saída Esperada

```
Database connected successfully!
Data inserted successfully!
```

### Operações Disponíveis

1. **Conexão com Banco de Dados**: Estabelece conexão usando as configurações JSON
2. **Inserção de Dados**: Adiciona novos jogos à tabela de games
3. **Verificação de Runtime**: Valida a integridade dos modelos

## 📊 Modelos de Dados

### Usuario

```go
type Usuario struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Password string `json:"password"`
}
```

**Campos:**
- `ID`: Identificador único do usuário
- `Name`: Nome do usuário
- `Password`: Senha do usuário

### Game (Implícito)

```json
{
  "name": "string",    // Nome do jogo
  "score": "number",   // Pontuação
  "tipo": "string"     // Tipo do jogo (ex: esports)
}
```

## 🔧 API

### Endpoints Planejados

```
GET    /api/users          # Listar usuários
POST   /api/users          # Criar usuário
GET    /api/users/:id      # Obter usuário específico
PUT    /api/users/:id      # Atualizar usuário
DELETE /api/users/:id      # Deletar usuário

GET    /api/games          # Listar jogos
POST   /api/games          # Criar jogo
GET    /api/games/:id      # Obter jogo específico
PUT    /api/games/:id      # Atualizar jogo
DELETE /api/games/:id      # Deletar jogo
```

*Nota: A API está em desenvolvimento. Atualmente, apenas operações básicas de banco de dados estão implementadas.*

## 🛠️ Desenvolvimento

### Adicionando Novos Modelos

1. **Crie o modelo** em `models/`
   ```go
   type NovoModelo struct {
       ID   int    `json:"id"`
       Nome string `json:"nome"`
   }
   ```

2. **Registre o modelo** em `models/all_models.go`
   ```go
   var AllModels = []interface{}{
       Usuario{},
       NovoModelo{}, // Adicione aqui
   }
   ```

3. **Crie a tabela JSON** em `database/tables/`
   ```json
   []
   ```

### Adicionando Novas Operações de Banco

1. **Implemente a função** em `database/database.go`
2. **Teste a operação** em `cmd/api.go`
3. **Documente** a nova funcionalidade

## 🧪 Testes

```bash
# Execute a aplicação para testar
go run cmd/api.go

# Verifique os logs para confirmar:
# - Conexão com banco de dados
# - Inserção de dados
# - Ausência de erros
```

## 📝 Logs e Debugging

O sistema fornece logs detalhados:

- ✅ **Sucesso na conexão**: "Database connected successfully!"
- ✅ **Inserção bem-sucedida**: "Data inserted successfully!"
- ❌ **Erros**: Mensagens detalhadas de erro

## 🤝 Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

### Diretrizes de Contribuição

- Mantenha o código limpo e bem documentado
- Siga as convenções de nomenclatura do Go
- Teste suas mudanças antes de submeter
- Atualize a documentação quando necessário

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

## 👥 Autores

- **Desenvolvedor Principal** - *Trabalho inicial* - [Seu Nome]

## 🙏 Agradecimentos

- Comunidade Go pela excelente documentação
- Contribuidores do projeto
- Inspiração em sistemas de apostas modernos

---

**Casa Aposta** - Transformando apostas em experiências! 🎰✨