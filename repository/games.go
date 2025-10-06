package repository

import (
	"casa-aposta/database/operations"
	"casa-aposta/models"
)

func GetAllGames() ([]models.Games, error) {
	var games []models.Games
	err := operations.Select(&games, models.AllModelsNames[1])
	if err != nil {
		return nil, err
	}
	return games, nil
}