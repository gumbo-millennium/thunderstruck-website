package scanner

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type ScannerController struct {
	Token string
}

func NewScannerController(token string) ScannerController {
	return ScannerController{
		Token: token,
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
		render.Render(w, r, ErrIncorrectToken(errors.New("given token is incorrect")))
		return
	}
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
