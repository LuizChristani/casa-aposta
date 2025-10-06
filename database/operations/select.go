package operations

import (
	"encoding/json"
	"os"
	"strings"
)

func Select(result interface{}, tableName string) error {
	// Capitalizar a primeira letra para corresponder aos nomes dos arquivos
	fileName := strings.Title(tableName) + ".json"
	dataFile := "database/tables/" + fileName
	if _, err := os.Stat(dataFile); err != nil {
		return nil
	}

	bytes, err := os.ReadFile(dataFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, result)
}