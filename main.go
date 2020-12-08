package main

import "EmployeeDetailsGoDynamoDB/database"

func main() {

	//database.CreateTable("Employee1", "userId", "S")

	// database.LoadData(".//employee.json", "Employee1")

	database.ReadAll("region", "CA", "Employee1")

}
