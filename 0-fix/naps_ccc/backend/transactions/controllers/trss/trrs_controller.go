package trss

import (
	"fmt"
	"net/http"
	"strconv"

	trs "github.com/devpablocristo/transactions/models/trss"
	service "github.com/devpablocristo/transactions/services"
	errors "github.com/devpablocristo/transactions/utils/errors"

	"github.com/gin-gonic/gin"
)

func CreateTrs(c *gin.Context) {
	var trs trs.Trs

	err := c.ShouldBindJSON(&trs)
	if err != nil {
		restErr := errors.BadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		fmt.Println(restErr)
		return
	}

	result, rErr := service.CreateTrs(trs)
	if err != nil {
		//restErr := errors.BadRequestError("error during creation Trs")
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetTrss(c *gin.Context) {
	result, rErr := service.GetTrss()
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetTrs(c *gin.Context) {
	id := c.Param("id")
	result, rErr := service.GetTrs(id)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func UpdateTrs(c *gin.Context) {
	var trs trs.Trs

	err := c.ShouldBindJSON(&trs)
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
			"Trsname":"LedZeppelin",
			"password":"rock",
			"email":"musica@rock.com"
		}

	*/

	uId := c.Param("id")
	result, rErr := service.UpdateTrs(trs, uId)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func DeleteTrs(c *gin.Context) {
	id := c.Param("id")
	del, rErr := service.DeleteTrs(id)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}
	nDel := int(*del)
	r := "Deleted " + strconv.Itoa(nDel) + " document/s."

	c.JSON(http.StatusCreated, r)
}
