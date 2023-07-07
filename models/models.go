package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"text;not null;"`
	GameName string `json:"gameName" gorm:"text;default:null"`
	TagLine  string `json:"tagLine" gorm:"text;default:null"`
	Puuid    string `json:"puuid" gorm:"text;"`
}
