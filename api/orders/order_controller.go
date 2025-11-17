package orders

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/mail"
	"strings"

	"github.com/go-chi/render"
	"github.com/gumbo-millennium/thunderstruck-website/internal/data"
)

var (
	ErrInvalidRequestBody error = errors.New("given request body is invalid")
)

type OrderController struct {
	Service       OrderService
}

func NewOrderController(service OrderService) OrderController {
	return OrderController{
		Service:       service,
	}
}

type NewOrderRequest struct {
	Email string
}

func (request *NewOrderRequest) Bind(r *http.Request) error {
	if request.Email == "" {
		return errors.New("missing required field email")
	}

	if _, err := mail.ParseAddress(request.Email); err != nil {
		return errors.New("given mail is invalid")
	}

	return nil
}

func (c OrderController) NewOrder(w http.ResponseWriter, r *http.Request) {
	request := &NewOrderRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	order, err := c.Service.NewOrder(request.Email)
	if err != nil {
		render.Render(w, r, ErrInternalError(err))
		return
	}

	render.Render(w, r, NewOrderReponse(order))
}

func (c OrderController) ConfirmOrder(w http.ResponseWriter, r *http.Request) {
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error(err.Error())
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	// A valid request by Mollie should look like `id=xyz`.
	// Make sure that there will be two items after splitting by `=`.
	split := strings.Split(string(raw), "=")
	if len(split) < 2 {
		slog.Error(ErrInvalidRequestBody.Error())
		render.Render(w, r, ErrInvalidRequest(ErrInvalidRequestBody))
		return
	}

	// The second position in the split should equal the value of the payment id.
	// Within our own order domain we refer to this id as `reference`.
	order, err := c.Service.ConfirmOrderByReference(split[1])
	if errors.Is(err, ErrOrderAlreadyPaid) {
		render.Render(w, r, NewOrderReponse(order))
		return
	} else if err != nil {
		slog.Error(err.Error())
		render.Render(w, r, ErrInternalError(err))
		return
	}

	slog.Info("created new ticket from mollie webhook", "order_id", order.ID, "ticket_id", order.TicketID)
	render.Render(w, r, NewOrderReponse(order))
}

type OrderResponse struct {
	data.Order
}

func NewOrderReponse(order data.Order) *OrderResponse {
	return &OrderResponse{Order: order}
}

func (response *OrderResponse) Render(w http.ResponseWriter, r *http.Request) error {
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
