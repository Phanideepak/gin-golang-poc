package models

import (
	"github.com/deepak/ems/config"
	_ "github.com/go-sql-driver/mysql"
)

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
	Salary        string `json:"salary"`
	DepartmentId  uint   `json:"department_id"`
}

func GetAllEmployees(employee *[]Employee) (err error) {
	if err := config.DB.Find(employee).Error; err != nil {
		return err
	}
	return nil
}

func GetEmployee(emp *Employee, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(emp).Error; err != nil {
		return err
	}
	return nil
}

func PostEmployee(emp *Employee) (err error) {
	if err := config.DB.Create(emp).Error; err != nil {
		return err
	}
	return nil
}

func EditEmployee(emp *Employee) (err error) {
	config.DB.Save(emp)
	return nil
}

func DeleteEmployee(emp *Employee, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(emp)
	return nil
}

func (b *Employee) TableName() string {
	return "employee"
}
