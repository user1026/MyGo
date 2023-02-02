package routes

import "github.com/gin-gonic/gin"

func RoleRouter(e *gin.Engine) {
	role := e.Group("/role")
	{
		role.POST("/add", AddRole)
		role.POST("/del", DelRole)
		role.POST("/edit", EditRole)
		role.POST("/query", QueryRole)
	}
}

func QueryRole(c *gin.Context) {

}

func EditRole(c *gin.Context) {

}

func DelRole(c *gin.Context) {

}

func AddRole(c *gin.Context) {

}
