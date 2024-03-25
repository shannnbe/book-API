package models

type BookModel struct {
	Id          int		`json:"id" gorm:"primaryKey"`
	Title       string	`json:"title"`
	Description string 	`json:"description"`
	Qty         int		`json:"quantity"`
}