package v1

import (
	"net/http"
	"time"

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
	g.GET("/:id/review", r.getById)

	return r
}

type tourGetByIdInput struct {
	Id int `param:"id" validate:"required"`
}

// @Summary		Get tour by id
// @Description	Get tour by id
// @Tags			tours
// @Produce		json
// @Success		200	{object}	entity.Tour
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/api/v1/tours/{id} [get]
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

	return c.JSON(http.StatusOK, tour)
}

type ToursSearch struct {
	FromName  string    `query:"fromName" validate:"omitempty"`
	ToName    string    `query:"toName" validate:"omitempty"`
	When      time.Time `query:"when"`
	NightsCnt int       `query:"nightsCnt" validate:"omitempty"`
	Adults    int       `query:"adults" validate:"omitempty"`
	Childrens int       `query:"childrens" validate:"omitempty"`
}

type ToursFilter struct {
	TourType   string `query:"tourType" validate:"omitempty"`
	PriceFrom  string `query:"priceFrom" validate:"omitempty"`
	PriceTo    string `query:"priceTo" validate:"omitempty"`
	Rating     int    `query:"rating" validate:"omitempty"`
	Guaranteed bool   `query:"guaranteed" validate:"omitempty"`
	Features   Features
	AgeGroupId int `query:"ageGroupId" validate:"omitempty"`
	Difficulty int `query:"difficulty" validate:"omitempty"`
	Comfort    int `query:"comfort" validate:"omitempty"`
	FoodId     int `query:"foodId" validate:"omitempty"`
}

type Features struct {
	Feature1 string `query:"feature1" validate:"omitempty"`
	Feature2 string `query:"feature2" validate:"omitempty"`
	Feature3 string `query:"feature3" validate:"omitempty"`
}

type toursSearchAndFilterInput struct {
	ToursSearch
	ToursFilter
}

// @Summary		Get tours according to search and filter params
// @Description	Get tours according to search and filter params
// @Tags			tours
// @Produce		json
// @Param			fromName	query	string	false	"From location name"
// @Param			toName		query	string	true	"To location name"
// @Param			when		query	string	true	"Date of the tour"
// @Param			nightsCnt	query	int		true	"Number of nights in the tour"
// @Param			adults		query	int		true	"Number of adults"
// @Param			childrens	query	int		true	"Number of children"
// @Param			tourType	query	string	true	"Type of the tour"
// @Param			priceFrom	query	string	true	"Minimum price"
// @Param			priceTo		query	string	true	"Maximum price"
// @Param			rating		query	int		true	"Minimum rating"
// @Param			guaranteed	query	boolean	true	"Guaranteed availability"
// @Param			feature1	query	string	true	"Feature 1"
// @Param			feature2	query	string	true	"Feature 2"
// @Param			feature3	query	string	true	"Feature 3"
// @Param			ageGroupId	query	int		true	"Age group ID"
// @Param			difficulty	query	int		true	"Tour difficulty level"
// @Param			comfort		query	int		true	"Comfort level"
// @Param			foodId		query	int		true	"Food ID"
// @Produce		json
// @Success		200	{object}	[]entity.SimplifiedTourView
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/api/v1/tours [get]
func (r *tourRoutes) getMany(c echo.Context) error {
	var input toursSearchAndFilterInput

	if err := c.Bind(&input); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid query parameters")
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

	simplifiedTours := make([]entity.SimplifiedTourView, len(tours))
	for _, tour := range tours {
		simplifiedTours = append(simplifiedTours, entity.SimplifyingTour(tour))
	}

	return c.JSON(http.StatusOK, simplifiedTours)
}
