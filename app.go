package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Patient struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Insurance  string    `json:"insurance"`
	DOB        string    `json:"DOB"`
	Occupation string    `json:"occupation"`
	Gender     string    `json:"gender"`
	TimeAdded  time.Time `json:"timeAdded"`
}

type MedicalHx struct {
	Code      string   `json:"code"`
	Diagnosis string   `json:"diagnosis"`
	PhysExam  string   `json:"physExam"`
	PastHx    []string `json:"pastHx"`
}

var patients []Patient

func init() {
	patients = make([]Patient, 0)
	patients = append(patients, Patient{
		ID: xid.New().String(), Name: "fly Boy", Insurance: "323u78r", DOB: "22-06-1994", Occupation: "Lawyer",
		Gender: "Male", TimeAdded: time.Now(),
	})
}

func getPatients(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, patients)
}

func NewPatientHandler(c *gin.Context) {
	var patient Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	patient.ID = xid.New().String()
	patient.TimeAdded = time.Now()
	patients = append(patients, patient)
	c.JSON(http.StatusOK, patient)
}

func UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var patient Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	index := -1
	for i := 0; i < len(patients); i++ {
		if patients[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Patient not found",
		})
		return
	}

	patients[index] = patient
	c.JSON(http.StatusOK, patient)
}

func deletePatientHandler(c *gin.Context) {
	id := c.Param("id")

	index := -1
	for i := 0; i < len(patients); i++ {
		if patients[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Patient with id " + id + " was not found",
		})
	}
	patients = append(patients[:index], patients[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Patient with id " + id + " has been deleted",
	})
}

func getPatientById(c *gin.Context) {
	id := c.Param("id")

	index := -1
	for i := 0; i < len(patients); i++ {
		if patients[i].ID == id {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Patient with id " + id + " was not found",
		})
	}

	c.JSON(http.StatusOK, patients[index])
}

func main() {
	fmt.Println(patients, "ei")
	router := gin.Default()
	router.GET("/patients", getPatients)
	router.GET("/patients/:id", getPatientById)
	router.POST("/patients", NewPatientHandler)
	router.PUT("/patients/:id", UpdatePatient)
	router.DELETE("/patients/:id", deletePatientHandler)
	router.Run("localhost:6600")
}
