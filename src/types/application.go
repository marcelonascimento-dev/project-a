package types

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type ApplicationContext struct {
	FiberApp     *fiber.App
	DB           *sql.DB
	SessionStore *session.Store
}
