package main

import (
	"casa-aposta/database"
	"casa-aposta/database/connections"
	"casa-aposta/repository"
	"casa-aposta/runtime"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	err := database.SetDatabase(map[string]interface{}{
		"operation": connections.OperationConnect,
	})

	if err != nil {
		panic(err)
	}
	
	errorRuntime := runtime.RuntimeVerificationTables() 

	if errorRuntime != nil {
		panic(errorRuntime)
	}

	router.GET("/games", func (c *gin.Context){
		games, err := repository.GetAllGames()
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(200, gin.H{
			"games": games,
		})
	})
}