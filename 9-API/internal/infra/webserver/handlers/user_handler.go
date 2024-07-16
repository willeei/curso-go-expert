package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/willbrr.dev/goexpert/9-API/internal/dto"
	"github.com/willbrr.dev/goexpert/9-API/internal/entity"
	"github.com/willbrr.dev/goexpert/9-API/internal/infra/database"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// GetJwtToken godoc
// @Summary     Get JWT token
// @Description Get JWT token
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       request body dto.GetTokenInput true "user credentials"
// @Success     200 {object} dto.GetTokenOutput
// @Failure     400 {object} Error
// @Failure     401 {object} Error
// @Failure     404 {object} Error
// @Router      /users/generate_token [post]
func (h *UserHandler) GetJwtToken(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwt_expires_in").(int)
	var user dto.GetTokenInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "Invalid request body"}
		json.NewEncoder(w).Encode(error)
		return
	}
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: "User not found"}
		json.NewEncoder(w).Encode(error)
		return
	}
	if !u.CheckPassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	_, token, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	accessToken := dto.GetTokenOutput{AccessToken: token}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary     Create a new user
// @Description Create a new user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       request body dto.CreateUserInput true "User object that needs to be created"
// @Success     201
// @Failure     400 {object} Error
// @Failure     500 {object} Error
// @Router      /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "Invalid request body"}
		json.NewEncoder(w).Encode(error)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: "Failed to create user"}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
