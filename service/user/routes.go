package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tanishqtrivedi27/ecom/config"
	"github.com/tanishqtrivedi27/ecom/service/auth"
	"github.com/tanishqtrivedi27/ecom/types"
	"github.com/tanishqtrivedi27/ecom/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /login", h.handleLogin)
	router.HandleFunc("POST /register", h.handleRegister)

	router.HandleFunc("GET /addresses", auth.JWTAuthMiddleware(http.HandlerFunc(h.handleGetAddresses)))
	router.HandleFunc("POST /addresses", auth.JWTAuthMiddleware(http.HandlerFunc(h.handleCreateAddress)))
	// router.HandleFunc("PUT /addresses/{id}", auth.JWTAuthMiddleware(http.HandlerFunc(h.handleUpdateAddress)))
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	//get JSON payload
	var payload types.LoginUserPayLoad
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	//check if the user already exits
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email %v", err))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid password"))
		return
	}

	err = h.store.UpdateLastLogin(u.Id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.Id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//get JSON payload
	var payload types.RegisterUserPayLoad
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	//check if the user already exits
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	//else create new user
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleGetAddresses(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserFromContext(r.Context())

	addresses, err := h.store.GetAddresses(userId)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, addresses)

}

func (h *Handler) handleCreateAddress(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserFromContext(r.Context())

	var payload types.CreateAddressPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateAddress(types.Address{
		UserID:  userId,
		Line1:   payload.Line1,
		Line2:   payload.Line2,
		City:    payload.City,
		Country: payload.Country,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

// func (h *Handler) handleUpdateAddress(w http.ResponseWriter, r *http.Request) {
// }
