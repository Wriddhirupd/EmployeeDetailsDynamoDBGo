package main

import (
	"EmployeeDetailsGoDynamoDB/database"
	"EmployeeDetailsGoDynamoDB/models"
	"encoding/json"
	"fmt"
	"os"
)

type EmployeeDetail struct {
	UserId            string `json:"userId,omitempty"`
	JobTitleName      string `json:"jobTitleName,omitempty"`
	FirstName         string `json:"firstName,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	PreferredFullName string `json:"preferredFullName,omitempty"`
	EmployeeCode      string `json:"employeeCode,omitempty"`
	Region            string `json:"region,omitempty"`
	PhoneNumber       string `json:"phoneNumber,omitempty"`
	EmailAddress      string `json:"emailAddress,omitempty"`
}

type EmployeeDetail1 struct {
	Pkey              string `json:"pkey,omitempty"`
	UserId            string `json:"userId,omitempty"`
	JobTitleName      string `json:"jobTitleName,omitempty"`
	FirstName         string `json:"firstName,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	PreferredFullName string `json:"preferredFullName,omitempty"`
	EmployeeCode      string `json:"employeeCode,omitempty"`
	Region            string `json:"region,omitempty"`
	PhoneNumber       string `json:"phoneNumber,omitempty"`
	EmailAddress      string `json:"emailAddress,omitempty"`
}

func main() {

	//TO Create TABLE
	// database.CreateTable("Employee1", "pkey", "S")

	//LOAD Data to DynamoDB By each item
	// LoadJSonFile()

	//Read All Items using Scan
	// database.ReadAll("pkey", "nirani", "Employee1")

	//Read An Item using Query
	// pkeyNames := []string{"rirani", "nirani", "thanks"}
	// for _, pkey := range pkeyNames {
	// 	database.ReadQuery("Employee1", "pkey", pkey)
	// }

	//Read An Item using GetItem
	// database.GetItem("Employee1", "pkey", "nirani")

}

//Opens JSON and Parses it to DB LoadData Function
func LoadJSonFile() {
	empData, err := os.Open("employee.json")
	defer empData.Close()
	if err != nil {
		fmt.Println("Could not open the moviedata.json file", err.Error())
		os.Exit(1)
	}

	var emp []models.EmployeeDetail
	err = json.NewDecoder(empData).Decode(&emp)
	if err != nil {
		fmt.Println("Could not decode the moviedata.json data", err.Error())
		os.Exit(1)
	}

	for _, em := range emp {
		database.LoadData(em.UserId, em, "Employee1")
	}

	fmt.Printf("We have processed %v records\n", len(emp))
}
