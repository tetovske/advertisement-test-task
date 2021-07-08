package models

type Advertisement struct {
	Id 			uint `json:"-"`
	Title 		string `json:"title"`
	Description string `json:"description"`
	Price		uint `json:"price"`
}
