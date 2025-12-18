package auth

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	// простая проверка
	if req.Username != "admin" || req.Password != "1234" {
		http.Error(w, "invalid credentials", 401)
		return
	}

	token, _ := Generate(req.Username)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
