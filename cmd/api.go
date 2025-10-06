package main

import (
	"casa-aposta/cmd/endpoints"
	"casa-aposta/database"
	"casa-aposta/database/connections"
	"casa-aposta/repository"
	"casa-aposta/service"
	"casa-aposta/validator"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configuração do banco de dados
	err := database.SetDatabase(map[string]interface{}{
		"operation": connections.OperationConnect,
	})
	if err != nil {
		panic(err)
	}
	
	// Verificação das tabelas
	errorRuntime := validator.RuntimeVerificationTables() 
	if errorRuntime != nil {
		panic(errorRuntime)
	}

	// Injeção de dependências: Repository -> Service -> Handler
	gamesRepo := repository.NewGamesRepository()
	gamesService := service.NewGamesService(gamesRepo)
	handler := endpoints.NewHandler(gamesService)

	// Configuração das rotas
	router.GET("/games", handler.GetAllGames)
	router.POST("/games", handler.CreateGame)

	// Inicialização do servidor
	router.Run(":8080")
}