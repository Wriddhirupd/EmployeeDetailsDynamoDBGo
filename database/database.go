package database

import (
	"EmployeeDetailsGoDynamoDB/models"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
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

func LoadData(pkey string, empData models.EmployeeDetail, tableName string) {

	empdb := models.EmployeeDetailDynamoDB{
		EmployeeDetail: empData,
		Pkey:           pkey,
	}
	fmt.Printf("\n%+v\n", empdb)

	info, err := dynamodbattribute.MarshalMap(empdb)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal the employees, %v", err))
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String(tableName),
	}

	_, err = dynamo.PutItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func ReadAll(Filter string, JsonFilter string, tableName string) {

	filt := expression.Name(Filter).Equal(expression.Value(JsonFilter))

	proj := expression.NamesList(expression.Name("pkey"),
		expression.Name("userId"),
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
		item := models.EmployeeDetailDynamoDB{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		numItems++

		fmt.Println("Pkey: " + item.Pkey)
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

func ReadQuery(tableName string, query string, queryValue string) {

	var queryInput = &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		KeyConditions: map[string]*dynamodb.Condition{
			query: {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(queryValue),
					},
				},
			},
		},
	}

	resp1, err1 := dynamo.Query(queryInput)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		personObj := []models.EmployeeDetailDynamoDB{}
		_ = dynamodbattribute.UnmarshalListOfMaps(resp1.Items, &personObj)
		fmt.Printf("\n%+v\n", personObj)
	}

}

func GetItem(tableName string, key string, value string) {
	result, err := dynamo.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			key: {
				S: aws.String(value),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if result.Item == nil {
		log.Println("Could not find item")
		return
	}

	item := models.EmployeeDetailDynamoDB{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Printf("\n%+v\n", item)

}
