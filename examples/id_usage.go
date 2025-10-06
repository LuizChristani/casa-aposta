package main

import (
	"casa-aposta/database/operations"
	"fmt"
)

func main() {
	// Exemplo de como usar as funções de ID

	// 1. Pegar o último ID criado na tabela Games
	lastID := operations.GetLastID("Games")
	fmt.Printf("Último ID na tabela Games: %d\n", lastID)

	// 2. Pegar o ID do último registro criado (último na lista)
	lastCreatedID := operations.GetLastCreatedID("Games")
	fmt.Printf("ID do último registro criado: %d\n", lastCreatedID)

	// 3. Gerar o próximo ID para uma nova inserção
	nextID := operations.GenerateID("Games")
	fmt.Printf("Próximo ID a ser usado: %d\n", nextID)

	// Exemplo com outras tabelas
	fmt.Println("\n--- Exemplo com tabela User ---")
	userLastID := operations.GetLastID("User")
	fmt.Printf("Último ID na tabela User: %d\n", userLastID)
	
	userNextID := operations.GenerateID("User")
	fmt.Printf("Próximo ID para User: %d\n", userNextID)
}