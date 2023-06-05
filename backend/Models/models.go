package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

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
	OrderId         string
	ProductName     string
	DeliveredAmount float64
	TotalAmount     float64
	UserID          uint
	User            User
	OrderDate       time.Time
}

//func (order *Order) BeforeCreate(tx *gorm.DB) error {
//	order.OrderId = uuid.New().String()
//	return nil
//}

//func (o *Order) FormattedCreatedAt() string {
//	timeAgo := timeutil.TimeAgo(o.CreatedAt)
//	formattedTime := o.CreatedAt.Format("Jan _2") + timeAgo[3:] + ", 3:04PM"
//	return formattedTime
//}
