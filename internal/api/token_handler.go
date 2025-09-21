package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/visakadev/go/internal/store"
	"github.com/visakadev/go/internal/tokens"
	"github.com/visakadev/go/internal/utils"
)

type TokenHandler struct {
	tokenStore store.TokenStore
	userStore  store.UserStore
	logger     *log.Logger
}

type createTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewTokenHandler(tokenStore store.TokenStore, userStore store.UserStore, logger *log.Logger) *TokenHandler {
	return &TokenHandler{
		tokenStore: tokenStore,
		userStore:  userStore,
		logger:     logger,
	}
}

func (h *TokenHandler) HandleCreateToken(w http.ResponseWriter, r *http.Request) {
	var req createTokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		h.logger.Panicf("Error: createTokenResponse: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelop{"error": "invalid request payload"})
		return
	}
	// get user
	user, err := h.userStore.GetUserByUsername(req.Username)
	if err != nil || user == nil {
		h.logger.Printf("Error: GetUserByUsername %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelop{"error": "invalid server error"})
		return
	}
	passwordsDoMatch, err := user.PasswordHash.Matches(req.Password)
	if err != nil {
		h.logger.Panicf("Error: PasswordHash.Matches: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelop{"error": "invalid request payload"})
		return
	}
	if !passwordsDoMatch {
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelop{"error": "invalid credentials"})
		return
	}
	token, err := h.tokenStore.CreateNewToken(user.ID, 24*time.Hour, tokens.ScopeAuth)
	if err != nil {
		h.logger.Panicf("Error: creating token  %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelop{"error": "invalid server error"})
		return
	}
	utils.WriteJSON(w, http.StatusCreated, utils.Envelop{"auth_token": token})

}
