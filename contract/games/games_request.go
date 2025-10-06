package games

type CreateGamesRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MinBet      float64 `json:"min_bet"`
	MaxBet      float64 `json:"max_bet"`
}