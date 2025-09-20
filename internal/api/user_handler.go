package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"

	"github.com/visakadev/go/internal/store"
	"github.com/visakadev/go/internal/utils"
)

type registerUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

type UserHandler struct {
	userStore store.UserStore
	logger    *log.Logger
}

func NewUserHandler(userStore store.UserStore, logger *log.Logger) *UserHandler {
	return &UserHandler{
		userStore: userStore,
		logger:    logger,
	}
}

func (h *UserHandler) validateRegisterRequest(req *registerUserRequest) error {
	// Add all the necessary constraints
	if req.Username == "" {
		return errors.New("username is required")
	}
	if req.Email == "" {
		return errors.New("email is required")
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(req.Email) {
		return errors.New("invalid email format")
	}
	// Password must be at least 8 characters long
	// contain at least one uppercase letter, one lowercase letter, one digit, and one special character

	if req.Password == "" {
		return errors.New("password is required")
	}
	return nil

}

func (h *UserHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	var req registerUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.Printf("Error: DecodeRegisterUser: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelop{"error": "invalid request sent"})
		return
	}
	err = h.validateRegisterRequest(&req)
	if err != nil {
		h.logger.Printf("Error: ValidateRegisterUser: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelop{"error": err.Error()})
		return
	}
	user := &store.User{
		Username: req.Username,
		Email:    req.Email,
	}
	if req.Bio != "" {
		user.Bio = req.Bio
	}
	// Password handler
	err = user.PasswordHash.Set(req.Password)
	if err != nil {
		h.logger.Printf("ERROR: Hashing password %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelop{"error": "internal server error"})
		return
	}

	err = h.userStore.CreateUser(user)
	if err != nil {
		h.logger.Printf("ERROR: registering user %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelop{"error": "internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelop{"user": user})

}
