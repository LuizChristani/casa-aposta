package service

import (
	"casa-aposta/models"
	"casa-aposta/repository"
	"errors"
)

// GamesService interface define os métodos do serviço de games
type GamesService interface {
	GetAllGames() ([]models.Games, error)
	GetGameByID(id int) (*models.Games, error)
	ValidateGame(game *models.Games) error
}

// gamesService implementa GamesService
type gamesService struct {
	repository repository.GamesRepository
}

// NewGamesService cria uma nova instância do serviço de games
func NewGamesService(repo repository.GamesRepository) GamesService {
	return &gamesService{
		repository: repo,
	}
}

// GetAllGames retorna todos os games com validações de negócio
func (s *gamesService) GetAllGames() ([]models.Games, error) {
	games, err := s.repository.GetAllGames()
	if err != nil {
		return nil, err
	}

	// Lógica de negócio: filtrar games inválidos
	var validGames []models.Games
	for _, game := range games {
		if s.ValidateGame(&game) == nil {
			validGames = append(validGames, game)
		}
	}

	return validGames, nil
}

// GetGameByID retorna um game específico por ID
func (s *gamesService) GetGameByID(id int) (*models.Games, error) {
	games, err := s.repository.GetAllGames()
	if err != nil {
		return nil, err
	}

	for _, game := range games {
		if game.ID == id {
			if err := s.ValidateGame(&game); err != nil {
				return nil, err
			}
			return &game, nil
		}
	}

	return nil, errors.New("game not found")
}

// ValidateGame valida as regras de negócio de um game
func (s *gamesService) ValidateGame(game *models.Games) error {
	if game == nil {
		return errors.New("game cannot be nil")
	}

	if game.Name == "" {
		return errors.New("game name cannot be empty")
	}

	if game.MinBet < 0 {
		return errors.New("minimum bet cannot be negative")
	}

	if game.MaxBet < 0 {
		return errors.New("maximum bet cannot be negative")
	}

	if game.MinBet > game.MaxBet {
		return errors.New("minimum bet cannot be greater than maximum bet")
	}

	return nil
}