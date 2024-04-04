package v1

import (
	"net/http"

	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/service"
	"github.com/labstack/echo/v4"
)

type reviewRoutes struct {
	reviewService service.Review
}

func newReviewRoutes(g *echo.Group, reviewService service.Review) *reviewRoutes {
	r := &reviewRoutes{
		reviewService: reviewService,
	}

	g.GET("", r.getById)

	return r
}

func (r *reviewRoutes) getById(c echo.Context) error {
	var input interface{}

	if err := c.Bind(&input); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	reviews, err := r.reviewService.GetMany(c.Request().Context())
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	type response struct {
		Reviews []entity.Review
	}

	return c.JSON(http.StatusOK, response{
		Reviews: reviews,
	})
}
