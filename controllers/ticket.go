package controllers

import (
	"fmt"

	"github.com/labstack/echo"
)

func BuyTicket(c echo.Context) error {
	fmt.Println(10000)
	return nil
}

// type TicketHandler struct {
// 	ticketStore ticket.Store
// }

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
