package endpoints

import (
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