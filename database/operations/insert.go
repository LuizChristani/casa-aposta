package operations

import (
	"encoding/json"
	"os"
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
