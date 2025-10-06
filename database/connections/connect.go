package connections

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// Tipos locais para evitar ciclo de importação
type Database struct {
	URL     string
	DBName  string
	DBSenha string
}

type DatabaseObj struct {
	Database struct {
		URL     string `json:"url"`
		DBName  string `json:"dbname"`
		DBSenha string `json:"dbsenha"`
	} `json:"name"`
}

type ConnectionStatus int
type DatabaseOperation string

// Enum para Status de Conexão
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

// Variável global para controlar o status da conexão
var currentConnectionStatus = Disconnected

func ConnectSimulationDatabaseInitial(db Database, data map[string]interface{}) (string, error) {

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
		jsonFile, err = os.Open("database/database/database.json")
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
	currentConnectionStatus = Disconnected
	return connectionString, nil
}

func ConnectDatabase(db Database) (string, error) {
	if currentConnectionStatus != Disconnected {
		return "", errors.New("database: connection is not disconnected")
	}

	if db.URL == "" || db.DBName == "" || db.DBSenha == "" {
		return "", errors.New("database: URL, DBName, DBSenha are required")
	}
	
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", db.DBName, db.DBSenha, db.URL, db.DBName)
	currentConnectionStatus = Connected
	return connectionString, nil
}