package v1

import (
	"net/http"
	"net/url"

	_ "github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/service"
	"github.com/labstack/echo/v4"
)

type cityRoutes struct {
	cityService service.City
}

func newCityRoutes(g *echo.Group, cityService service.City) *cityRoutes {
	r := &cityRoutes{
		cityService: cityService,
	}

	g.GET("/:prefix", r.getMany)

	return r
}

type cityGetManyLimitInput struct {
	Prefix string `param:"prefix" validate:"required"`
	Limit  int    `query:"limit"`
}

// @Summary		Get city by prefix
// @Description	Get city by prefix
// @Tags			cities
// @Produce		json
// @Param			prefix	path		string	true	"City name prefix"
// @Param			limit	query		int		false	"Output limit"
// @Success		200		{object}	[]entity.City
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/api/v1/cities/{prefix} [get]
func (r *cityRoutes) getMany(c echo.Context) error {
	var input cityGetManyLimitInput

	if err := c.Bind(&input); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	prefix, err := url.QueryUnescape(input.Prefix)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "incorrect parameter")
		return err
	}

	// default limit value
	if input.Limit == 0 {
		input.Limit = 10
	}

	input.Prefix = prefix
	cities, err := r.cityService.GetMany(c.Request().Context(), input.Prefix, input.Limit)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusOK, cities)
}
