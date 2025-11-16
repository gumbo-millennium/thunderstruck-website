package tickets

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/go-chi/render"
)

type TicketController struct {
	Service TicketService
}

func NewTicketController(service TicketService) TicketController {
	return TicketController{
		Service: service,
	}
}

type PurchaseRequest struct {
	Email string
}

func (p *PurchaseRequest) Bind(r *http.Request) error {
	if p.Email == "" {
		return errors.New("missing required field email")
	}

	if _, err := mail.ParseAddress(p.Email); err != nil {
		return errors.New("given mail is invalid")
	}

	return nil
}

func (c TicketController) Purchase(w http.ResponseWriter, r *http.Request) {
	request := &PurchaseRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	ticket, err := c.Service.NewTicket(request.Email)
	if err != nil {
		render.Render(w, r, ErrInternalError(err))
		return
	}

	fmt.Println(ticket)
}

func (c TicketController) GetById(w http.ResponseWriter, r *http.Request) {}

func (c TicketController) Index(w http.ResponseWriter, r *http.Request) {}

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
