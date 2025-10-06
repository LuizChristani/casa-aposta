package database

import (
	"casa-aposta/database/connections"
	"errors"
)

const (
	URL     = "connection@mydb"
	DBName  = "casa_aposta"
	DBSenha = "123456"
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

func SetDatabase(data map[string] interface{}) error {
	if URL == "" || DBName == "" || DBSenha == "" {
		return errors.New("database: URL, DBName, DBSenha are not valid")
	}
	db := connections.Database{
		URL: URL,
		DBName: DBName,
		DBSenha: DBSenha,
	}
	_, err := connections.ConnectSimulationDatabaseInitial(db, data)

	if err != nil {
		return err
	}
	return nil
}