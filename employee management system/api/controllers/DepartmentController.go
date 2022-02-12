package controllers

import (
	"net/http"

	"github.com/deepak/ems/api/models"
	"github.com/gin-gonic/gin"
)

func GetDepartments(c *gin.Context) {
	var department []models.Department
	err := models.GetAllDepartments(&department)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": department,
		})
	}
}

func GetDepartment(c *gin.Context) {
	id := c.Param("id")
	var department models.Department
	err := models.GetDepartment(&department, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": department,
		})
	}
}

func PostDepartment(c *gin.Context) {
	var department models.Department
	c.BindJSON(&department)
	err := models.PostDepartment(&department)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    department,
			"success": "inserted",
		})
	}
}

func EditDepartment(c *gin.Context) {
	id := c.Param("id")
	var department models.Department
	err := models.GetDepartment(&department, id)
	if err != nil {
		c.JSON(http.StatusNotFound, department)
	}
	c.BindJSON(&department)
	err = models.EditDepartment(&department, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    department,
			"success": "updated",
		})
	}
}

func DeleteDepartment(c *gin.Context) {
	id := c.Query("id")
	var department models.Department
	err := models.DeleteDepartment(&department, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "deleted",
		})
	}
}
