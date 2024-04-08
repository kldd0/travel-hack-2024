package v1

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/repository/repoerrs"
	"github.com/kldd0/travel-hack-2024/internal/service"
	"github.com/labstack/echo/v4"
)

type tourRoutes struct {
	tourService service.Tour
	tourReview  service.Review
	tourOrder   service.Order
}

func newTourRoutes(
	g *echo.Group,
	tourService service.Tour,
	reviewService service.Review,
	orderService service.Order,
) *tourRoutes {
	r := &tourRoutes{
		tourService: tourService,
		tourReview:  reviewService,
		tourOrder:   orderService,
	}

	g.GET("", r.getMany)
	g.GET("/hot", r.getHot)
	g.GET("/recommendations", r.getRecommendations)

	g.GET("/:id", r.getById)
	g.POST("/:id/order", r.makeOrder)

	return r
}

type tourGetByIdInput struct {
	Id int `param:"id" validate:"required"`
}

// @Summary		Get tour by id
// @Description	Get tour by id
// @Tags			tours
// @Produce		json
// @Param			id	path		int	true	"Tour id"
// @Success		200	{object}	entity.DTOTour
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
		if errors.Is(err, repoerrs.ErrNotFound) {
			ErrorResponse(c, http.StatusNotFound, err.Error())
			return err
		}

		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	//получаем отзывы
	reviews, err := r.tourReview.GetAllByTourId(c.Request().Context(), tour.Id)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	tour.Reviews = reviews

	return c.JSON(http.StatusOK, tour)
}

type ToursSearch struct {
	FromName  string    `query:"from_name"`
	ToName    string    `query:"to_name"`
	WhenDate  time.Time `query:"when"`
	NightsCnt int       `query:"nights_count"`
	Adults    int       `query:"adults"`
	Childrens int       `query:"childrens"`
}

type ToursFilter struct {
	Tags       []string `query:"tags"`
	PriceFrom  string   `query:"price_from"`
	PriceTo    string   `query:"price_to"`
	Rating     int      `query:"rating"`
	Guaranteed bool     `query:"guaranteed"`
	Features   Features
	AgeGroupId int `query:"age_group"`
	Difficulty int `query:"difficulty"`
	Comfort    int `query:"comfort"`
	FoodId     int `query:"food_id"`
}

type Features struct {
	WithFlight       bool `query:"with_flight"`
	WithAccomodation bool `query:"with_acc"`
	WithFood         bool `query:"with_food"`
	DayOff           bool `query:"day_off"`
	LowCost          bool `query:"low_cost" `
}

func ToFilters(s ToursSearch, f ToursFilter) map[string]interface{} {
	result := make(map[string]interface{})

	if s.ToName != "" {
		result["location"] = s.ToName
	}

	if !s.WhenDate.IsZero() {
		result["date"] = s.WhenDate
	}

	if s.NightsCnt != 0 {
		result["nights_count"] = s.NightsCnt
	}
	// result["tags"] = f.Tags

	return result
}

// @Summary		Get tours according to search and filter params
// @Description	Get tours according to search and filter params
// @Tags			tours
// @Produce		json
// @Param			from_name		query	string	false	"From location name"
// @Param			to_name			query	string	false	"To location name"
// @Param			when			query	string	false	"Date of the tour"
// @Param			nights_count	query	int		false	"Number of nights in the tour"
// @Param			adults			query	int		false	"Number of adults"
// @Param			childrens		query	int		false	"Number of children"
// @Param			tags			query	string	false	"Tags of the tour, multiple values separated by commas"
// @Param			price_from		query	string	false	"Minimum price"
// @Param			price_to		query	string	false	"Maximum price"
// @Param			rating			query	int		false	"Minimum rating"
// @Param			guaranteed		query	boolean	false	"Guaranteed availability"
// @Param			with_flight		query	bool	false	"Flight is included"
// @Param			with_acc		query	bool	false	"Accomodation is included"
// @Param			with_food		query	bool	false	"Nutrition is included""
// @Param			day_off			query	bool	false	"The tour takes place on a weekend"
// @Param			low_cost		query	bool	false	"Low cost tour"
// @Param			age_group		query	int		false	"Age group ID"
// @Param			difficulty		query	int		false	"Tour difficulty level"
// @Param			comfort			query	int		false	"Comfort level"
// @Param			food_id			query	int		false	"Food ID"
// @Produce		json
// @Success		200	{array}		entity.SimplifiedTourView
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/api/v1/tours [get]
func (r *tourRoutes) getMany(c echo.Context) error {
	var inputSearch ToursSearch
	var inputFilter ToursFilter

	if err := c.Bind(&inputSearch); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid query parameters")
		return err
	}

	if err := c.Bind(&inputFilter); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid query parameters")
		return err
	}

	filters := ToFilters(inputSearch, inputFilter)

	fmt.Println(filters)

	tours, err := r.tourService.GetMany(c.Request().Context(), filters)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusOK, tours)
}

// @Summary		Get hot tours according to search and filter params
// @Description	Get hot tours according to search and filter params
// @Tags			tours
// @Produce		json
// @Param			from_name		query	string	false	"From location name"
// @Param			to_ame			query	string	false	"To location name"
// @Param			when			query	string	false	"Date of the tour"
// @Param			nights_count	query	int		false	"Number of nights in the tour"
// @Param			adults			query	int		false	"Number of adults"
// @Param			childrens		query	int		false	"Number of children"
// @Param			tags			query	string	false	"Type of the tour"
// @Param			price_from		query	string	false	"Minimum price"
// @Param			price_to		query	string	false	"Maximum price"
// @Param			rating			query	int		false	"Minimum rating"
// @Param			guaranteed		query	boolean	false	"Guaranteed availability"
// @Param			with_flight		query	bool	false	"Flight is included"
// @Param			with_acc		query	bool	false	"Accomodation is included"
// @Param			with_food		query	bool	false	"Nutrition is included""
// @Param			day_off			query	bool	false	"The tour takes place on a weekend"
// @Param			low_cost		query	bool	false	"Low cost tour"
// @Param			age_group		query	int		false	"Age group ID"
// @Param			difficulty		query	int		false	"Tour difficulty level"
// @Param			comfort			query	int		false	"Comfort level"
// @Param			foodId			query	int		false	"Food ID"
// @Produce		json
// @Success		200	{array}		entity.SimplifiedTourView
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/api/v1/tours/hot [get]
func (r *tourRoutes) getHot(c echo.Context) error {
	var inputSearch ToursSearch
	var inputFilter ToursFilter

	if err := c.Bind(&inputSearch); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid query parameters")
		return err
	}

	if err := c.Bind(&inputFilter); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid query parameters")
		return err
	}

	filters := ToFilters(inputSearch, inputFilter)

	tours, err := r.tourService.GetHotMany(c.Request().Context(), filters)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusOK, tours)
}

type tourGetByTagInput struct {
	Tag *string `param:"tags"`
}

// @Summary		Get reccomended tours according to search and filter params
// @Description	Get recommended tours according to search and filter params
// @Tags			tours
// @Produce		json
// @Param			from_name	query	string	false	"From location name"
// @Param			to_name		query	string	false	"To location name"
// @Param			when_date	query	string	false	"Date of the tour"
// @Param			nights_cnt	query	int		false	"Number of nights in the tour"
// @Param			adults		query	int		false	"Number of adults"
// @Param			childrens	query	int		false	"Number of children"
// @Param			tags		query	string	false	"Type of the tour"
// @Param			price_from	query	string	false	"Minimum price"
// @Param			price_to	query	string	false	"Maximum price"
// @Param			rating		query	int		false	"Minimum rating"
// @Param			guaranteed	query	boolean	false	"Guaranteed availability"
// @Param			with_flight	query	bool	false	"Flight is included"
// @Param			with_acc	query	bool	false	"Accomodation is included"
// @Param			with_food	query	bool	false	"Nutrition is included""
// @Param			day_off		query	bool	false	"The tour takes place on a weekend"
// @Param			low_cost	query	bool	false	"Low cost tour"
// @Param			age_group	query	int		false	"Age group ID"
// @Param			difficulty	query	int		false	"Tour difficulty level"
// @Param			comfort		query	int		false	"Comfort level"
// @Param			food_id		query	int		false	"Food ID"
// @Produce		json
// @Success		200	{array}		entity.SimplifiedTourView
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/api/v1/tours/recommendations [get]
func (r *tourRoutes) getRecommendations(c echo.Context) error {
	var input tourGetByTagInput

	var tags []string
	if strings.Contains(c.QueryParam("tags"), ",") {
		tags = strings.Split(c.QueryParam("tags"), ",")
	}

	if err := c.Bind(&input); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid query parameters")
		return err
	}

	filters := make(map[string]interface{}, 1)
	filters["tags"] = input.Tag
	if tags != nil {
		filters["tags"] = tags
	}

	tours, err := r.tourService.GetMany(c.Request().Context(), filters)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusOK, tours)
}

func (r *tourRoutes) makeOrder(c echo.Context) error {
	var order entity.Order

	if err := c.Bind(&order); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid query parameters")
		return err
	}

	if err := c.Validate(order); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	order, err := r.tourOrder.GetOrderById(c.Request().Context(), order.OrderId)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusOK, order)
}
