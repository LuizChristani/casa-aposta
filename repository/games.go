package repository

import (
	"casa-aposta/database/operations"
	"casa-aposta/models"
)

// GamesRepository interface define os métodos do repositório de games
type GamesRepository interface {
	GetAllGames() ([]models.Games, error)
	CreateGame(game *models.Games) (*models.Games, error)
}

// gamesRepository implementa GamesRepository
type gamesRepository struct{}

// NewGamesRepository cria uma nova instância do repositório de games
func NewGamesRepository() GamesRepository {
	return &gamesRepository{}
}

// GetAllGames retorna todos os games do banco de dados
func (r *gamesRepository) GetAllGames() ([]models.Games, error) {
	var games []models.Games
	err := operations.Select(&games, models.AllModelsNames[1])
	if err != nil {
		return nil, err
	}
	return games, nil
}

// CreateGame cria um novo game no banco de dados
func (r *gamesRepository) CreateGame(game *models.Games) (*models.Games, error) {
	err := operations.InsertModel(game, models.AllModelsNames[1])
	if err != nil {
		return nil, err
	}
	return game, nil
}