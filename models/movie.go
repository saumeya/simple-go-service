package models

type Movie struct {
	ID           uint   `json:"id" gorm:"primary_key`
	Title        string `json:"title`
	Poster_Path  string `json:poster_path"`
	Language     string `json:"language":`
	Overview     string `json:"overview"`
	Release_Date string `json:"release_date`
}
