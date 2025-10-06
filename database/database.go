package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Definições de tipos e constantes
type ConnectionStatus int
type DatabaseOperation string

const (
	Disconnected ConnectionStatus = iota
	Connecting
	Connected
	Error
)

const (
	OperationCreate DatabaseOperation = "create"
	OperationRead   DatabaseOperation = "read"
	OperationUpdate DatabaseOperation = "update"
	OperationDelete DatabaseOperation = "delete"
)

var currentConnectionStatus = Disconnected



func CreateTablesIfNotExist(data map[string] interface{}, filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return false
	}

	err = os.WriteFile(filename, []byte("[]"), 0644)
	if err != nil {
		fmt.Printf("database: error creating file %s: %v\n", filename, err)
		return false
	}
	return true
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
			"operation": OperationCreate,
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

func ConnectSimulationDatabase() error {
	if currentConnectionStatus != Disconnected {
		return errors.New("database: connection is not disconnected")
	}
	currentConnectionStatus = Connecting
	return nil
}

func DisconnectSimulationDatabase() error {
	if currentConnectionStatus != Connected {
		return errors.New("database: connection is not connected")
	}
	currentConnectionStatus = Disconnected
	return nil
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
