package routers

import (
	"net/http"

	"github.com/AnggaArdhinata/k-stylehub/src/controllers"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from routers")
	})

	e.GET("/member", controllers.GetAllMember)
	e.POST("/member", controllers.StoreMember)
	e.PUT("/member", controllers.UpdateMember)
	e.DELETE("/member/:id_member", controllers.DeleteMember)

	return e
}
