package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ShinnosukeSuzuki/web-app-develop-golang-todo/entity"
	"github.com/go-playground/validator/v10"
)

type RegisterUser struct {
	Service  RegisterUserService
	Validate *validator.Validate
}

func (ru *RegisterUser) ServeHttp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	u, err := ru.Service.RegisterUser(ctx, b.Name, b.Password, b.Role)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID entity.UserID `json:"id"`
	}{ID: u.ID}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
