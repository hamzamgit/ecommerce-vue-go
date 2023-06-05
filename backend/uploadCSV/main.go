package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

type Company struct {
	gorm.Model
	Name  string
	Users []User
}

type User struct {
	gorm.Model
	Username  string
	Name      string
	Orders    []Order
	CompanyID uint
	Company   Company
}

type Order struct {
	gorm.Model
	OrderId         string `gorm:"unique"`
	ProductName     string
	DeliveredAmount float64
	TotalAmount     float64
	UserID          uint
	User            User
	OrderDate       time.Time
}

func main() {
	e := godotenv.Load("../.env")
	if e != nil {
		panic("unable to load env")
	}
	ConnectDB()
	UploadSCV()
}

func ConnectDB() {
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("DBPASSWORD")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, gorm_err := gorm.Open(dialect, dbURI)
	if gorm_err != nil {
		fmt.Println(gorm_err)
		panic("Failed to establish connection")
	} else {
		fmt.Println("Database connection is Successfully established")
	}

	// Close connection to database when the main function finishes

	db.AutoMigrate(&Company{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Order{})
	DB = db
}

func UploadSCV() {
	filePath := os.Getenv("CSV_PATH")

	// Open the temporary file
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	defer csvFile.Close()

	// Parse the CSV file
	reader := csv.NewReader(csvFile)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	// Insert the data into the database
	for i, row := range rows {
		if i == 0 {
			continue // Skip header row
		}

		orderID, productName, companyName, username, orderDate, deliveredAmount, totalAmount := row[0], row[1], row[2], row[3], row[4], row[5], row[6]

		// Insert or fetch the company
		company := Company{}
		err = DB.Where("name = ?", companyName).FirstOrCreate(&company).Error
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}

		// Insert or fetch the user
		user := User{}
		err = DB.Where("username = ?", username).FirstOrCreate(&user).Error
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}

		// Parse the amount as a float64
		price, err := strconv.ParseFloat(deliveredAmount, 64)
		if err != nil {
			price = '-'
		}
		totalPrice, err := strconv.ParseFloat(totalAmount, 64)
		if err != nil {
			totalPrice = '-'
		}

		_orderDate, err := time.Parse("2006-01-02T15:04:05.99999-07:00", orderDate)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}

		// Create the order
		order := Order{
			OrderId:         orderID,
			ProductName:     productName,
			UserID:          user.ID,
			OrderDate:       _orderDate,
			DeliveredAmount: price,
			TotalAmount:     totalPrice,
		}

		err = DB.Create(&order).Error
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}
	}
	fmt.Println("Successfully loaded csv!")
}
