package interfaces

import (
	"github.com/emirmuminoglu/first-ddd/application"
	"github.com/emirmuminoglu/lloyd"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	Interactor application.UserInteractor
}

func (h Handler) Init(r lloyd.Router) {
	r.GET("/user/{id}", h.getUser)
}

func (h Handler) getUser(ctx *lloyd.Ctx) {
	idStr := ctx.UserValue("id").(string)

	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.SetStatusCode(400)

		return
	}

	user, err := h.Interactor.GetUser(ctx, objID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.SetStatusCode(404)
			return
		}

		ctx.SetStatusCode(500)
		return
	}
	ctx.JSONInterfaceResponse(user)

	return
}
