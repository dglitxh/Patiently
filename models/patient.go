package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	Name       string `json:"name"`
	Insurance  string `json:"insurance"`
	DOB        string `json:"DOB"`
	Occupation string `json:"occupation"`
	Gender     string `json:"gender"`
}
