package v1

import (
	"net/http"
	"sort"
	"time"

	"github.com/kldd0/travel-hack-2024/internal/entity"
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
	FromName  string    `query:"fromName"`
	ToName    string    `query:"toName"`
	WhenDate  time.Time `query:"whenDate"`
	NightsCnt int       `query:"nightsCnt"`
	Adults    int       `query:"adults"`
	Childrens int       `query:"childrens"`
}

type ToursFilter struct {
	TourType   string `query:"tourType"`
	PriceFrom  string `query:"priceFrom"`
	PriceTo    string `query:"priceTo"`
	Rating     int    `query:"rating"`
	Guaranteed bool   `query:"guaranteed"`
	Features   Features
	AgeGroupId int `query:"ageGroupId"`
	Difficulty int `query:"difficulty"`
	Comfort    int `query:"comfort"`
	FoodId     int `query:"foodId"`
}

type Features struct {
	WithFlight       bool `query:"withFlight"`
	WithAccomodation bool `query:"withAcc"`
	WithFood         bool `query:"withNutrition"`
	DayOff           bool `query:"dayOff"`
	LowCost          bool `query:"lowCost" `
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
// @Param			toName		query	string	false	"To location name"
// @Param			whenDate	query	string	false	"Date of the tour"
// @Param			nightsCnt	query	int		false	"Number of nights in the tour"
// @Param			adults		query	int		false	"Number of adults"
// @Param			childrens	query	int		false	"Number of children"
// @Param			tourType	query	string	false	"Type of the tour"
// @Param			priceFrom	query	string	false	"Minimum price"
// @Param			priceTo		query	string	false	"Maximum price"
// @Param			rating		query	int		false	"Minimum rating"
// @Param			guaranteed	query	boolean	false	"Guaranteed availability"
// @Param			withFlight	query	bool	false	"Flight is included"
// @Param			withAcc		query	bool	false	"Accomodation is included"
// @Param			withFood	query	bool	false	"Nutrition is included""
// @Param			dayOff		query	bool	false	"The tour takes place on a weekend"
// @Param			lowCost		query	bool	false	"Low cost tour"
// @Param			ageGroupId	query	int		false	"Age group ID"
// @Param			difficulty	query	int		false	"Tour difficulty level"
// @Param			comfort		query	int		false	"Comfort level"
// @Param			foodId		query	int		false	"Food ID"
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

	tours, err := r.tourService.GetMany(c.Request().Context())
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	simplifiedTours := make([]entity.SimplifiedTourView, 0, len(tours))
	for _, tour := range tours {
		simplifiedTours = append(simplifiedTours, entity.SimplifyingTour(tour))
	}

	return c.JSON(http.StatusOK, simplifiedTours)
}

// @Summary		Get hot tours according to search and filter params
// @Description	Get hot tours according to search and filter params
// @Tags			tours
// @Produce		json
// @Param			fromName	query	string	false	"From location name"
// @Param			toName		query	string	false	"To location name"
// @Param			whenDate	query	string	false	"Date of the tour"
// @Param			nightsCnt	query	int		false	"Number of nights in the tour"
// @Param			adults		query	int		false	"Number of adults"
// @Param			childrens	query	int		false	"Number of children"
// @Param			tourType	query	string	false	"Type of the tour"
// @Param			priceFrom	query	string	false	"Minimum price"
// @Param			priceTo		query	string	false	"Maximum price"
// @Param			rating		query	int		false	"Minimum rating"
// @Param			guaranteed	query	boolean	false	"Guaranteed availability"
// @Param			withFlight	query	bool	false	"Flight is included"
// @Param			withAcc		query	bool	false	"Accomodation is included"
// @Param			withFood	query	bool	false	"Nutrition is included""
// @Param			dayOff		query	bool	false	"The tour takes place on a weekend"
// @Param			lowCost		query	bool	false	"Low cost tour"
// @Param			ageGroupId	query	int		false	"Age group ID"
// @Param			difficulty	query	int		false	"Tour difficulty level"
// @Param			comfort		query	int		false	"Comfort level"
// @Param			foodId		query	int		false	"Food ID"
// @Produce		json
// @Success		200	{object}	[]entity.SimplifiedTourView
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/api/v1/tours/hot [get]
func (r *tourRoutes) getHot(c echo.Context) error {
	var input toursSearchAndFilterInput

	if err := c.Bind(&input); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid query parameters")
		return err
	}

	tours, err := r.tourService.GetMany(c.Request().Context())
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	simplifiedTours := make([]entity.SimplifiedTourView, 0, len(tours))
	for _, tour := range tours {
		// check if tour has no dates
		if len(tour.Dates) < 1 {
			continue
		}

		// sort tour dates to get the nearest date to compare with today date
		// to find out if tour is hot (start in or less than 72 hours)
		sort.Slice(tour.Dates, func(i, j int) bool { return tour.Dates[i].Start.Before(tour.Dates[j].Start) })
		if !(time.Until(tour.Dates[0].Start) < time.Hour*72) {
			continue
		}

		simplifiedTours = append(simplifiedTours, entity.SimplifyingTour(tour))
	}

	return c.JSON(http.StatusOK, simplifiedTours)
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
