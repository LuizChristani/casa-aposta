package database

import (
	"errors"
)

const (
	URL     = "connection@mydb"
	DBName  = "casa_aposta"
	DBSenha = "123456"
)

func SetDatabase(data map[string] interface{}) error {
	if URL == "" || DBName == "" || DBSenha == "" {
		return errors.New("database: URL, DBName, DBSenha are not valid")
	}
	db := Database{
		URL: URL,
		DBName: DBName,
		DBSenha: DBSenha,
	}
	_, err := ConnectSimulationDatabase(db, data)

	if err != nil {
		return err
	}
	return nil
}