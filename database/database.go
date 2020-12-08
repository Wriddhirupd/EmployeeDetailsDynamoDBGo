package database

import (
	"EmployeeDetailsGoDynamoDB/models"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

var dynamo *dynamodb.DynamoDB

func init() {
	dynamo = createSession()
}

func createSession() *dynamodb.DynamoDB {
	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})))
}

func CreateTable(TableName string, Pkey string, PkeyType string) {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(Pkey),
				AttributeType: aws.String(PkeyType),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(Pkey),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String(TableName),
	}

	result, err := dynamo.CreateTable(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}

// func Load() models.EmployeeDetails {
// 	jsonFile, err := os.Open("employee.json")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Successfully Opened employee.json")
// 	defer jsonFile.Close()

// 	byteValue, _ := ioutil.ReadAll(jsonFile)

// 	var emp models.EmployeeDetails
// 	json.Unmarshal(byteValue, &emp)
// 	return emp
// }

// func JsonToMap(fileName string) map[string]interface{} {
// 	jsonFile, err := os.Open(fileName)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("Successfully Opened ", fileName)

// 	b, err := ioutil.ReadAll(jsonFile)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var result map[string]interface{}
// 	json.Unmarshal([]byte(b), &result)

// 	return result
// }

func LoadData(FileName string, TableName string) {

	moviesData, err := os.Open(FileName)
	defer moviesData.Close()
	if err != nil {
		fmt.Println("Could not open the moviedata.json file", err.Error())
		os.Exit(1)
	}

	var movies []models.EmployeeDetail
	err = json.NewDecoder(moviesData).Decode(&movies)
	if err != nil {
		fmt.Println("Could not decode the moviedata.json data", err.Error())
		os.Exit(1)
	}

	fmt.Printf("%+v", movies)

	for _, em := range movies {

		info, err := dynamodbattribute.MarshalMap(em)
		if err != nil {
			panic(fmt.Sprintf("failed to marshal the employees, %v", err))
		}

		input := &dynamodb.PutItemInput{
			Item:      info,
			TableName: aws.String(TableName),
		}

		_, err = dynamo.PutItem(input)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	}

	fmt.Printf("We have processed %v records\n", len(movies))

}

func ReadAll(Filter string, JsonFilter string, tableName string) {

	filt := expression.Name(Filter).Equal(expression.Value(JsonFilter))

	proj := expression.NamesList(expression.Name("userId"),
		expression.Name("jobTitleName"),
		expression.Name("firstName"),
		expression.Name("lastName"),
		expression.Name("preferredFullName"),
		expression.Name("employeeCode"),
		expression.Name("region"),
		expression.Name("phoneNumber"),
		expression.Name("emailAddress"))

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := dynamo.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	fmt.Println(result.Items)

	numItems := 0

	for _, i := range result.Items {
		item := models.EmployeeDetail{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Which ones had a higher rating than minimum?
		//if item.Rating > minRating {
		// Or it we had filtered by rating previously:
		//   if item.Year == year {
		numItems++

		fmt.Println("User Id: " + item.UserId)
		fmt.Println("JobTitleName: " + item.JobTitleName)
		fmt.Println("FirstName: " + item.FirstName)
		fmt.Println("LastName: " + item.LastName)
		fmt.Println("PreferredFullName: " + item.PreferredFullName)
		fmt.Println("EmployeeCode: " + item.EmployeeCode)
		fmt.Println("Region: " + item.Region)
		fmt.Println("PhoneNumber: " + item.PhoneNumber)
		fmt.Println("EmailAddress: " + item.EmailAddress)
		fmt.Println()
	}
}
