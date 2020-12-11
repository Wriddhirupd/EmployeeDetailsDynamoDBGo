# EmployeeDetailsDynamoDBGo

# 1. API Call Structure: 

  main.go -> handler.go -> service.go -> database.go 

# 2. Data Return Structure

  main.go <- handler.go <- service.go <- database.go
  
# GET ALL 
http://localhost:8080/employeeDetails/
{"employees":[{"userId":"thanks","jobTitleName":"Program
Directory","firstName":"Tom","lastName":"Hanks","preferredFullName":"Tom
Hanks","employeeCode":"E3","region":"CA","phoneNumber":"408-2222222","emailAddress":"tomhanks@gmail.com","pkey":"thanks"},{"userId":"rirani","jobTitleName":"Developer","firstName":"Romin","lastName":"Irani","preferredFullName":"Romin
Irani","employeeCode":"E1","region":"CA","phoneNumber":"408-1234567","emailAddress":"romin.k.irani@gmail.com","pkey":"rirani"},{"userId":"wriddhirupd","jobTitleName":"Developer","firstName":"wriddhirupdd","lastName":"Dutta","preferredFullName":"wriddhirupdd
dutta","employeeCode":"E1","region":"CA","phoneNumber":"408-1234567","emailAddress":"romin.k.irani@gmail.com","pkey":"wriddhirupd"},{"userId":"nirani","jobTitleName":"Developer","firstName":"Neil","lastName":"Irani","preferredFullName":"Neil
Irani","employeeCode":"E2","region":"CA","phoneNumber":"408-1111111","emailAddress":"neilrirani@gmail.com","pkey":"nirani"}]}

# GET ITEM
http://localhost:8080/employeeDetails/nirani
  {
    "nirani":{
              "userId": "nirani",
              "jobTitleName": "Developer",
              "firstName": "Neil",
              "lastName": "Irani",
              "preferredFullName": "Neil Irani",
              "employeeCode": "E2",
              "region": "CA",
              "phoneNumber": "408-1111111",
              "emailAddress": "neilrirani@gmail.com",
              "pkey": "nirani"
          }
  }

# POST ITEM
http://localhost:8080/employeeDetails/
  1. Body
      {
    "userId":"wriddhirupdd1",
    "jobTitleName":"Developer",
    "firstName":"wriddhirupdd",
    "lastName":"Dutta",
    "preferredFullName":"wriddhirupdd dutta",
    "employeeCode":"E1",
    "region":"CA",
    "phoneNumber":"408-1234567",
    "emailAddress":"romin.k.irani@gmail.com"
    }
    
 2. Output
     {
        "item": {
            "userId": "wriddhirupdd1",
            "jobTitleName": "Developer",
            "firstName": "wriddhirupdd",
            "lastName": "Dutta",
            "preferredFullName": "wriddhirupdd dutta",
            "employeeCode": "E1",
            "region": "CA",
            "phoneNumber": "408-1234567",
            "emailAddress": "romin.k.irani@gmail.com"
        },
        "message": "Data Posted in Employee1 !"
    }

# PATCH ITEM
http://localhost:8080/employeeDetails/
  1. Body 
    {
      "tableName": "Employee1",
      "key": "pkey",
      "value": "wriddhirupd",
      "keyToBeUpdated": "lastName",
      "valueToBeUpdated": "dutta44"
    }
  
  2. Output
      {
        "updated item": {
            "userId": "wriddhirupd",
            "jobTitleName": "Developer",
            "firstName": "neel44",
            "lastName": "dutta44",
            "preferredFullName": "wriddhirupdd dutta",
            "employeeCode": "E1",
            "region": "CA",
            "phoneNumber": "408-1234567",
            "emailAddress": "romin.k.irani@gmail.com",
            "pkey": "wriddhirupd"
        }
      }
