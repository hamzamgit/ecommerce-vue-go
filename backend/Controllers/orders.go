package controllers

import (
	"backend/Database"
	models "backend/Models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order
	json.NewDecoder(r.Body).Decode(&newOrder)
	fmt.Println("=> new order", newOrder)
	createdOrder := Database.DB.Create(&newOrder)
	err = createdOrder.Error
	if err != nil {
		err := json.NewEncoder(w).Encode(err)
		if err != nil {
			return
		}
	} else {
		json.NewEncoder(w).Encode(&newOrder)
	}
	fmt.Println(time.Now(), "Create Order")
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page") // Get the page parameter from the request query string
	if page == "" {
		page = "1"
	}
	pageSize := 5
	search := r.URL.Query().Get("search")        // Get the search parameter from the request query string
	startDate := r.URL.Query().Get("start_date") // Get the start_date parameter from the request query string
	endDate := r.URL.Query().Get("end_date")     // Get the end_date parameter from the request query string

	var orders []models.Order
	var totalCount int64

	// Create a new query
	query := Database.DB.Preload("User").Preload("User.Company").Find(&orders)

	// Apply search filter if provided
	if search != "" {
		query = query.Where("product_name LIKE ?", "%"+search+"%")
	}

	// Apply date range filter if provided
	if startDate != "" && endDate != "" {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	// Count the total number of orders
	query.Count(&totalCount)

	// Calculate the offset based on the current page and page size
	offset, err := strconv.Atoi(page)
	if err != nil {
		// Handle invalid page parameter
		// Return an appropriate response or error message
		return
	}
	offset = (offset - 1) * pageSize

	// Retrieve the paginated orders with search filter applied
	query.Offset(offset).Limit(pageSize).Find(&orders)

	// Create a response object with the paginated orders
	response := struct {
		TotalCount int64          `json:"totalCount"`
		PageSize   int            `json:"pageSize"`
		Page       int            `json:"page"`
		Orders     []models.Order `json:"orders"`
	}{
		TotalCount: totalCount,
		PageSize:   pageSize,
		Page:       offset/pageSize + 1,
		Orders:     orders,
	}

	// Send the response as JSON
	json.NewEncoder(w).Encode(response)

	//Database.DB.Find(&orders)
	//json.NewEncoder(w).Encode(&orders)
	//w.WriteHeader(200)
	//fmt.Println(time.Now(), "Get Orders")
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	params := mux.Vars(r)
	Database.DB.First(&order, params["id"])
	json.NewEncoder(w).Encode(&order)
	fmt.Printf("%s Get Order %s \n", time.Now(), params["id"])
}
