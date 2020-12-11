package handler

import (
	"EmployeeDetailsGoDynamoDB/models"
	"EmployeeDetailsGoDynamoDB/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func PostItemHandler(c *gin.Context) {

	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}

	var emp models.EmployeeDetail

	er := json.Unmarshal(value, &emp)
	if er != nil {
		panic(er)
	}

	fmt.Printf("%+v", emp)

	//Check if Item already exists in Employee1 Table
	exist := false
	emps := service.GetAllItemService("CA")
	for _, es := range emps {
		if emp.UserId == es.UserId {
			exist = true
			break
		}
	}

	if !exist {
		service.PostItemService(emp)
		// if !err {
		c.JSON(200, gin.H{
			"message": "Data Posted in Employee1 !",
			"item":    emp,
		})
		// }
	} else {
		c.JSON(200, gin.H{
			"message": "Data Not Posted in Employee1 !",
		})
	}

}

func PatchItemHandler(c *gin.Context) {

	var patchDetails map[string]string

	body := c.Request.Body

	value, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}

	er := json.Unmarshal(value, &patchDetails)
	if er != nil {
		panic(er)
	}

	res := service.PatchItemService(patchDetails)
	if !res {
		c.JSON(200, gin.H{
			"message": "Data Not Updated in Employee1 !",
		})
	} else {
		// emp := service.GetItemService(patchDetails["value"])
		c.JSON(200, gin.H{
			"updated item": service.GetItemService(patchDetails["value"]),
		})
	}

}
