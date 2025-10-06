package operations

import (
	"casa-aposta/database"
	"encoding/json"
	"os"
	"reflect"
)

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

// InsertModel insere um modelo genérico na tabela especificada
func InsertModel(model interface{}, tableName string) error {
	// Obter o caminho da tabela
	tablePath, err := database.GetTableName(tableName)
	if err != nil {
		return err
	}

	// Ler dados existentes
	var existingData []interface{}
	if _, err := os.Stat(tablePath); err == nil {
		bytes, err := os.ReadFile(tablePath)
		if err != nil {
			return err
		}
		if len(bytes) > 0 {
			json.Unmarshal(bytes, &existingData)
		}
	}

	// Gerar ID único
	newID := len(existingData) + 1
	
	// Definir ID no modelo usando reflection
	modelValue := reflect.ValueOf(model).Elem()
	idField := modelValue.FieldByName("ID")
	if idField.IsValid() && idField.CanSet() {
		idField.SetInt(int64(newID))
	}

	// Adicionar novo modelo
	existingData = append(existingData, model)

	// Salvar dados atualizados
	bytes, err := json.MarshalIndent(existingData, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(tablePath, bytes, 0644)
}
