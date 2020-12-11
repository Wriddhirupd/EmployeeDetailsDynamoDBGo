package main

import (
	"EmployeeDetailsGoDynamoDB/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	handler.LoadJSON()
	res := handler.Validate("CA")
	fmt.Println(res)

	r := gin.Default()

	// http://localhost:8080/employeeDetails/nirani
	r.GET("/employeeDetails/:userId", handler.GetItemHandler)

	// http://localhost:8080/employeeDetails/
	r.GET("/employeeDetails", handler.GetAllHandler)

	r.Run()

	// /:8080/employeedetails -> GET ALL

	// /:8080/employeedetails/userId -> Get item

	// /:8080/employeedetails/userID -> PATCH

	// /:8080/employeedetails -> PUT

	// include checks for UserID for validation 2nd,3rd, for 4th check if UserID already exists,

	// Create /handler/handler.go Create fns to create, modify
	// Create /service/service.go Create DB calls here

}
