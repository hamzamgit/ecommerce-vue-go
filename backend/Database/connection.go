package Database

import (
	models "backend/Models"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Connect() {
	e := godotenv.Load(".env")
	if e != nil {
		panic("unable to load env")
	}
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("DBPASSWORD")
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println("DATABASE: ", dbURI)

	db, gorm_err := gorm.Open(dialect, dbURI)
	if gorm_err != nil {
		panic("Failed to establish connection")
	} else {
		fmt.Println("Database connection is Successfully established")
	}

	// Close connection to database when the main function finishes

	db.AutoMigrate(&models.Company{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Order{})
	DB = db
}
