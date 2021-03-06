package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tuananhvp25081995/bitkingreturn_golang/controllers"
	"github.com/tuananhvp25081995/bitkingreturn_golang/db"
	"gopkg.in/go-playground/validator.v9"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// Create creates the gin engine with all routes.
func Create(db *db.DB) *echo.Echo {

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// ticketStore := store.NewTicketStore(db)
	// ticketHandler := handler.NewTicketHandler(ticketStore)

	e.POST("/ticket/buy", controllers.BuyTicket)

	return e
}
