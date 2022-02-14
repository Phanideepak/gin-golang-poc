package controllers

import (
	"net/http"
	"strconv"

	"github.com/deepak/ems/api/dto"
	"github.com/deepak/ems/api/models"
	"github.com/deepak/ems/utils"
	"github.com/gin-gonic/gin"
)

func GetAllEmployees(c *gin.Context) {
	var empList []dto.Employee
	err := models.GetAllEmployees(&empList)
	if err != nil {

		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": empList,
		})
	}
}

func GetEmployee(c *gin.Context) {
	var emp models.Employee
	id := c.Param("id")
	err := models.GetEmployee(&emp, id)
	if err != nil {

		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": emp,
		})
	}
}

func PostEmployee(c *gin.Context) {
	var emp models.Employee
	c.BindJSON(&emp)
	err := models.PostEmployee(&emp)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    emp,
			"success": "inserted",
		})
	}
}

func EditEmployee(c *gin.Context) {
	var emp models.Employee
	id := c.Param("id")
	err := models.GetEmployee(&emp, id)
	if err != nil {
		c.JSON(http.StatusNotFound, emp)
	}
	c.BindJSON(&emp)
	err = models.EditEmployee(&emp)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    emp,
			"success": "updated",
		})
	}
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	var emp models.Employee
	err := models.DeleteEmployee(&emp, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "deleted",
		})
	}
}

func GiveEmployeeBonus(c *gin.Context) {
	var emp models.Employee
	id := c.Query("id")
	err := models.GetEmployee(&emp, id)
	if err != nil {
		c.JSON(http.StatusNotFound, emp)
	}

	bonus, _ := strconv.Atoi(c.Query("bonus"))
	utils.UpdateEmployeeSalary(&emp, bonus)
	err = models.EditEmployee(&emp)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    emp,
			"success": "salary hiked",
		})
	}
}
