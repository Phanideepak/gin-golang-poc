package routes

import (
	"github.com/deepak/ems/api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	group1 := r.Group("/department")
	{
		group1.GET("/all", controllers.GetDepartments)
		group1.GET(":id", controllers.GetDepartment)
		group1.POST("add", controllers.PostDepartment)
		group1.PUT("edit/:id", controllers.EditDepartment)
		group1.DELETE("delete", controllers.DeleteDepartment)
	}

	group2 := r.Group("/employee")
	{
		group2.GET("/all", controllers.GetAllEmployees)
		group2.GET(":id", controllers.GetEmployee)
		group2.POST("add", controllers.PostEmployee)
		group2.PUT("edit/:id", controllers.EditEmployee)
		group2.DELETE("delete/:id", controllers.DeleteEmployee)
		group2.PUT("/giveHike", controllers.GiveEmployeeBonus)
	}
	return r
}
