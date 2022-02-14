package dto

type Employee struct {
	Id            uint   `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	Age           uint   `json:"age"`
	Job           string `json:"job"`
	BirthDate     string `json:"birth_date"`
	JoiningDate   string `json:"joining_date"`
	CreatedBy     string `json:"created_by"`
	LastUpdatedBy string `json:"last_updated_by"`
	Salary        uint   `json:"salary"`
	DeptName      string `json:"department_name"`
}
