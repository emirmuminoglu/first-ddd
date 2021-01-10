package interfaces

import (
	"bytes"

	"github.com/emirmuminoglu/first-ddd/application"
	"github.com/emirmuminoglu/first-ddd/domain"
	"github.com/emirmuminoglu/first-ddd/interfaces/request"
	"github.com/emirmuminoglu/first-ddd/interfaces/response"

	"github.com/emirmuminoglu/lloyd"
)

var jsonContentType = []byte("application/json")

type Handler struct {
	Interactor application.UserInteractor
}

func (h Handler) Init(r lloyd.Router) {
	r.POST("/user/register", h.register)
	//	r.GET("/user/{id}", h.getUser)
}

func (h Handler) register(ctx *lloyd.Ctx) {
	req := request.AcquireRegister()
	defer request.ReleaseRegister(req)

	if !bytes.Equal(ctx.Request.Header.ContentType(), jsonContentType) {
		ctx.SetStatusCode(415)
		return
	}

	err := req.UnmarshalJSON(ctx.PostBody())
	if err != nil {
		ctx.SetStatusCode(400)

		return
	}

	err = h.Interactor.CreateUser(ctx, req.EMail, req.Password)
	if err != nil {
		resp := response.AcquireRegister()
		defer response.ReleaseRegister(resp)

		switch err {
		case domain.ErrHashError:
			resp.Reason = "server error"
			ctx.SetStatusCode(500)
		case domain.ErrInvalidEMail:
			resp.Reason = err.Error()
			ctx.SetStatusCode(400)
		case domain.ErrInvalidPassword:
			resp.Reason = err.Error()
			ctx.SetStatusCode(400)
		case application.ErrDuplicateKey:
			resp.Reason = "email is in use"
			ctx.SetStatusCode(400)
		default:
			resp.Reason = "unexpected error"
			ctx.SetStatusCode(500)
		}

		ctx.JSONInterfaceResponse(resp)
		return
	}
}
