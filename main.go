package main

import (
	"EmployeeDetailsGoDynamoDB/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// http://localhost:8080/employeeDetails/nirani
	r.GET("/employeeDetails/:userId", handler.GetItemHandler)

	// http://localhost:8080/employeeDetails/
	r.GET("/employeeDetails", handler.GetAllHandler)

	// http://localhost:8080/employeeDetails/
	r.POST("employeeDetails/", handler.PostItemHandler)

	// http://localhost:8080/employeeDetails/
	r.PATCH("employeeDetails/", handler.PatchItemHandler)

	r.Run()

	// /:8080/employeedetails -> GET ALL

	// /:8080/employeedetails/userId -> Get item

	// /:8080/employeedetails/userID -> PATCH

	// /:8080/employeedetails -> PUT

	// include checks for UserID for validation 2nd,3rd, for 4th check if UserID already exists,

	// Create /handler/handler.go Create fns to create, modify
	// Create /service/service.go Create DB calls here

}
