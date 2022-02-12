package models

import (
	"github.com/deepak/ems/config"
	_ "github.com/go-sql-driver/mysql"
)

type Department struct {
	Id            uint   `json:"id"`
	DeptName      string `json:"dept_name"`
	CreatedBy     string `json:"created_by"`
	LastUpdatedBy string `json:"last_updated_by"`
}

func GetAllDepartments(department *[]Department) (err error) {
	if err = config.DB.Find(department).Error; err != nil {
		return err
	}
	return nil
}

func PostDepartment(department *Department) (err error) {
	if err = config.DB.Create(department).Error; err != nil {
		return err
	}
	return nil
}

func GetDepartment(department *Department, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(department).Error; err != nil {
		return err
	}
	return nil
}

func EditDepartment(department *Department, id string) (err error) {
	config.DB.Save(department)
	return nil
}

func DeleteDepartment(department *Department, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(department)
	return nil
}

func (b *Department) TableName() string {
	return "department"
}
