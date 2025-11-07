package handlers

import (
	"encoding/json"
	"fmt"
	"jwt-token/Model"
	"jwt-token/security"
	"net/http"
)

func LoginHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var user Model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		return
	}
	if user.Username == "isco_naifa" && user.Password == "naifa245" {
		tokenString, err := security.CreateToken(user.Username)
		if err != nil {
			http.Error(writer, fmt.Sprintf("The user not found%v", err), http.StatusNotFound)
		}
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, tokenString)
		return
	} else {
		http.Error(writer, "The username or password is incorrect", http.StatusUnauthorized)
	}

}

func ProtectedHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	tokenString := request.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(writer, "No token provided", http.StatusUnauthorized)
		return
	}
	tokenString = tokenString[len("Bearer "):]
	err := security.VerifyToken(tokenString)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(writer, "The token is invalid")
		return
	}
	fmt.Fprintf(writer, "Welcome to the protected area")
}
