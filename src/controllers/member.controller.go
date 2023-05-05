package controllers

import (
	"net/http"
	"strconv"

	"github.com/AnggaArdhinata/k-stylehub/src/models"
	"github.com/labstack/echo/v4"
)

func GetAllMember(c echo.Context) error {
	result, err := models.GetAllMember()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreMember(c echo.Context) error {
	username := c.FormValue("username")
	gender := c.FormValue("gender")
	skintype := c.FormValue("skintype")
	skincolor := c.FormValue("skincolor")

	result, err := models.StoreMember(username, gender, skintype, skincolor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateMember(c echo.Context) error  {
	id_member := c.FormValue("id_member")
	username := c.FormValue("username")
	gender := c.FormValue("gender")
	skintype := c.FormValue("skintype")
	skincolor := c.FormValue("skincolor")

	idint, err := strconv.Atoi(id_member)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateMember(idint, username, gender, skintype, skincolor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteMember(c echo.Context) error {
	id := c.Param("id_member")

	idconv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteMember(idconv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}