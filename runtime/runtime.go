package runtime

import (
	"casa-aposta/database"
	"casa-aposta/models"
	"reflect"
)

func RuntimeVerificationTables() error {
	for _, model := range models.AllModels {
		modelType := reflect.TypeOf(model)
		modelName := modelType.Name()
		
		_, err := database.GetTableName(modelName)
		if err != nil {
			return err
		}
	}
	return nil
}