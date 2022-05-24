package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"time"
)

type Patient struct {
	ID         string    `json:"id"`
	Name       string    `json:"title"`
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

func main() {
	fmt.Println(patients, "ei")
	router := gin.Default()
	router.GET("/patients", getPatients)
	router.POST("/patients", NewPatientHandler)
	router.Run("localhost:6600")

}
