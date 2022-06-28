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

type MedicalHx struct {
	Code      string   `json:"code"`
	Diagnosis string   `json:"diagnosis"`
	PhysExam  string   `json:"physExam"`
	PastHx    []string `json:"pastHx"`
}

type Records struct {
	gorm.Model
	Patient Patient     `json:"patient"`
	User    User        `json:"user" gorm:"foreignKey:Id"`
	History []MedicalHx `json:"history"`
}
