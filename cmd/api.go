package main

import (
	"casa-aposta/database"
	"casa-aposta/runtime"
	"fmt"
)

func main() {
	// Primeiro conecta ao banco
	err := database.SetDatabase(map[string]interface{}{
		"operation": database.OperationConnect,
	})

	if err != nil {
		panic(err)
	}
	
	// Depois verifica as tabelas (agora que está conectado)
	errorRuntime := runtime.RuntimeVerificationTables() 

	if errorRuntime != nil {
		panic(errorRuntime)
	}

	data := map[string]interface{}{
		"name": "Tigrinho",
		"tipo": "esports", 
		"score": 100,
	}

	err = database.Insert(data)
	if err != nil {
		fmt.Println("Error inserting data:", err)
		return
	}

	fmt.Println("Data inserted successfully")
}