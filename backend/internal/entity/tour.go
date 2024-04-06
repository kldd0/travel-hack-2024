package entity

import (
	"time"
)

// выплевывает мини инфу о мини карточке
type SimplifiedTourView struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Location   string `json:"location"`
	Category   string `json:"category"`
	NightCount int    `json:"night_count"`
	Type       []Tag  `json:"type"`
	Rating     int    `json:"rating"`
}

type Tour struct {
	Id              int        `json:"id"`
	Title           string     `json:"title"`
	Location        string     `json:"location"`
	Category        string     `json:"category"`
	Type            []Tag      `json:"type"`
	Desc            string     `json:"description"`
	NightCount      int        `json:"night_count"`
	Program         []string   `json:"program"`
	Included        []string   `json:"included"`
	NotIncluded     []string   `json:"not_included"`
	DifficultyLevel string     `json:"difficulty_level"`
	ComfortLevel    string     `json:"comfort_level"`
	NearestDate     time.Time  `json:"nearest_date"` /* must be updated every req */
	GroupDates      []TourDate `json:"group_dates"`  /* must be updated every req */
	ImportantInfo   string     `json:"important_info"`
	Media           []Image    `json:"media"`
	Faq             string     `json:"faq"`
	Rating          int        `json:"rating"`
	Reviews         []Review   `json:"reviews"` // массиов отзывов
	/* embed count of companions and their age group */
}

type Tag struct {
	Title string `json:"title"`
}

type Image struct {
	Caption string `json:"caption"`
	URL     string `json:"url"`
}

type TourDate struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func SimplifyingTour(tour Tour) SimplifiedTourView {
	return SimplifiedTourView{
		Id:         tour.Id,
		Title:      tour.Title,
		Location:   tour.Location,
		Category:   tour.Category,
		NightCount: tour.NightCount,
		Type:       tour.Type,
		Rating:     tour.Rating,
	}
}
