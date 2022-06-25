package models

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Name       string `json:"name"`
	Insurance  string `json:"insurance"`
	DOB        string `json:"DOB"`
	Occupation string `json:"occupation"`
	Gender     string `json:"gender"`
}

type MedicalHx struct {
	gorm.Model
	Code      string   `json:"code"`
	Diagnosis string   `json:"diagnosis"`
	PhysExam  string   `json:"physExam"`
	PastHx    []string `gorm:"type:text[]" json:"pastHx"`
}

type Records struct {
	gorm.Model
	Patient Patient     `json:"patient" gorm:"foreignKey:ID"`
	User    User        `json:"user" gorm:"foreignKey:Id"`
	History []MedicalHx `gorm:"foreignKey:ID" json:"history"`
}
