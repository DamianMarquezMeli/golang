package users

import (
	"fmt"
	"net/http"
	"strconv"

	users "github.com/devpablocristo/users/models/users"
	service "github.com/devpablocristo/users/services"
	errors "github.com/devpablocristo/users/utils/errors"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.BadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		fmt.Println(restErr)
		return
	}

	result, rErr := service.CreateUser(user)
	if err != nil {
		//restErr := errors.BadRequestError("error during creation user")
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUsers(c *gin.Context) {
	result, rErr := service.GetUsers()
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	result, rErr := service.GetUser(id)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func UpdateUser(c *gin.Context) {
	var u users.User

	err := c.ShouldBindJSON(&u)
	if err != nil {
		restErr := errors.BadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		fmt.Println(restErr)
		return
	}

	/*
		ejemplo para actualizar

		id: 609a4ad964b4678593e14d6a
		{
			"username":"LedZeppelin",
			"password":"rock",
			"email":"musica@rock.com"
		}

	*/

	uId := c.Param("id")
	result, rErr := service.UpdateUser(u, uId)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	del, rErr := service.DeleteUser(id)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}
	nDel := int(*del)
	r := "Deleted " + strconv.Itoa(nDel) + " document/s."

	c.JSON(http.StatusCreated, r)
}
