package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tuananhvp25081995/bitkingreturn_golang/models"
	"github.com/tuananhvp25081995/bitkingreturn_golang/utils"
)

type TicketHandler struct {
	ticketStore ticket.Store
}

// func NewTicketHandler(ts ticket.Store) *TicketHandler {
// 	return &TicketHandler{
// 		ticketStore: ts,
// 	}
// }

// func (th *TicketHandler) BuyTicket(c echo.Context) error {
// 	var t model.Ticket
// 	req := &api.RequestBuyTicket{}
// 	if err := req.Bind(c, &t); err != nil {
// 		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
// 	}
// 	if result := th.ticketStore.GetByID(t.ID); result != nil {
// 		return c.JSON(http.StatusCreated, api.BuyTicketResponse(&t))
// 	} else {
// 		return c.JSON(http.StatusUnprocessableEntity, errors.New("Ticket does not exist"))
// 	}

// }
