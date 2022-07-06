package db

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	//ID                uint `gorm:"auto_increment"` gorm:"primarykey"
	//CreatedAt         time.Time
	//UpdatedAt         time.Time
	gorm.Model
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	OrderUID          string         `gorm:"unique" db:"order_uid" json:"order_uid" validate:"required"`
	TrackNumber       string         `db:"track_number" json:"track_number"`
	Entry             string         `db:"entry" json:"entry"`
	Delivery          Delivery       `db:"delivery" json:"delivery" gorm:"foreignkey:Order;references:OrderUID"`
	Payment           Payment        `db:"payment" json:"payment" gorm:"foreignkey:Order;references:OrderUID"`
	Items             []Items        `db:"items" json:"items" gorm:"foreignKey:Order;references:OrderUID"`
	Locale            string         `db:"locale" json:"locale"`
	InternalSignature string         `db:"internal_signature" json:"internal_signature"`
	CustomerId        string         `db:"customer_id" json:"customer_id"`
	DeliveryService   string         `db:"delivery_service" json:"delivery_service"`
	Shardkey          string         `db:"shardkey" json:"shardkey"`
	SmId              int            `db:"sm_id" json:"sm_id"`
	DateCreated       time.Time      `db:"date_created" json:"date_created"`
	OofShard          string         `db:"oof_shard" json:"oof_shard"`
}

type Delivery struct {
	gorm.Model
	Name    string `json:"name" db:"name"`
	Phone   string `json:"phone"  db:"phone" validate:"e164"`
	Zip     string `json:"zip"  db:"zip"`
	City    string `json:"city" db:"city"`
	Address string `json:"address"  db:"address"`
	Region  string `json:"region" db:"region"`
	Email   string `json:"email"  db:"email" validate:"email"`
	Order   string `db:"order" json:"order"`
}

type Payment struct {
	gorm.Model
	Transaction  string `json:"transaction"  db:"transaction"`
	RequestId    string `json:"request_id" db:"request_id"`
	Currency     string `json:"currency"  db:"currency"`
	Provider     string `json:"provider"  db:"provider"`
	Amount       int    `json:"amount"  db:"amount"`
	PaymentDt    int    `json:"payment_dt"  db:"payment_dt"`
	Bank         string `json:"bank"  db:"bank"`
	DeliveryCost int    `json:"delivery_cost"  db:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total" db:"goods_total"`
	CustomFee    int    `json:"custom_fee" db:"custom_fee"`
	Order        string `db:"order" json:"order"`
}

type Items struct {
	gorm.Model
	ChrtId      int    `json:"chrt_id" db:"chrt_id"`
	TrackNumber string `json:"track_number" db:"track_number"`
	Price       int    `json:"price" db:"price"`
	Rid         string `json:"rid" db:"rid"`
	Name        string `json:"name" db:"name"`
	Sale        int    `json:"sale" db:"sale"`
	Size        string `json:"size" db:"size"`
	TotalPrice  int    `json:"total_price" db:"total_price"`
	NmId        int    `json:"nm_id" db:"nm_id"`
	Brand       string `json:"brand" db:"brand"`
	Status      int    `json:"status" db:"status"`
	Order       string `db:"order" json:"order"`
}
