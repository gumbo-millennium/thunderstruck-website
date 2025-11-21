package tickets

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/gumbo-millennium/thunderstruck-website/internal/data"
)

type TicketController struct {
	Service TicketService
}

func NewTicketController(service TicketService) TicketController {
	return TicketController{
		Service: service,
	}
}

func (c TicketController) GetTicket(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		render.Render(w, r, ErrInvalidRequest(errors.New("no ticket id given")))
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	ticket, err := c.Service.GetTicketByID(uid)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
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

func ErrInternalError(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal server error.",
		Error:          err.Error(),
	}
}
