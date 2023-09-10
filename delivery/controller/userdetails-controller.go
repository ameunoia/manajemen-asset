package controller

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/usecase"
	"github.com/gin-gonic/gin"
)

type UserDetailsController struct {
	udetailsUC usecase.UserDetailsUsecase
	gin        *gin.Engine
}

func (u *UserDetailsController) SaveUserHandler(c *gin.Context) {
	var udetails model.UserDetails

	if err := u.udetailsUC.NewUserDetails(udetails); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": "Failed to save user details %v"})
		return
	}
}

// define routing in here
func (u *UserDetailsController) Route() {
	//create a group
	ug := u.gin.Group("/app")
	{
		ug.POST("/save-user", u.SaveUserHandler)
	}
}

func NewUserDetailsController(udetails usecase.UserDetailsUsecase, g *gin.Engine) *UserDetailsController {
	return &UserDetailsController{
		udetailsUC: udetails,
		gin:        g,
	}
}
