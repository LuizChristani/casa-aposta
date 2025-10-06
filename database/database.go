package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

type Database struct {
	URL     string
	DBName  string
	DBSenha string
}

type DatabaseObj struct {
	Database struct {
		URL     string  `json:"URL"`
		DBName  string  `json:"DBName"`
		DBSenha string  `json:"DBSenha"`
	} `json:"name"`
}



// Definindo os tipos primeiro
type ConnectionStatus int
type DatabaseOperation string

// Enum para Status de Conexão
const (
	Disconnected ConnectionStatus = iota
	Connected
	Connecting
	Error
)

// Enum para Operações do Banco
const (
	OperationConnect    DatabaseOperation = "Connect"
	OperationDisconnect DatabaseOperation = "Disconnect"
	OperationQuery      DatabaseOperation = "Query"
	OperationInsert     DatabaseOperation = "Insert"
	OperationUpdate     DatabaseOperation = "Update"
	OperationDelete     DatabaseOperation = "Delete"
)

// Variável global para controlar o status da conexão
var currentConnectionStatus = Disconnected

func ConnectSimulationDatabase(db Database, data map[string] interface{}) (string, error) {

	if currentConnectionStatus != Disconnected {
		return "", errors.New("database: connection is not disconnected")
	}

	if data == nil {
		return "", errors.New("database: data is nil")
	}

	if data["operation"] == nil {
		return "", errors.New("database: operation is nil")
	}

	if db.URL == "" || db.DBName == "" || db.DBSenha == "" {
		return "", errors.New("database: URL, DBName, DBSenha are required")
	}
	var jsonFile *os.File
	var err error
	jsonFile, err = os.Open("database/database/database.json")
	if err == nil {
		currentConnectionStatus = Connected
	} else {
		currentConnectionStatus = Error
	}
	if currentConnectionStatus == Error {
		return "", errors.New("database: connection is error")
	}

	if currentConnectionStatus == Connecting {
		time.Sleep(300 * time.Second)
		jsonFile, err = os.Open("database/databases/database.json")
		if err != nil {
			return "", errors.New("database: connection is connecting")
		}
		defer jsonFile.Close()
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return "", err
	}

	objDatabase := DatabaseObj{}
	err = json.Unmarshal(byteValue, &objDatabase)
	if err != nil {
		return "", err
	}

	if objDatabase.Database.URL != db.URL || objDatabase.Database.DBName != db.DBName || objDatabase.Database.DBSenha != db.DBSenha {
		currentConnectionStatus = Disconnected
		return "", errors.New("database: URL, DBName, DBSenha are not valid")
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", db.DBName, db.DBSenha, db.URL, db.DBName)
	fmt.Println("database: connected")
	return connectionString, nil
}

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

func Insert(data map[string]interface{}) error {
	var db []map[string]interface{}
	dataFile := "database/tables/games_data.json"
	if _, err := os.Stat(dataFile); err == nil {
		bytes, err := os.ReadFile(dataFile)
		if err != nil {
			return err
		}
		json.Unmarshal(bytes, &db)
	}

	db = append(db, data)

	bytes, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(dataFile, bytes, 0644)
}

func GetTableName(model string) (string, error){
	fmt.Println("database: get table name", model)
	if model == "" {
		return "", errors.New("database: model is empty")
	}

	tablePath := "database/tables/" + model + ".json"
	_, err := os.Stat(tablePath)
	if err != nil {
		// Se não conseguir acessar o arquivo, a tabela não existe
		fmt.Printf("database: table %s does not exist, creating...\n", model)
		if !CreateTablesIfNotExist(map[string]interface{}{
			"operation": OperationConnect,
		}, tablePath) {
			return "", errors.New("database: failed to create table " + model)
		}
		fmt.Printf("database: table %s created successfully\n", model)
	}

	// Se chegou até aqui, a tabela existe
	fmt.Println("database: table", model, "exists")

	return tablePath, nil
}

// Função para obter o status atual da conexão
func GetCurrentConnectionStatus() ConnectionStatus {
	return currentConnectionStatus
}

// Função para definir o status da conexão
func SetConnectionStatus(status ConnectionStatus) {
	currentConnectionStatus = status
}

// Método para obter status de uma instância específica
func (connection *ConnectionStatus) GetStatusConnection() ConnectionStatus {
	return *connection
}

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
