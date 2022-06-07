package models

type Post struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Text string `json:"post"`
}
