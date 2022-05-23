package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
)

type Patient struct {
	ID         string `json:"id"`
	Name       string `json:"title"`
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

var patients []Patient

func init() {
	patients = make([]Patient, 0)
	patients = append(patients, Patient{
		ID: xid.New().String(), Name: "fly Boy", Insurance: "323u78r", DOB: "2/3/2022", Occupation: "Lawyer",
		Gender: "Male",
	})
}

func getPatients(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, patients)
}

func main() {
	fmt.Println(patients, "ei")
	router := gin.Default()
	router.GET("/patients", getPatients)
	router.Run("localhost:6600")

}
