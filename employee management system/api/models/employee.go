package models

import (
	"github.com/deepak/ems/api/dto"
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
	Salary        uint   `json:"salary"`
	DepartmentId  uint   `json:"department_id"`
}

func GetAllEmployees(employee *[]dto.Employee) (err error) {

	if err := config.DB.Table("employee").Select("employee.id,employee.first_name, employee.last_name, employee.email, employee.phone_number, employee.age, employee.job, employee.birth_date, employee.joining_date, employee.created_by, employee.last_updated_by, employee.salary,department.dept_name").Joins("join department on employee.department_id = department.id").Find(employee).Error; err != nil {
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

func GetNthHighestSalary(emp *[]dto.Employee, n string) (err error) {
	query := "SELECT employee.id,employee.first_name, employee.last_name, employee.email, employee.phone_number, employee.age, employee.job, employee.birth_date, employee.joining_date, employee.created_by, employee.last_updated_by, employee.salary,department.dept_name from employee join department on employee.department_id = department.id where salary = (select g.salary from (Select e.salary from  employee e order by salary desc limit " + n + ") as g order by salary limit 1);"

	if err := config.DB.Raw(query).Find(&emp).Error; err != nil {
		return err
	}
	return nil
}

func (b *Employee) TableName() string {
	return "employee"
}
