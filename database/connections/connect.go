package connections

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type Database struct {
	URL     string
	DBName  string
	DBSenha string
}

// Estrutura compatível com o arquivo database/database/database.json
type DatabaseObj struct {
	Database struct {
		URL     string `json:"URL"`
		DBName  string `json:"DBName"`
		DBSenha string `json:"DBSenha"`
	} `json:"name"`
}

type ConnectionStatus int
type DatabaseOperation string

const (
	Disconnected ConnectionStatus = iota
	Connected
	Connecting
	Error
)

const (
	OperationConnect    DatabaseOperation = "Connect"
	OperationDisconnect DatabaseOperation = "Disconnect"
	OperationQuery      DatabaseOperation = "Query"
	OperationInsert     DatabaseOperation = "Insert"
	OperationUpdate     DatabaseOperation = "Update"
	OperationDelete     DatabaseOperation = "Delete"
)

var currentConnectionStatus = Disconnected

func ConnectSimulationDatabaseInitial(db Database, data map[string]interface{}) (string, error) {
	if data == nil || data["operation"] == nil {
		return "", errors.New("database: operation is required")
	}

	if db.URL == "" || db.DBName == "" || db.DBSenha == "" {
		return "", errors.New("database: URL, DBName, DBSenha are required")
	}

	jsonFile, err := os.Open("database/database/database.json")
	if err != nil {
		currentConnectionStatus = Error
		return "", err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		currentConnectionStatus = Error
		return "", err
	}

	var objDatabase DatabaseObj
	if err := json.Unmarshal(byteValue, &objDatabase); err != nil {
		currentConnectionStatus = Error
		return "", err
	}

	if objDatabase.Database.URL != db.URL || objDatabase.Database.DBName != db.DBName || objDatabase.Database.DBSenha != db.DBSenha {
		currentConnectionStatus = Disconnected
		return "", errors.New("database: URL, DBName, DBSenha are not valid")
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", db.DBName, db.DBSenha, db.URL, db.DBName)
	currentConnectionStatus = Connected
	return connectionString, nil
}
