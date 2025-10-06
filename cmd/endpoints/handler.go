package endpoints

import (
	"casa-aposta/contract/games"
	"casa-aposta/models"
	"casa-aposta/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler estrutura que contém os serviços
type Handler struct {
	GamesService service.GamesService
}

// NewHandler cria uma nova instância do handler
func NewHandler(gamesService service.GamesService) *Handler {
	return &Handler{
		GamesService: gamesService,
	}
}

// GetAllGames endpoint para buscar todos os games
func (h *Handler) GetAllGames(c *gin.Context) {
	games, err := h.GamesService.GetAllGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, games)
}

// CreateGame endpoint para criar um novo game
func (h *Handler) CreateGame(c *gin.Context) {
	var game games.CreateGamesRequest
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	

	createdGame, err := h.GamesService.CreateGame(&models.Games{
		Name:        game.Name,
		Description: game.Description,
		MinBet:      game.MinBet,
		MaxBet:      game.MaxBet,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdGame)
}