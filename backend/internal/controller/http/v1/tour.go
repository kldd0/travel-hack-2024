package v1

import (
	"net/http"

	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/service"
	"github.com/labstack/echo/v4"
)

type tourRoutes struct {
	tourService service.Tour
}

func newTourRoutes(g *echo.Group, tourService service.Tour) *tourRoutes {
	r := &tourRoutes{
		tourService: tourService,
	}

	g.GET("/:id", r.getById)
	g.GET("", r.getMany)

	return r
}

type tourGetByIdInput struct {
	Id int `param:"id" validate:"required"`
}

// @Summary Get tour by id
// @Description Get tour by id
// @Tags tours
// @Accept
// @Produce json
// @Success 200 {object} v1.tourRoutes.getById.response
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /api/v1/tours/{id} [get]
func (r *tourRoutes) getById(c echo.Context) error {
	var input tourGetByIdInput

	if err := c.Bind(&input); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	tour, err := r.tourService.GetById(c.Request().Context(), input.Id)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	type response struct {
		Tour entity.Tour `json:"tour"`
	}

	return c.JSON(http.StatusOK, response{
		Tour: tour,
	})
}

type toursSearchAndFilterInput struct {
	Param1 int `query:"param1" validate:"omitempty"`
}

// @Summary Get tours according to search and filter params
// @Description Get tours according to search and filter params
// @Tags tours
// @Param collection query []string false "string collection"
// @Produce json
// @Success 200 {object} v1.tourRoutes.getById.response
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /api/v1/tours [get]
func (r *tourRoutes) getMany(c echo.Context) error {
	var input toursSearchAndFilterInput

	if err := c.Bind(&input); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	tours, err := r.tourService.GetMany(c.Request().Context())
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	type response struct {
		Tours []entity.Tour
	}

	return c.JSON(http.StatusOK, response{
		Tours: tours,
	})
}
