package utils

import (
	"github.com/deepak/ems/api/models"
)

func UpdateEmployeeSalary(emp *models.Employee, bonus_percentage int) {
	emp.Salary = emp.Salary * (uint(bonus_percentage) + 100) / 100

}
