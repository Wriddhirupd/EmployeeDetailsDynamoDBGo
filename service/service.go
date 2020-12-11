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

func PostItemService(emp models.EmployeeDetail) bool {
	return database.CreateItem(emp, "Employee1")
}

func PatchItemService(m map[string]string) bool {
	return database.UpdateItem(m["tableName"], m["key"],
		m["value"], m["keyToBeUpdated"], m["valueToBeUpdated"])
}
