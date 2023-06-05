package controllers

import (
	"backend/Database"
	models "backend/Models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	Database.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
	w.WriteHeader(200)
	fmt.Println(time.Now(), "Get users")
}
