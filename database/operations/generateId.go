package operations

import (
	"casa-aposta/database"
	"encoding/json"
	"os"
)

func GenerateID(tableName string) int64 {
	lastID := GetLastID(tableName)
	return lastID + 1
}

func GetLastID(tableName string) int64 {
	tablePath, err := database.GetTableName(tableName)
	if err != nil {
		return 0
	}

	var existingData []map[string]interface{}
	if _, err := os.Stat(tablePath); err == nil {
		bytes, err := os.ReadFile(tablePath)
		if err != nil {
			return 0
		}
		if len(bytes) > 0 {
			json.Unmarshal(bytes, &existingData)
		}
	}

	if len(existingData) == 0 {
		return 1
	}

	var maxID int64 = 0
	for _, item := range existingData {
		if id, ok := item["id"].(float64); ok {
			if int64(id) > maxID {
				maxID = int64(id)
			}
		}
	}

	return maxID
}

func GetLastCreatedID(tableName string) int64 {
	tablePath, err := database.GetTableName(tableName)
	if err != nil {
		return 0
	}

	var existingData []map[string]interface{}
	if _, err := os.Stat(tablePath); err == nil {
		bytes, err := os.ReadFile(tablePath)
		if err != nil {
			return 0
		}
		if len(bytes) > 0 {
			json.Unmarshal(bytes, &existingData)
		}
	}

	if len(existingData) == 0 {
		return 0
	}

	lastItem := existingData[len(existingData)-1]
	if id, ok := lastItem["id"].(float64); ok {
		return int64(id)
	}

	return 0
}