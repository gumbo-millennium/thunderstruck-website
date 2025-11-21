package scanner

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/gumbo-millennium/thunderstruck-website/internal/data"
	"github.com/gumbo-millennium/thunderstruck-website/tickets"
)

var (
	ErrIncorrectTokenGiven error = errors.New("given token is incorrect")
)

type ScannerController struct {
	Token         string
	TicketService tickets.TicketService
}

func NewScannerController(token string, ticketService tickets.TicketService) ScannerController {
	return ScannerController{
		Token:         token,
		TicketService: ticketService,
	}
}

type TokenRequest struct {
	Token string
}

func (request *TokenRequest) Bind(r *http.Request) error {
	if request.Token == "" {
		return errors.New("token may not be empty")
	}

	return nil
}

func (c ScannerController) ValidateToken(w http.ResponseWriter, r *http.Request) {
	request := &TokenRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if request.Token != c.Token {
		render.Render(w, r, ErrIncorrectToken(ErrIncorrectTokenGiven))
		return
	}
}

type ScanTicketRequest struct {
	Token  string
	Ticket string
}

func (request *ScanTicketRequest) Bind(r *http.Request) error {
	if request.Token == "" {
		return errors.New("token may not be empty")
	}

	if request.Ticket == "" {
		return errors.New("ticket may not be empty")
	}

	return nil
}

func (c ScannerController) ScanTicket(w http.ResponseWriter, r *http.Request) {
	request := &ScanTicketRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if request.Token != c.Token {
		render.Render(w, r, ErrIncorrectToken(ErrIncorrectTokenGiven))
		return
	}

	ticket, err := c.TicketService.GetTicketByValue(request.Ticket)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	ticket, err = c.TicketService.Repository.UpdateTicket(context.Background(), data.UpdateTicketParams{
		ID:    ticket.ID,
		Type:  ticket.Type,
		State: data.TicketStateUsed,
		Email: ticket.Email,
		Value: ticket.Value,
	})
	if err != nil {
		render.Render(w, r, ErrInternalError(err))
		return
	}

	render.Render(w, r, NewTicketReponse(ticket))
}

type TicketResponse struct {
	data.Ticket
}

func NewTicketReponse(order data.Ticket) *TicketResponse {
	return &TicketResponse{Ticket: order}
}

func (response *TicketResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type ErrResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
	Error          string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		Error:          err.Error(),
	}
}

func ErrIncorrectToken(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 401,
		StatusText:     "Incorect token",
		Error:          err.Error(),
	}
}

func ErrInternalError(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal server error.",
		Error:          err.Error(),
	}
}
