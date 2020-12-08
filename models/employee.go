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
