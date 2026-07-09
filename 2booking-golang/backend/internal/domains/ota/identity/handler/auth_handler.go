package handler

import (
	"encoding/json"
	"net/http"

	"1booking-golang/backend/internal/domains/ota/identity/model"
	"1booking-golang/backend/internal/domains/ota/identity/service"
)

// AuthHandler handles authentication-related HTTP requests.
type AuthHandler struct {
	authService *service.AuthService
}

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login handles user login requests.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Implementation will go here
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Login endpoint"}`))
}

// Logout handles user logout requests.
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Implementation will go here
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Logout endpoint"}`))
}

// RefreshToken handles token refresh requests.
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Implementation will go here
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Refresh token endpoint"}`))
}

// Register handles user registration requests.
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.authService.Register(r.Context(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
		"user_id": user.ID,
	})
}