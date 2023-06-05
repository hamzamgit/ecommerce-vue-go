package controllers

import (
	"backend/Database"
	models "backend/Models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GenerateCompany(w http.ResponseWriter, r *http.Request) {
	var newCompany models.Company
	json.NewDecoder(r.Body).Decode(&newCompany)
	createdCompany := Database.DB.Create(&newCompany)
	err = createdCompany.Error
	if err != nil {
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			return
		}
	} else {
		json.NewEncoder(w).Encode(&newCompany)
	}
	fmt.Println(time.Now(), "Create Company")
}

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	fmt.Print("db =>", Database.DB)
	var companies []models.Company
	result := Database.DB.Find(&companies)

	if result.Error != nil {
		fmt.Println("Failed to fetch companies:", result.Error)
		return
	}

	json.NewEncoder(w).Encode(&companies)
	w.WriteHeader(200)
	fmt.Println(time.Now(), "Get tasks")
}
