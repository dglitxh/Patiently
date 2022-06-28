package models

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Name       string      `json:"name"`
	Insurance  string      `json:"insurance"`
	DOB        string      `json:"DOB"`
	Occupation string      `json:"occupation"`
	Gender     string      `json:"gender"`
	History    []MedicalHx `gorm:"foreignKey:Patient_id" json:"history"`
}

type MedicalHx struct {
	gorm.Model
	Code       string   `json:"code"`
	Diagnosis  string   `json:"diagnosis"`
	PhysExam   string   `json:"physExam"`
	PastHx     []string `gorm:"type:text[]" json:"pastHx"`
	Patient_id uint     `json:"patient_id"`
	User_id    string   `json:"user_id"`
}
