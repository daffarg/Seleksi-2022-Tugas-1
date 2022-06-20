package playerRoutes

import (
	"github.com/gofiber/fiber/v2"
	playerHandler "api/internal/handlers/player"

)	

func SetupPlayerRoutes(router fiber.Router) {
	player := router.Group("/player")

	player.Get("/", playerHandler.GetPlayer) // read all players
	player.Get("/:playerId", playerHandler.GetPlayer) // read player by id
}