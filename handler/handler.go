package handler

import (
	"EmployeeDetailsGoDynamoDB/models"
	"EmployeeDetailsGoDynamoDB/service"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var emp []models.EmployeeDetail

func LoadJSON() {

	empData, err := os.Open(".//employee.json")
	defer empData.Close()
	if err != nil {
		fmt.Println("Could not open the moviedata.json file", err.Error())
		os.Exit(1)
	}
	fmt.Println()
	err = json.NewDecoder(empData).Decode(&emp)
	if err != nil {
		fmt.Println("Could not decode the moviedata.json data", err.Error())
		os.Exit(1)
	}

}

func Validate(uid string) (result bool) {
	result = false
	for _, em := range emp {
		if em.UserId == uid {
			result = true
			break
		} else if em.Region == uid {
			result = true
			break
		}
	}
	return result
}

func GetItemHandler(c *gin.Context) {

	userId := c.Param("userId")

	if Validate(userId) {
		res := service.GetItemService(userId)

		c.JSON(200, gin.H{
			userId: res,
		})
	} else {
		c.JSON(200, gin.H{
			userId: "Not found!",
		})
	}
}

func GetAllHandler(c *gin.Context) {

	res := service.GetAllItemService("CA")
	c.JSON(200, gin.H{
		"employees": res,
	})
}
