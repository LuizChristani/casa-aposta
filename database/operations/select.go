package operations

import (
	"encoding/json"
	"os"
)

func Select(result interface{}) error {
	dataFile := "database/tables/games_data.json"
	if _, err := os.Stat(dataFile); err != nil {
		// Se o arquivo não existe, retorna slice vazio
		return nil
	}

	bytes, err := os.ReadFile(dataFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, result)
}