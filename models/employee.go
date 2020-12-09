package models

type EmployeeDetails struct {
	Employees []EmployeeDetail `json:"employees"`
}

//omit empty what will happen if we dont add omit empty
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

type EmployeeDetailDynamoDB struct {
	EmployeeDetail
	Pkey string `json:"pkey,omitempty"`
}
