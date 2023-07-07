package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	name     string `json:"name"`
	gameName string `json:"gameName"`
	tagLine  string `json:"tagLine"`
	puuid    string `json:"puuid"`
}
