package service

import (
	"EmployeeDetailsGoDynamoDB/database"
	"EmployeeDetailsGoDynamoDB/models"
)

func GetItemService(userId string) models.EmployeeDetailDynamoDB {
	return database.GetItem("Employee1", "pkey", userId)
}

func GetAllItemService(region string) []models.EmployeeDetailDynamoDB {
	return database.ReadAll("region", region, "Employee1")
}
